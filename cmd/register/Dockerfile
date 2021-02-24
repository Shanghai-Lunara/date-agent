FROM ubuntu:18.04

LABEL maintainer="lunara-developer@lunara.net"

WORKDIR /server/
COPY date-agent /server/
COPY templates /server/templates
COPY statics /server/statics

CMD ["/server/date-agent", "-alsologtostderr=true", "-v=4"]
