# Palo Alto everyday tasks

Initial setup PA-VM 

### Baseline configuration

#### MGMT IP address  
- Load PA-VM image and wait for the prompt #PA-VM, password admin:admin

```
configure
set deviceconfig system type static
set deviceconfig system ip-address {{ firewall ip }}    netmask {{ netmask }}default-gateway {{ gateway IP} }dns-setting servers primary {{ dns ip }}
```
#### Default username:password and hostname

```
#change password default user
set mgt-config users admin password

# new user
set mgt-config users {{ admin-user}} permissions role-based superuser yes 
set mgt-config users {{ admin-user}} password
# {{ password }}
set deviceconfig system hostname {{ hostname }}
set deviceconfig system domain {{ domain }
```

#### SSL profile for MGMT access 

```
set shared ssl-tls-service-profile {{ CERT PROFILE} protocol-settings min-version tls1-2
set shared ssl-tls-service-profile {{ CERT PROFILE} protocol-settings max-version max
set shared ssl-tls-service-profile {{ CERT PROFILE} certificate fw1-sfo
```

#### IP address assigment 
create object 
```
set address {{ object-name }} ip-netmask {{ ip-address}}/{{ subnet-mask}}
```
Assign object to physical interface

```
set zone {{ zone-name }} network layer3 {{ interface-name }}
set network virtual-router default interface {{ interface-name }}
set network interface ethernet {{ interface-name }} layer3 ndp-proxy enabled no
set network interface ethernet {{ interface-name }} layer3 ip {{ object-name }}
set network interface ethernet {{ interface-name }} layer3 lldp enable no
```
####  Allow ICMP 

```
set rulebase security rules ALLOW_ICMP to any
set rulebase security rules ALLOW_ICMP from any
set rulebase security rules ALLOW_ICMP source any
set rulebase security rules ALLOW_ICMP destination any
set rulebase security rules ALLOW_ICMP source-user any
set rulebase security rules ALLOW_ICMP category any
set rulebase security rules ALLOW_ICMP application [ icmp ipv6-icmp ]
set rulebase security rules ALLOW_ICMP service application-default
set rulebase security rules ALLOW_ICMP hip-profiles any
set rulebase security rules ALLOW_ICMP action allow
```


##### Reference

[ Set Up a Firewall Administrative Account and Assign CLI Privileges ](https://docs.paloaltonetworks.com/pan-os/9-0/pan-os-cli-quick-start/get-started-with-the-cli/give-administrators-access-to-the-cli/set-up-a-firewall-administrative-account-and-assign-cli-privileges)
