#!/bin/bash

echo '127.0.0.1 centos6 nanodbserver springtrader rabbitmq' >> /etc/hosts
service rabbitmq-server start

tail -f /dev/null
