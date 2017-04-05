import sys
import csv
from netmiko import ConnectHandler
from ciscoconfparse import CiscoConfParse

host=sys.argv[1]
platform=sys.argv[2]
username=sys.argv[3]
password=sys.argv[4]
device = ConnectHandler(device_type=platform, ip=host, username=username, password=password)
output = str(device.send_command('show ip int brief'))
Interfaces = output.split("\n")
string='up'
string2='Up'
for i in Interfaces:
    if string in i:
        print i
#connect to a cisco_ios a list all active interfaces on vrf default
