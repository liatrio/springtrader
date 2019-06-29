#!/bin/bash
echo '127.0.0.1 centos6 nanodbserver' >> /etc/hosts

# Start SQLFire
echo 'STARTING SQLFIRE DB'
sqlf locator start -peer-discovery-address=127.0.0.1 -peer-discovery-port=3241 -dir=/opt/vmware/vfabric-sqlfire/vFabric_SQLFire_103/locator1 -client-port=1527 -client-bind-address=127.0.0.1
sqlf server start -dir=/opt/vmware/vfabric-sqlfire/vFabric_SQLFire_103/server1 -client-bind-address=127.0.0.1 -client-port=1528 -locators=127.0.0.1[3241]
echo 'SQLFIRE DATABASE IS UP'

tail -f /dev/null
