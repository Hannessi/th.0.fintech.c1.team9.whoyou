FROM centos:7

RUN yum update -y
RUN yum upgrade
RUN yum install -y httpd

COPY ./build/ /var/www/html/
COPY ./httpd.conf /etc/httpd/conf/
COPY ./.htaccess /var/www/html/

ENTRYPOINT ["/usr/sbin/httpd", "-D", "FOREGROUND"]

