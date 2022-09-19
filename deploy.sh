#!/bin/bash
GOOS=linux GOARCH=arm go build -o pr_order
scp pr_order pi@rpi-airplay.local:~/ # 192.168.219.104
