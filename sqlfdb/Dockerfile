FROM centos:centos6
ENV JAVA_HOME=/usr

RUN yum install git wget unzip java-1.7.0-openjdk-devel -y
RUN rpm -ivhf http://repo.vmware.com/pub/rhel6/vfabric/5.1/vfabric-5.1-repo-5.1-1.noarch.rpm
RUN rpm -Uvh https://download.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
COPY entry.sh .
RUN chmod +x entry.sh

ENTRYPOINT ["./entry.sh"]
