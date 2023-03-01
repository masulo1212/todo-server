FROM alpine:latest

RUN mkdir /app

COPY db/migration ./db/migration
COPY brokerApp /app

CMD [ "/app/brokerApp" ]