# Usre And Organization Management Of ContainerOps
![Crew](docs/crew.jpg "Crew")

Crew is a REST API service integrating user and organization management, GPG & SSH keys management, OpenID connect identity(OIDC), OAuth 2.0 provider and RABC management.

# containerops.conf 


# Runtime Conf File Example

```
runmode = dev

listenmode = https
httpscertfile = cert/containerops/containerops.crt
httpskeyfile = cert/containerops/containerops.key

[log]
filepath = log/backend.log
level = info

[database]
driver = mysql
uri = containerops:containerops@/containerops?charset=utf8&parseTime=True&loc=Asia%2FShanghai

[deployment]
domains = containerops.me
```
