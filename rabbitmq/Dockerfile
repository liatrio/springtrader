FROM centos:centos6

EXPOSE 4369 5671 5672 25672

RUN mkdir -p /etc/vmware/vfabric/
RUN echo 'I_ACCEPT_EULA_LOCATED_AT=http://www.vmware.com/download/eula/vfabric_app-platform_eula.html' > /etc/vmware/vfabric/accept-vfabric5.1-eula.txt

RUN rpm -ivhf http://repo.vmware.com/pub/rhel6/vfabric/5.1/vfabric-5.1-repo-5.1-1.noarch.rpm
RUN rpm -Uvh https://download.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
RUN yum install vfabric-tc-server-standard erlang vfabric-rabbitmq-server -y

ADD entry.sh .
RUN chmod +x entry.sh

ENTRYPOINT ["./entry.sh"]
