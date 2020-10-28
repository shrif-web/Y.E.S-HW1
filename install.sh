#!/bin/bash

SERVER_IP_ADDRESS=$1

########################## nginx ####################################
yum update -y
yum install epel-release -y
yum install nginx -y
cp nginx/www.yes.io.conf /etc/nginx/conf.d/

########################## install nodejs & yarn ####################
curl -sL https://rpm.nodesource.com/setup_10.x | bash -
yum install nodejs -y

curl --silent --location https://dl.yarnpkg.com/rpm/yarn.repo | tee /etc/yum.repos.d/yarn.repo
rpm --import https://dl.yarnpkg.com/rpm/pubkey.gpg
yum install yarn -y

########################## firewall #################################
firewall-cmd --permanent --zone=public --add-service=http
firewall-cmd --permanent --zone=public --add-service=https
firewall-cmd --reload

######################### selinux ###################################
setenforce 0

######################### service files #############################
PWD=$(pwd)
NODE_PATH=$(which node)

######################## go
chmod +x go/main
ln -s $PWD/go/main /usr/bin/go-server
####################### nodejs
sed "1 i\\#!$NODE_PATH " -i nodejs/main.js
chmod +x nodejs/main.js
ln -s $PWD/nodejs/main.js /usr/bin/node-server
cd nodejs
yarn install

cp services/go-server.service /etc/systemd/system/multi-user.target.wants/
cp services/nodejs-server.service /etc/systemd/system/multi-user.target.wants/

######################## yarn install ###############################

######################## running ####################################
systemctl daemon-reload
systemctl enable nginx
systemctl enable go-server
systemctl enable nodejs-server

systemctl start go-server
systemctl start nodejs-server
systemctl start nginx

