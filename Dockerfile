FROM centos:centos6
RUN yum install git -y
RUN yum install wget -y
RUN yum install unzip -y
RUN yum install java-1.7.0-openjdk-devel -y

# Install Groovy
WORKDIR /usr
RUN mkdir groovy 
WORKDIR groovy
RUN wget http://dl.bintray.com/groovy/maven/groovy-binary-2.3.0-beta-2.zip
RUN unzip groovy-binary-2.3.0-beta-2.zip
RUN rm -f groovy-binary-2.3.0-beta-2.zip 
ENV GROOVY_HOME=/usr/groovy/groovy-2.3.0-beta-2
ENV PATH=$PATH:$GROOVY_HOME/bin
RUN echo $PATH


# Accept VMware certificate 
RUN mkdir -p /etc/vmware/vfabric/
RUN echo 'I_ACCEPT_EULA_LOCATED_AT=http://www.vmware.com/download/eula/vfabric_app-platform_eula.html' > /etc/vmware/vfabric/accept-vfabric5.1-eula.txt

RUN rpm -ivhf http://repo.vmware.com/pub/rhel6/vfabric/5.1/vfabric-5.1-repo-5.1-1.noarch.rpm
RUN yum install vfabric-tc-server-standard -y 
RUN yum install vfabric-sqlfire -y 
RUN rpm -Uvh https://download.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
RUN yum install erlang -y
RUN yum install vfabric-rabbitmq-server -y 
#RUN echo '127.0.0.1 centos6 nanodbserver' >> /etc/hosts
ENV JAVA_HOME=/usr/

#RUN service rabbitmq-server start

WORKDIR /opt/vmware/vfabric-sqlfire/vFabric_SQLFire_103
RUN mkdir locator1 server1
#RUN sqlf locator start -peer-discovery-address=127.0.0.1 -peer-discovery-port=3241 -dir=locator1 -client-port=1527 -client-bind-address=127.0.0.1
#RUN sqlf server start -dir=server1 -client-bind-address=127.0.0.1 -client-port=1528 -locators=127.0.0.1[3241]

WORKDIR /root
RUN git clone https://github.com/vFabric/springtrader.git
RUN cp /opt/vmware/vfabric-sqlfire/vFabric_SQLFire_103/lib/sqlfireclient.jar /root/springtrader/lib/sqlfireclient-1.0.3.jar
RUN cp /opt/vmware/vfabric-sqlfire/vFabric_SQLFire_103/lib/sqlfireclient.jar /root/springtrader/templates/springtrader/lib/sqlfireclient-1.0.3.jar
RUN wget https://repo.spring.io/plugins-release/com/gemstone/gemfire/gemfire/8.2.0/gemfire-8.2.0.jar
RUN cp gemfire-8.2.0.jar /root/springtrader/lib/gemfire.jar
ENV GRADLE_OPTS='-Xmx1024m -Xms256m -XX:MaxPermSize=512m'

WORKDIR /root/springtrader
RUN ./gradlew clean build release
RUN mkdir dist/libs

WORKDIR /root/springtrader/dist
RUN unzip DataGenerator.zip 
#RUN ./createSqlfSchema

WORKDIR /root/springtrader
RUN cp -r /root/springtrader/templates/springtrader /opt/vmware/vfabric-tc-server-standard/templates

WORKDIR /opt/vmware/vfabric-tc-server-standard
RUN ./tcruntime-instance.sh create springtrader -t springtrader -f templates/springtrader/sqlfire.properties
RUN cp /root/springtrader/dist/spring-nanotrader-asynch-services-1.0.1.BUILD-SNAPSHOT.war /opt/vmware/vfabric-tc-server-standard/springtrader/webapps/spring-nanotrader-asynch-services.war
RUN cp /root/springtrader/dist/spring-nanotrader-services-1.0.1.BUILD-SNAPSHOT.war /opt/vmware/vfabric-tc-server-standard/springtrader/webapps/spring-nanotrader-services.war
RUN cp /root/springtrader/dist/spring-nanotrader-web-1.0.1.BUILD-SNAPSHOT.war /opt/vmware/vfabric-tc-server-standard/springtrader/webapps/spring-nanotrader-web.war

WORKDIR /opt/vmware/vfabric-tc-server-standard/springtrader/bin
RUN echo 'JVM_OPTS="-Xmx1024m -Xss192K -XX:MaxPermSize=192m"' >> setenv.sh
RUN yes | cp /root/springtrader/lib/sqlfireclient-1.0.3.jar /opt/vmware/vfabric-tc-server-standard/springtrader/lib/sqlfireclient-1.0.3.jar
COPY entry.sh .
RUN chmod +x entry.sh

ENTRYPOINT ["./entry.sh"]
###Access Web UI at http://localhost:8080/spring-nanotrader-web/#login
