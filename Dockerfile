##########################################
## Kong Plugins BuildUp
##########################################

FROM kong/go-plugin-tool:2.0.4-alpine-latest AS builder

#Setting Up kong-plugins Directory
RUN mkdir -p /tmp/kong-plugins/
RUN cd /tmp/kong-plugins/ && go mod init kong-plugins

#Placement of Custom Developed Kong Plugins into kong-plugins Directory
COPY /auth/plugin.go /tmp/kong-plugins/auth-plugin.go
COPY /correlation/plugin.go /tmp/kong-plugins/correlation-plugin.go

#Building Up Custom Developed Kong Plugins into kong-plugins Directory
RUN cd /tmp/kong-plugins/ && go get -d -v github.com/Kong/go-pluginserver
RUN cd /tmp/kong-plugins/ && go build github.com/Kong/go-pluginserver
RUN cd /tmp/kong-plugins/ && go build -buildmode plugin auth-plugin.go
RUN cd /tmp/kong-plugins/ && go build -buildmode plugin correlation-plugin.go

##########################################
## Release Kong Image
##########################################

FROM kong:2.0.4-alpine

#Placement of Build files of Custom Developed Kong Plugins
RUN mkdir /tmp/kong-plugins
COPY --from=builder /tmp/kong-plugins/go-pluginserver /usr/local/bin/go-pluginserver
COPY --from=builder /tmp/kong-plugins/auth-plugin.so /tmp/kong-plugins/auth-plugin.so
COPY --from=builder /tmp/kong-plugins/correlation-plugin.so /tmp/kong-plugins/correlation-plugin.so

#Placement of Supporting files of Custom Developed Kong Plugins
COPY config.yml /tmp/config.yml
COPY /auth/public.key /tmp/public.key

#Choose root User for Below Operations with 777 permission
USER root
RUN chmod -R 777 /tmp

#Copy Go files
RUN /usr/local/bin/go-pluginserver -version
RUN cd /tmp/kong-plugins
RUN /usr/local/bin/go-pluginserver -dump-plugin-info kong-plugins

#Set kong User as Default
USER kong
