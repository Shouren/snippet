#! /usr/env python

from time import sleep
from scapy.all import srp, Ether, IP, TCP

hosts = ['128.199.107.127', '198.35.44.222']
result = {}

for h in hosts:
    result.setdefault(h, None)

def tcp_ping(hosts):
    packet = Ether()/IP(dst=hosts)/TCP(dport=80, flags='S')
    ans, unans = srp(packet, filter='tcp', verbose=0, timeout=2)

    for ret in ans:
        delta = ret[1].time - ret[0].sent_time
        if ret[1][IP].src in result:
            result[ret[1][IP].src] = delta

if __name__ == '__main__':
    for count in range(0, 30):
        tcp_ping(hosts)
        print "Count: {0}".format(count)
        for (host, delay) in result.iteritems():
            print '    Host: {0} Time: {1}'.format(host, delay*1000)
        sleep(1)
