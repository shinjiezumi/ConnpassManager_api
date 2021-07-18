#!/bin/bash
cd /usr/share/nginx/html/webapps/go/connpass-manager_api
git pull
cd src
go build -o ./api main.go
sudo supervisorctl restart connpass-manager