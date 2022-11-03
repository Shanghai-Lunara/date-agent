FROM ubuntu:20.04

WORKDIR /server/
COPY date-agent /server/
COPY templates /server/templates
COPY statics /server/statics

CMD ["/server/date-agent", "-alsologtostderr=true", "-v=4"]
