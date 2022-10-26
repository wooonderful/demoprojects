- 修改 pod 的 nodeName 为期待的节点名
- `kubectl apply -f dlv-pod.yaml`
- `kubectl exec -it ssh-privileged-deployment -- bash`
- `ps -ef | grep manager | grep 8080 | awk '{print $2}'`
- `dlv attach $(ps -ef | grep manager | grep 8080 | awk '{print $2}') --listen=:2345`

**用完记得删除对应pod**