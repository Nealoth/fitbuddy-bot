FROM golang:1.16

RUN git clone https://github.com/Nealoth/fitbuddy-bot
WORKDIR ./fitbuddy-bot
RUN make build
ENTRYPOINT ["make", "container-run"]
