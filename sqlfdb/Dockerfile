FROM centos:centos6
ENV JAVA_HOME=/usr

RUN yum install java-1.7.0-openjdk-devel -y

# Accept VMWare scripts
RUN mkdir -p /etc/vmware/vfabric/
RUN echo 'I_ACCEPT_EULA_LOCATED_AT=http://www.vmware.com/download/eula/vfabric_app-platform_eula.html' > /etc/vmware/vfabric/accept-vfabric5.1-eula.txt

# Install SQLFire
RUN rpm -ivhf http://repo.vmware.com/pub/rhel6/vfabric/5.1/vfabric-5.1-repo-5.1-1.noarch.rpm
RUN yum install vfabric-sqlfire -y
RUN mkdir /opt/vmware/vfabric-sqlfire/vFabric_SQLFire_103/locator1
RUN mkdir /opt/vmware/vfabric-sqlfire/vFabric_SQLFire_103/server1
