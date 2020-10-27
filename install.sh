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
cp services/go-server.service /etc/systemd/system/multi-user.target.wants/
cp services/nodejs-server.service /etc/systemd/system/multi-user.target.wants/

######################## running ####################################
systemctl daemon-reload
systemctl enable nginx
systemctl enable go-server
systemctl enable nodejs-server

systemctl start go-server
systemctl start nodejs-server
systemctl start nginx

