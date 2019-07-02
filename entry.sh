#!/bin/bash

# Make sure SQLfire is up
echo 'GOING TO SLEEP'
sleep 180
echo 'DONE SLEEPING'

echo '127.0.0.1 centos6 nanodbserver springtrader rabbitmq' >> /etc/hosts

cd /app
echo 'createSqlfSchema'
./createSqlfSchema
echo 'SPRINGTRADER START'
/opt/vmware/vfabric-tc-server-standard/springtrader/bin/tcruntime-ctl.sh start springtrader
cd /app
echo 'GENERATE DATA'
./generateData
tail -f /dev/null
