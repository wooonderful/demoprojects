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

## 说明
其实这种方式调试不太好用， 远程代码和本地代码不能实时同步， 需要手动更新远端代码。
并且由于远端是直接 sttach 到 pid 上的， 直接在 vscode 等 debugger 上打断点是没用的。

建议最好还是使用 dlv debug 的方式去调试， 另外简化的方式是直接远端目标容器用 sshd 容器替换， 
比如在原本的 Dockerfile 上加上 sshd 功能， 然后使用 goland 的远程调试工具， 这种效果是最好的。