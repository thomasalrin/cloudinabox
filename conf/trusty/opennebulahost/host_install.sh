#!/bin/bash

ONE_INSTALL_LOG="/var/log/megam/megamcib/opennebulahost.log"

echo "host_install.sh start execution ====>" >> $ONE_INSTALL_LOG

ping -c 1 downloads.opennebula.org &> /dev/null

sudo apt-get -y install build-essential genromfs autoconf libtool qemu-utils libvirt0 bridge-utils >> $ONE_INSTALL_LOG

sudo apt-get -y install lvm2 ssh iproute iputils-arping make >> $ONE_INSTALL_LOG

if [ $? -ne 0 ]; then
  echo "`date`: check your network connection. downloads.opennebula.org is not reachable!" >> $ONE_INSTALL_LOG
  exit 1
fi

wget -q -O- http://downloads.opennebula.org/repo/Ubuntu/repo.key | apt-key add -

echo "deb http://downloads.opennebula.org/repo/4.12/Ubuntu/14.04 stable opennebula" > /etc/apt/sources.list.d/opennebula.list

echo "File created /etc/apt/sources.list.d/opennebula.list " >> $ONE_INSTALL_LOG



sudo apt-get -y update


sudo apt-get -y install opennebula-node >> $ONE_INSTALL_LOG

echo "Changing password for oneadmin user(Password = oneadmin)" >> $ONE_INSTALL_LOG
sudo usermod -p $(echo oneadmin | openssl passwd -1 -stdin) oneadmin

echo "host_install.sh ends execution ====>" >> $ONE_INSTALL_LOG
