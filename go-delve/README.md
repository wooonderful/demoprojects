# golang 程序通过 dlv 特权容器远程调试

## remote
- 修改 pod 的 nodeName 为期待的节点名
- `kubectl apply -f dlv-pod.yaml`
- `kubectl exec -it ssh-privileged-deployment -- bash`
- `ps -ef | grep manager | grep 8080 | awk '{print $2}'`
- `dlv attach $(ps -ef | grep manager | grep 8080 | awk '{print $2}') --listen=:2345 --accept-multiclient --api-version=2  --log=true --log-output=debugger,debuglineerr,gdbwire,lldbout,rpc --headless=true`


## local
- `dlv connect $(minikube ip):32345`

**用完记得删除对应pod**

**退出 dlv 时可能会杀掉远端程序**

**远程 k8s 探针可能会在断点时重启目标容器**