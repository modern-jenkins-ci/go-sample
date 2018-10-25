FROM scratch

ARG GIT_COMMIT=unspecified
LABEL git_commit=$GIT_COMMIT

ADD go-sample /

EXPOSE 8080
ENTRYPOINT ["/go-sample"]