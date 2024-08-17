## Link
https://github.com/kubernetes/client-go/tree/master/examples

### 5 SimpleClient
```
GOOS=linux go build -o ./app .
docker build -t in-cluster -t ysyoo/hello:latest .

```

```
docker images
REPOSITORY     TAG         IMAGE ID       CREATED          SIZE
in-cluster     latest      b885ecd39fd6   26 seconds ago   117MB
```

```
 kubectl apply -f ./
```