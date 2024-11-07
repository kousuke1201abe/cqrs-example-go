FROM alpine:latest
ARG CPUARCH=amd64
RUN apk update && apk add mysql-client && mysql --version && apk add curl && curl --version

# https://github.com/k0kubun/sqldef/releases
RUN curl -OL https://github.com/k0kubun/sqldef/releases/download/v0.16.9/mysqldef_linux_${CPUARCH}.tar.gz
RUN tar xf mysqldef_linux_${CPUARCH}.tar.gz -C /usr/local/bin/ && mysqldef --help
COPY db/schema.sql ./
CMD ["/bin/sh", "-c", "mysqldef -h\"${DB_HOST}\" -P\"${DB_PORT}\" -u\"${DB_USER}\" -p\"${DB_PASS}\" \"${DB_NAME}\" < ./schema.sql"]
