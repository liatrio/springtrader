#!/bin/bash

echo '127.0.0.1 centos6 nanodbserver springtrader rabbitmq' >> /etc/hosts
service rabbitmq-server start
echo 'RABBITMQ STARTED'

tail -f /dev/null
