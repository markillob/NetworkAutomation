#!/bin/bash
a=1
for i in {1001..1039}
do
printf "\ninterface port-channel"$a
printf "\nswitchport trunk allowed vlan add 3100"
printf "\ninterface Ethernet1/"$a
printf "\ndescription SERVER_"$i
printf "\nswitchport mode trunk"
printf "\nswitchport trunk native vlan "$i
printf "\nswitchport trunk allowed vlan 70,72-74,1000,3100,3102,"$i
printf "\nspanning-tree port type edge trunk"
printf "\nchannel-group %s mode active" $a
printf "\nno shutdown\n"
a=$((a + 1))
done
