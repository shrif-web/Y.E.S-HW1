#!/bin/bash

SERVER_IP_ADDRESS=$1

########################## nginx ####################################
yum install epel-release
yum install nginx
cp nginx/www.yes.io.conf /etc/nginx/conf.d/

########################## install nodejs & yarn ####################
curl -sL https://rpm.nodesource.com/setup_10.x | bash -
yum install nodejs

curl --silent --location https://dl.yarnpkg.com/rpm/yarn.repo | tee /etc/yum.repos.d/yarn.repo
rpm --import https://dl.yarnpkg.com/rpm/pubkey.gpg
yum install yarn

########################## firewall #################################
firewall-cmd --permanent --zone=public --add-service=http
firewall-cmd --permanent --zone=public --add-service=https
firewall-cmd --reload

######################### selinux ###################################
setenforce 0

######################### service files #############################
cp /etc/systemd/system/



######################## running ####################################
systemctl enable nginx
systemctl start nginx

