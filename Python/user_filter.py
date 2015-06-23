#! /usr/bin/env python3

import re
import subprocess

#res = subprocess.Popen(["ls", "-lh"], stdout=subprocess.PIPE)

# Get ip and tenant from server-list
f = open('./server_list', 'r')
ip_to_tenant = dict()

for line in f:
    rec = line.split('|')
    ips = rec[3].strip().split(',')
    if len(ips) == 2:
        ip_to_tenant[ips[1]] = rec[9].strip()

#print(ip_to_tenant)
f.close()

# Get ip and domains 
f = open('./domains', 'r')
ip_to_domains = dict()
user = dict()

for line in f:
    rec = line.split()
    if rec[1].strip() in ip_to_domains:
        ip_to_domains[rec[1].strip()].append(rec[0].strip())
    else:
        ip_to_domains[rec[1].strip()] = list()
        ip_to_domains[rec[1].strip()].append(rec[0].strip())

    if rec[1].strip() in ip_to_tenant:
        if rec[1].strip() in user:
            user[ip_to_tenant[rec[1].strip()]].append(ip_to_domains[rec[1].strip()])
        else:
            user[ip_to_tenant[rec[1].strip()]] = list()
            user[ip_to_tenant[rec[1].strip()]].append(ip_to_domains[rec[1].strip()])

f.close()

#print(ip_to_domains)
#print (user)

# write tenant name to file tenant_name

f = open('./tenant_email', 'w')
for key in user:
    f.write(key + ' ' + str(user[key]) + '\n')
f.close

