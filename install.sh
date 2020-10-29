#!/bin/bash

SERVER_IP_ADDRESS=$1

########################## nginx ####################################
yum update -y
yum install epel-release -y
yum install nginx -y
sed "s@<path_to_front>@$(pwd)\/front@gm" nginx/www.yes.io.conf > /etc/nginx/conf.d/www.yes.io.conf

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
NODE_PATH=$(which node)

######################## go
chmod +x go/main
ln -s $(pwd)/go/main /usr/bin/go-server
####################### nodejs
sed "1 i\\#!$NODE_PATH " -i nodejs/main.js
chmod +x nodejs/main.js
ln -s $(pwd)/nodejs/main.js /usr/bin/node-server

cd nodejs
  yarn install
cd ..

sed "s@<path_to_binary>@$(pwd)\/go@gm" services/go-server.service > /etc/systemd/system/multi-user.target.wants/go-server.service
sed "s@<path_to_binary>@$(pwd)\/nodejs@gm" services/nodejs-server.service > /etc/systemd/system/multi-user.target.wants/nodejs-server.service

#cp services/go-server.service /etc/systemd/system/multi-user.target.wants/
#cp services/nodejs-server.service /etc/systemd/system/multi-user.target.wants/

######################## yarn install ###############################

######################## running ####################################
systemctl daemon-reload
systemctl enable nginx
systemctl enable go-server
systemctl enable nodejs-server

systemctl restart go-server
systemctl restart nodejs-server
systemctl restart nginx

nginx -s reload
