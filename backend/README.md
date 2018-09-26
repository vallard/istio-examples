# Sample Stuff for Istio

```
export CGO_ENABLED=0
export GOOS=linux
go build -ldflags '-s' -o backend
docker build -t vallard/istio-backend . 
docker run -d -p 3000:3000 -e VERSION=v1 vallard/istio-backend
```

Now curl the porgram:

```
curl localhost:3000/hello
```

It will now return the version
