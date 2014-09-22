FROM scratch
MAINTAINER CoreOS

EXPOSE 80
WORKDIR /opt/validate
ENTRYPOINT ["bin/validate"]

ADD bin/validate /opt/validate/bin/validate
