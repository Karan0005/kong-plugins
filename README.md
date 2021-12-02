# Prerequisites:

1. Windows Docker.

# Build Command:

docker build -t kong-plugins .

# Run Command:

docker run -ti --rm --name kong-plugins -e "KONG_DATABASE=off" -e "KONG_GO_PLUGINS_DIR=/tmp/kong-plugins" -e "KONG_DECLARATIVE_CONFIG=/tmp/config.yml" -e "KONG_PLUGINS=auth-plugin,correlation-plugin" -e "KONG_PROXY_LISTEN=0.0.0.0:8000" -p 8000:8000 kong-plugins

# Other Useful Docker Commands:

1. docker ps
2. docker ps -a
3. docker system prune
4. docker image prune
5. docker inspect [YOUR_CONTAINER_ID]

# Method Definitions

1. Certificate: Executed during the SSL certificate serving phase of the SSL handshake.
2. Rewrite: Executed for every request upon its reception from a client as a rewrite phase handler
3. Access: Executed for every request from a client and before it is being proxied to the upstream service.
   Response: Executed after the whole response has been received from the upstream service, but before sending any part of it to the client.
4. Preread: Executed once for every connection
5. Log: Executed once for each connection after it has been closed.

# References

1. https://levelup.gitconnected.com/kong-custom-plugin-development-using-go-abab906b89b4
2. https://master--kongdocs.netlify.app/gateway-oss/2.1.x/go/
3. https://pkg.go.dev/github.com/Kong/go-pdk@v0.7.1/service/request#pkg-overview
