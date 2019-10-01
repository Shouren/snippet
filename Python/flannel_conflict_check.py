import paramiko
from ipaddress import ip_address as addr
from ipaddress import IPv4Interface as Iface


def get_host_list():
    return [
            ('172.27.129.20', 'root', '123456'),
            ('172.27.129.21', 'root', '123456'),
            ('172.27.129.22', 'root', '123456'),
            ]


class HostInfo():
    def __init__(self, host, user, password):
        self.host = host
        self. user = user
        self.passwd = password
        self.connect_host()
        self.get_vtep_ips_from_config()
        self.get_vtep_ips_from_dev()
        self.disconnect_host()

    def __str__(self):
        return "Host: {0}}".format(self.host)

    def connect_host(self):
        self.client = paramiko.SSHClient()
        self.client.load_system_host_keys()
        self.client.set_missing_host_key_policy(paramiko.client.WarningPolicy)
        self.client.connect(
                self.host, username=self.user, password=self.passwd)

    def disconnect_host(self):
        self.client.close()

    def get_vtep_ips_from_config(self):
        stdin, stdout, stderr = self.client.\
                exec_command("cat /var/run/flannel/subnet.env | grep SUBNET")
        self.config_info = []
        for line in stdout:
            data = line.strip("\n").split("=")[1]
            ret = data.split("/")
            ip, mask = ret[0], ret[1]
            self.config_info.append((ip, mask))

    def get_vtep_ips_from_dev(self):
        stdin, stdout, stderr = self.client.\
                exec_command("ip -o addr show | grep flannel | grep -v inet6")
        self.dev_info = []
        for line in stdout:
            data = line.strip("\n").split()[3]
            ret = data.split("/")
            ip, mask = ret[0], ret[1]
            self.dev_info.append((ip, mask))

    def check_config_error(self):
        config_ips = list(map(
                lambda x: Iface(x[0]+"/"+x[1]).network.network_address,
                self.config_info))
        err_ips = [x for x in self.dev_info if addr(x[0]) not in config_ips]
        e = ""
        if len(err_ips) > 0:
            e += "To prevent IP conflict in the future, fix {0} with:\n"\
                    .format(self.host)
        for data in err_ips:
            e += "  'ip addr del dev flannel.1 {1}/{2}'\n".format(
                self.host, data[0], data[1])
        if len(e) > 0:
            print(e + "\n")


if __name__ == "__main__":
    targets = get_host_list()
    vtep_hosts = {}
    for host_info in targets:
        host, user, passwd = host_info
        host_obj = HostInfo(host, user, passwd)
        host_obj.check_config_error()
        for info in host_obj.config_info:
            vtep_hosts.setdefault(info, []).append(host_obj.host)
    for hosts in vtep_hosts.values():
        print(hosts)
        if len(hosts) == 1:
            continue
        err = "Flannel IP conflicts between Hosts: "
        err += str(hosts)
        print(err)
