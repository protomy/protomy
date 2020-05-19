FROM alpine:3

COPY protomy /bin/protomy

ENTRYPOINT ["protomy"]