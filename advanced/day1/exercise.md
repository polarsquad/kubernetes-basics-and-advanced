# Deployments

These exercises assume that there is a resource quota in your namespace. You can use 1 gigabyte of memory and 1000m CPU as "hard limits" when thinking about resource usage. Once you're done you should delete the resources you've created. You can check your namespace's quota with `kubectl get resourcequota`.

An example quota could be set with a manifest like this:
```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: mem-cpu-demo
spec:
  hard:
    requests.cpu: "1"
    requests.memory: 1Gi
    limits.cpu: "1"
    limits.memory: 1Gi
```

The cluster admins could then set the quota with `kubectl apply -f my-quota.yaml -n <namespace>`

You can use the demos as starting points but pay attention to the labels.

The app code is under app/ if you want to laugh at my Go skills. 

# First Exercise

Create a deployment for europe-north1-docker.pkg.dev/tuomaspal-sandbox/training/hello-http:0.0.1

It uses port 8080. We want three replicas. 

Also create liveness and readiness probes. You can use `/` as the path.

Use whatever labeling you feel is useful here, but also add a `version: v1` to `spec.metadata.labels`.

# Second Exercise

Now create a new deployment using europe-north1-docker.pkg.dev/tuomaspal-sandbox/training/hello-http:0.0.2 with three replicas

Remember the quota limitations and health checks for this one as well. Add a label `version: v2` like in the first exercise.

This new version of the app adds Redis capabilities, but don't worry about that for now. It also adds a new endpoint /helloworld, which depends on Redis. Requests to that endpoint will fail, but all the relevant info is actually provided by the root of the app.

Check that you have two different deployments running.

Create a service that routes traffic to both deployments. Use the service in demo4_2.yaml as a starting point.

You can then create an ingress for the app using this as a basis:
```yaml
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: demo-ingress
  labels:
    app: demo
spec:
  rules:
  - host: <your namespace here>.k8s.polarsquad.training
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: <your service name here>
            port:
              number: 80
```

Since the ingress routes traffic to port 80 of the service, you'll need to make sure that the service provides port 80, but routes the traffic to the pods' ports. Make sure to use your namespace as part of the host.

Now test that you're getting a response from the ingress with a browser, curl or postman. Take note of the message.

Now scale down the number of replicas in v1 of your deployment. See what happens. Scale it down to zero.

# Third Exercise

Try to build a green/blue deployment based on these two earlier examples. Use the demo as guidance.
