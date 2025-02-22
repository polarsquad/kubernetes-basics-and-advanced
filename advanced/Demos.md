# Connecting to Kubernetes

```sh
kubectl cluster-info
```

```sh
kubectl config current-context
kubectl config get-contexts
kubectl config use-context <context name>
```

```sh
kubectl config set-context --current --namespace=<namespace>
```

# Using a service account from a pod

```sh
kubectl apply -f advanced/day1/demo1_0.yaml
kubectl exec -it pods/sa-demo -- /bin/bash
printenv
curl -H "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \-ik https://172.20.0.1:443/api/v1/namespaces/tr-tuomas/pods
```



# Pod Security Policies and Standards

```sh
kubectl label --dry-run=server --overwrite ns <your own namespace> pod-security.kubernetes.io/warn=restricted
```

# Persistence

```sh
kubectl patch pv <your-pv-name> -p '{"spec":{"persistentVolumeReclaimPolicy":"Retain"}}'
```

or

```sh
kubectl edit pv <your-pv-name>
```

Manually delete the pv

# Upgrades

```sh
kubectl scale deployment my-nginx --replicas=4
kubectl set resources deployment my-nginx --limits=cpu=100m,memory=64Mi --requests=cpu=100m,memory=64Mi
watch kubectl get all
kubectl set image deployment/my-nginx nginx=nginx:1.23.4
```

