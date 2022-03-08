```
make golang-container-start
root@CSMachine:/app# go run main.go 
```

The golang intro: https://aristorinjuang.com/go-crash-course-part-1.html

The `kubectl port-forward svc/ingress-nginx-controller -n ingress-nginx 8080:80` part was inspired by this [article](https://banzaicloud.com/blog/kind-ingress/).