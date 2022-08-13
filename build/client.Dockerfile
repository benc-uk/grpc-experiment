FROM ubuntu:latest
WORKDIR /app 

RUN apt update && apt install -y wget dnsutils
COPY main ./client

CMD [ "bash" ]
