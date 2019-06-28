FROM openjdk:7 as builder
RUN mkdir springtrader
WORKDIR /springtrader
ADD build.gradle gradle.properties settings.gradle gradlew ./
ADD .wrapper/ ./.wrapper
RUN ./gradlew build
ADD . .
RUN ./gradlew clean build release

################################################################################

FROM centos:centos6
ENV GROOVY_HOME=/usr/groovy/groovy-2.3.0-beta-2
ENV PATH=$PATH:$GROOVY_HOME/bin
ENV JAVA_HOME=/usr

RUN yum install git wget unzip java-1.7.0-openjdk-devel -y

# Install Groovy
RUN mkdir /usr/groovy/
WORKDIR groovy/
RUN wget http://dl.bintray.com/groovy/maven/groovy-binary-2.3.0-beta-2.zip
RUN unzip groovy-binary-2.3.0-beta-2.zip && rm -f groovy-binary-2.3.0-beta-2.zip

# Accept VMware certificate
RUN mkdir -p /etc/vmware/vfabric/
RUN echo 'I_ACCEPT_EULA_LOCATED_AT=http://www.vmware.com/download/eula/vfabric_app-platform_eula.html' > /etc/vmware/vfabric/accept-vfabric5.1-eula.txt

# Install vFabric software
RUN rpm -ivhf http://repo.vmware.com/pub/rhel6/vfabric/5.1/vfabric-5.1-repo-5.1-1.noarch.rpm
RUN rpm -Uvh https://download.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
RUN yum install vfabric-tc-server-standard erlang -y

# Handle SQLFire jars
WORKDIR /
COPY --from=builder /springtrader/dist /dist
RUN wget https://repo.spring.io/plugins-release/com/vmware/sqlfire/sqlfireclient/1.0.3/sqlfireclient-1.0.3.jar -O sqlfireclient.jar
RUN mkdir /templates

WORKDIR /dist
RUN unzip DataGenerator.zip
RUN cp /sqlfireclient.jar /dist/libs/sqlfireclient-1.0.3.jar

# Retrieve SpringTrader template
COPY --from=builder /springtrader/templates/ /templates
RUN ls /templates
RUN cp /sqlfireclient.jar /templates/springtrader/lib/sqlfireclient.jar

# Copy artifacts to server
WORKDIR /opt/vmware/vfabric-tc-server-standard
RUN cp -r /templates/springtrader/ templates/
RUN ./tcruntime-instance.sh create springtrader -t springtrader -f templates/springtrader/sqlfire.properties
RUN cp /dist/spring-nanotrader-asynch-services-1.0.1.BUILD-SNAPSHOT.war /opt/vmware/vfabric-tc-server-standard/springtrader/webapps/spring-nanotrader-asynch-services.war
RUN cp /dist/spring-nanotrader-services-1.0.1.BUILD-SNAPSHOT.war /opt/vmware/vfabric-tc-server-standard/springtrader/webapps/spring-nanotrader-services.war
RUN cp /dist/spring-nanotrader-web-1.0.1.BUILD-SNAPSHOT.war /opt/vmware/vfabric-tc-server-standard/springtrader/webapps/spring-nanotrader-web.war

WORKDIR /opt/vmware/vfabric-tc-server-standard/springtrader/bin
RUN echo 'JVM_OPTS="-Xmx1024m -Xss192K -XX:MaxPermSize=192m"' >> setenv.sh
RUN yes | cp /sqlfireclient.jar /opt/vmware/vfabric-tc-server-standard/springtrader/lib/sqlfireclient.jar
COPY entry.sh .
RUN chmod +x entry.sh

ENTRYPOINT ["./entry.sh"]
