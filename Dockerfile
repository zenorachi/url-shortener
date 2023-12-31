FROM alpine:latest

WORKDIR /root

# Copy binary
COPY ./.bin/app /root/.bin/app

# Copy configs
COPY ./.env /root/
COPY ./configs/main.yml /root/configs/main.yml

COPY ./docs /root/docs

CMD ["./.bin/app"]