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
curl -H "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \-ik https://172.20.0.1:443/api/v1/namespaces/tr-tuomas/pods
```



# Pod Security Policies and Standards

```sh
kubectl label --dry-run=server --overwrite ns <your own namespace> pod-security.kubernetes.io/warn=restricted
```