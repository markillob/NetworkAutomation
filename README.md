# Network Automation Examples

* Networking uses cases to apply automation*

- Deploy same command on multiple devices
- [Backup configurations from multiple devices and store them as a txt files](pending)
- "show ip route " and parse the output
- display all active interfaces on a list of cisco_ios routers
```python
python interfaces_up.py 192.168.73.137 cisco_ios cisco cisco
GigabitEthernet1       192.168.73.137  YES DHCP   up                    up
```
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
  
