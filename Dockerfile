FROM centos:centos6 

# Accept VMware certificate
RUN mkdir -p /etc/vmware/vfabric/ && \
    echo 'I_ACCEPT_EULA_LOCATED_AT=http://www.vmware.com/download/eula/vfabric_app-platform_eula.html' \
    > /etc/vmware/vfabric/accept-vfabric5.1-eula.txt

# Install vFabric software
RUN rpm -ivhf http://repo.vmware.com/pub/rhel6/vfabric/5.1/vfabric-5.1-repo-5.1-1.noarch.rpm && \
    yum install wget unzip java-1.7.0-openjdk-devel vfabric-tc-server-standard -y
    
# Build the application
WORKDIR /app
COPY . .
RUN ./gradlew clean build release

# Copy artifacts to server
WORKDIR /opt/vmware/vfabric-tc-server-standard
RUN ./tcruntime-instance.sh create springtrader -t springtrader -f templates/springtrader/sqlfire.properties
COPY /springtrader/dist/spring-nanotrader-asynch-services-1.0.1.BUILD-SNAPSHOT.war springtrader/webapps/spring-nanotrader-asynch-services.war
COPY /springtrader/dist/spring-nanotrader-services-1.0.1.BUILD-SNAPSHOT.war springtrader/webapps/spring-nanotrader-services.war
COPY /springtrader/dist/spring-nanotrader-web-1.0.1.BUILD-SNAPSHOT.war springtrader/webapps/spring-nanotrader-web.war

WORKDIR /app
ENTRYPOINT ["/opt/vmware/vfabric-tc-server-standard/springtrader/bin/tcruntime-ctl.sh", "run", "springtrader"]