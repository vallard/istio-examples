```
export BACKEND=http://localhost:32771/hello
```

export CGO_ENABLED=0
export GOOS=linux
go build -ldflags '-s' 
docker build -t vallard/istio-frontend .
docker run -d -p 3001:3001 -e BACKEND=http://localhost:32771/hello vallard/istio-frontend
```

Now curl the fronend:

```
curl localhost:3001/
```

It will now return the version that the backend reports. 
