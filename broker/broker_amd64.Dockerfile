FROM alpine:latest

RUN mkdir /app

COPY brokerApp_amd64 /app

CMD [ "/app/brokerApp_amd64" ]