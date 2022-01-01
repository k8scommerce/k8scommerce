FROM governmentpaas/psql:latest

# https://github.com/pressly/goose/releases/download/v2.7.0-rc4/goose-linux386
ARG GOOSE_BINARY_URL=https://github.com/pressly/goose/releases/download/v3.4.1/goose_linux_x86_64

RUN wget -O /bin/goose ${GOOSE_BINARY_URL} && chmod 755 /bin/goose

COPY ./wait-for-postgres.sh /
COPY ./run-migrations.sh /
COPY ./migrations/ ./migrations/

ENV PATH="/bin:/sbin:${PATH}"

CMD /bin/sh
