FROM alpine:latest
RUN apk update && apk add mysql-client && mysql --version && apk add curl && curl --version

# https://github.com/go-testfixtures/testfixtures/releases
RUN curl -OL https://github.com/go-testfixtures/testfixtures/releases/download/v3.12.0/testfixtures_linux_arm64.tar.gz
RUN tar xf testfixtures_linux_arm64.tar.gz -C /usr/local/bin/ && testfixtures --version

COPY db/seed ./seed
RUN chmod +x ./seed/entrypoint.sh
ENTRYPOINT ["seed/entrypoint.sh"]
