FROM golang:1.13.4

RUN mkdir /cron

ADD . /cron

WORKDIR /cron


CMD ["go","run","cronjob.go"]
