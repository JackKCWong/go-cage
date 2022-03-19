#!/usr/bin/env bash


sudo umount /tmp/empty/proc
sudo rm -rf /tmp/empty/.pivot_root

go build

sudo ./go-cage /tmp/empty /lsroot
