install gin

```
go get github.com/gin-gonic/gin
```

build image
```
docker build -tag app-runner-test:{version} .
```

run container
```
docker run -p 8080:8080 app-runner-test:{version}
```
