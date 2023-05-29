# Persistence and configuration

## Exercise 1

Create a PVC using the default StorageClass which is gp2 for our cluster.

Create a deployment that uses the PVC to mount a directory inside the pods. 

You can also try rerunning the StatefulSet from day1/demo6_1.yaml and see how it handles PVCs.

## Exercise 2 - Services 

Create a PVC and mount it inside an Nginx deployment.

Create and edit an index.html in /usr/share/nginx/html and scale your deployment down to zero. Check that the PVC is retained. Scale the deployment back up again and check that your changes are in place.

Now create a service that allows you to access your pods.

Test the functionality with `kubectl port-forward`.

## Exercise 3 - Ingresses

We're using an Nginx ingress controller. 

You'll need to create several deployments and services in addition to the ingress.

First, you'll need a Redis deployment and a Redis service. Create a PVC for the data and mount it in `/data` inside the Redis pod. This could be a statefulSet if you want, but a simple deployment will do for now.

Second, create a deployment and a service using the Go app in europe-north1-docker.pkg.dev/tuomaspal-sandbox/training/hello-http:0.0.2 as the image. The deployment needs two env variables called REDIS and REDIS_PORT. These need to correspond with your Redis deployments address and port. You can create and mount the settings via a configMap if you want. The Go app has two endpoints, `/`and `/helloworld`. We want to access both of these.

Third, create a simple Nginx deployment and a matching service.

Now you'll need to create an ingress for your apps. There are two apps we want to expose to the wider internet: Nginx and the Go app. Route nginx to the root of your hostname. Then route the Go app to some other path, like /go/. Check that all the endpoints work.

Use this as a starting point:
```yaml
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
            name: nginx
            port:
              number: 80
```

Important tip: https://kubernetes.github.io/ingress-nginx/examples/rewrite/

Once you're done, delete your deployments, services, ingress and the PVC.