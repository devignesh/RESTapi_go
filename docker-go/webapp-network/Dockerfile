FROM ubuntu:16.04

RUN apt update && apt install -y arp-scan

RUN apt install -y iputils-ping \
    iproute2

COPY webapp /

CMD [ "/bin/bash" ]