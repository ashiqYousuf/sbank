# sbank

```
Migration command:-

migrate create -ext sql -dir db/migration -seq <migration_name>
```

```
Install evans (go install github.com/ktr0731/evans@latest)

Update PATH: 

echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc

go install package (goes to ~/go/bin/)
```

Note:- gRPC interceptor is a function that gets called for every request before its sent to the gRPC handler (Doesn't work for HTTP Requests, need to seperately add that for gRPC Gateway)
