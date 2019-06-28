#!/bin/bash
export GROOVY_HOME=/usr/groovy/groovy-2.3.0-beta-2
export PATH=$PATH:$GROOVY_HOME/bin

# Install Groovy
mkdir /usr/groovy
cd /usr/groovy/
wget http://dl.bintray.com/groovy/maven/groovy-binary-2.3.0-beta-2.zip
unzip groovy-binary-2.3.0-beta-2.zip && rm -f groovy-binary-2.3.0-beta-2.zip

# Accept VMWare scripts
mkdir -p /etc/vmware/vfabric/
echo 'I_ACCEPT_EULA_LOCATED_AT=http://www.vmware.com/download/eula/vfabric_app-platform_eula.html' > /etc/vmware/vfabric/accept-vfabric5.1-eula.txt

# Install SQLFire
rpm -ivhf http://repo.vmware.com/pub/rhel6/vfabric/5.1/vfabric-5.1-repo-5.1-1.noarch.rpm
yum install vfabric-sqlfire -y
mkdir /opt/vmware/vfabric-sqlfire/vFabric_SQLFire_103/locator1
mkdir /opt/vmware/vfabric-sqlfire/vFabric_SQLFire_103/server1

# Install RabbitMQ
rpm -Uvh https://download.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
yum install erlang -y
#yum install vfabric-rabbitmq-server -y

echo '127.0.0.1 centos6 nanodbserver springtrader rabbitmq' >> /etc/hosts
#service rabbitmq-server start

# Start SQLFire
sleep 30
sqlf locator start -peer-discovery-address=127.0.0.1 -peer-discovery-port=3241 -dir=/opt/vmware/vfabric-sqlfire/vFabric_SQLFire_103/locator1 -client-port=1527 -client-bind-address=127.0.0.1
sqlf server start -dir=/opt/vmware/vfabric-sqlfire/vFabric_SQLFire_103/server1 -client-bind-address=127.0.0.1 -client-port=1528 -locators=127.0.0.1[3241]

cd /dist
./createSqlfSchema
/opt/vmware/vfabric-tc-server-standard/springtrader/bin/tcruntime-ctl.sh start springtrader
sleep 30
cd /dist
./generateData
tail -f /dev/null
