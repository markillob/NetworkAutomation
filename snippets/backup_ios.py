#!usr/bin/python
#coding: utf-8

import sys
import csv
from netmiko import ConnectHandler

ip = '192.168.73.137'
un = 'cisco'
pw = 'cisco'
platform='cisco_ios'
device = ConnectHandler(device_type=platform, ip=ip, username=un, password=pw)
disable_paging= device.disable_paging()
#fo = open ("show_version.txt","a")
output = device.send_command('show version', delay_factor=20)
print output
#fo.write(output)
#fo.close
close_ssh=device.disconnect()


