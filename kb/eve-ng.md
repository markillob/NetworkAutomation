# Basics

## Create own linux

- Download ISO and copy to eve-ng server
- Create directory and rename the fiie 
```
mv $HOME/ubuntu-16.04.6-server-amd64.iso /opt/unetlab/addons/qemu/linux-ubuntu-xenial/cdrom.iso
```
- Create QEMU IMG, size based on your needs.
```
/opt/qemu/bin/qemu-img create -f qcow2 virtioa.qcow2 15G
```
- Create a new lab and add a new linux node, start normal Linux installation and  when the installation is complete, last confirmation screen, shutdown the node (via eve-ng-GUI ) and remove the `cdrom.iso `, from the follwing directory
```
/opt/unetlab/addons/qemu/linux-ubuntu-xenial/cdrom.iso
```
- Get the lab ID from `Lab Details` from eve-ng-GUI,  something like  `0270944f-c23e-4fba-8...`, and move to that directory on command line 
```
cd /opt/unetlab/tmp/0/2a912b88-5041-4273-9a32-4e78e88/1/
```
- Delete the virtioa.qcow2 from default eve-ng linux location.
```
rm /opt/unetlab/addons/qemu/linux-ubuntu-xenial/virtioa.qcow2
```
- Create a dummy file on the above directory and commit the image 
```
touch /opt/unetlab/addons/qemu/linux-ubuntu-xenial/virtioa.qcow2
/opt/qemu/bin/qemu-img commit virtioa.qcow2

Image committed.
```
- You should be able to run your custom Linux anytime you add a new node.