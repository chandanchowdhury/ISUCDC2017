#!/bin/bash

#if [[ $EUID -ne 0 ]]; then
#   echo "This script must be run as root" 1>&2
#   exit 1
#fi

go get
go build
sudo -H -u www-data ./api 0.0.0.0:8080