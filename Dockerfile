FROM centos:7
WORKDIR /app
COPY /apisix-admin /app/
ENTRYPOINT ["/app/apisix-admin"]
