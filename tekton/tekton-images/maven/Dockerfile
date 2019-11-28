FROM debian:stretch-slim

USER root

ENV MAVEN_VERSION 3.5.4
ENV MAVEN_HOME /usr/local/maven
ENV JAVA_HOME /usr/local/java/jdk1.8.0_171
ENV PATH $MAVEN_HOME/bin:$JAVA_HOME/bin:$PATH

COPY apache-maven-3.5.4-bin.tar.gz .
#RUN wget http://archive.apache.org/dist/maven/maven-3/$MAVEN_VERSION/binaries/apache-maven-$MAVEN_VERSION-bin.tar.gz && \
RUN  tar -zxvf apache-maven-$MAVEN_VERSION-bin.tar.gz && \
  rm apache-maven-$MAVEN_VERSION-bin.tar.gz && \
  mv apache-maven-$MAVEN_VERSION /usr/local/maven

COPY settings.xml /usr/local/maven/conf/

COPY jdk1.8.0_171 /usr/local/java/jdk1.8.0_171
