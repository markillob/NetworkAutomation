# Network Automation Examples

*Uses cases to apply automation*

- Deploy same command on multiple devices
- "show ip route " and parse the output
- display all active interfaces on a list of cisco_ios routers
- [List all UP interfaces on a list of routers stored on a CSV file](https://github.com/markillob/NetworkAutomation/blob/master/snippets/if_up_ios_csv.py)
  
  ```
  python if_up_ios_csv.py
  GigabitEthernet1       192.168.73.137  YES DHCP   up                    up
  Ethernet0/0                192.168.73.138  YES DHCP   up                    up
  Ethernet0/1                192.168.73.139  YES DHCP   up                    up
  ```
CSV
 ```
192.168.73.137
192.168.73.138
192.168.73.139
 ```
  
