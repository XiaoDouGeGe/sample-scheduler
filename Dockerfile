FROM debian:stretch-slim

WORKDIR /

COPY bin/sample-scheduler /usr/local/bin

RUN chmod +x /usr/local/bin/sample-scheduler

CMD ["sample-scheduler"]
