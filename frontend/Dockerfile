FROM alpine:3.3

RUN apk update && \ 
    apk add \ 
      ca-certificates && \
    rm -rf /var/cache/apk/*
EXPOSE 3001
ADD frontend /
CMD ["/frontend"]
