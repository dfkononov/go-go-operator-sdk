FROM alpine:3.11.6

ENV OPERATOR=/usr/local/bin/go-go-operator-sdk \
    USER_UID=1001 \
    USER_NAME=go-go-operator-sdk

# install operator binary
COPY go-go-operator-sdk ${OPERATOR}

COPY build/bin /usr/local/bin

RUN  chmod u+x /usr/local/bin/user_setup && chmod ugo+x /usr/local/bin/entrypoint && /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
