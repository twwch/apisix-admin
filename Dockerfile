FROM scratch
WORKDIR /app
COPY apisix-admin /app/
ENTRYPOINT ["/app/apisix-admin"]
