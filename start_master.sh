#!/usr/bin/env bash
git checkout master
git reset HEAD --hard
pm2 stop "piedpiper"
git pull
go build -o piedpiper
pm2 start "piedpiper"
pm2 logs "piedpiper"
