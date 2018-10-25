#!usr/bin/python
#coding: utf-8
import sys
import csv
from netmiko import ConnectHandler
platform='cisco_ios'
username='cisco' 
password='cisco'
ip_host=[]
i=0
with open('hosts.csv', 'rb') as f:
    reader = csv.reader(f)
    for result in reader:
        result = ''.join(result)
        ip_host.append(result)
for ip_add in ip_host:
    device = ConnectHandler(device_type=platform, ip=ip_add, username=username, password=password)
    output = str(device.send_command('show ip int brief'))
    Interfaces = output.split("\n")
    string='up'
    string2='Up'
    #print output
    for i in Interfaces:
        if string in i:
            print i
