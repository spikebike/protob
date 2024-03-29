#!/bin/bash
# call this script with an email address (valid or not).
# like:
# ./makecert.sh joe@random.com
if [ ! "$#" -eq 1 ];then
	echo "You MUST include an email address as an argument"
	exit;
fi
if [ ! -d certs ]; then
	mkdir certs
fi
rm -f certs/*
echo
echo "make server cert"
openssl req -new -nodes -x509 -out certs/server.pem -keyout certs/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=$1"
echo
echo "make client cert"
openssl req -new -nodes -x509 -out certs/client.pem -keyout certs/client.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=$1"

