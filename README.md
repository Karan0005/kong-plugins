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
