#!usr/bin/python
#coding: utf-8
import sys
import csv
from netmiko import ConnectHandler
platform='cisco_ios'
#username=raw_input(' Username --> ')
username='cisco'
password='cisco'
#password=raw_input(' Password --> ')''
#cef=raw_input(' Configuration to search --> ')
csv_file=sys.argv[1]

def read_csv(csvname):
    ip_host=[]
    with open(csvname, 'rb') as f:
        reader = csv.reader(f)
        for result in reader:
                result = ''.join(result)
                ip_host.append(result)
    return ip_host
def backup_configurations(devices):
    for ip_add in devices:
        device = ConnectHandler(device_type=platform, ip=ip_add, username=username, password=password)
        no_breakpage = device.disable_paging()
        output = device.send_command_expect('show version')
        print type(output)
        print "----------------------------------------"
        #Interfaces = output.split("\n")
        #for i in Interfaces:
        #    if cef in i:net_connect.send_command_timing
        #       print i
    return  
devices=read_csv(csv_file)
print backup_configurations(devices)



'''
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
'''