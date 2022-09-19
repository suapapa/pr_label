#!/bin/bash/
GOOS=linux GOARCH=arm go build -o pr_label
scp pr_label pi@rpi-airplay.local:~/ # 192.168.219.104
