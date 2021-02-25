FROM alpine
RUN apk add --no-cache coreutils bash curl

WORKDIR /app
COPY ./wait_for_status_code.sh .
RUN chmod +x ./wait_for_status_code.sh

ENTRYPOINT ["/app/wait_for_status_code.sh"]