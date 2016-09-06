# Usre And Organization Management Of ContainerOps
![Crew](docs/crew.jpg "Crew")

# containerops.conf 


# Runtime Conf File Example

```
runmode = dev

listenmode = https
httpscertfile = cert/containerops/containerops.crt
httpskeyfile = cert/containerops/containerops.key

[log]
filepath = log/containerops-log

[db]
uri = localhost:6379
db = 8
```
