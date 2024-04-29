# sample-scheduler

## 编译插件
```
go build -o bin/sample-scheduler main.go
```

## 制作镜像
```
docker build -f Dockerfile . -t sample-scheduler:v0.0.3
```

## 部署插件
```
kubectl apply -f sample-scheduler.yaml
```

## 部署应用
```
kubectl apply -f nginx-deployment-sample-scheduler.yaml
```

## 查看日志
```
kubectl -n kube-system logs -f sample-scheduler-5ff99554cf-4ztqd
```

## 节点标签
```
kubectl label node node02 sample=true
```
