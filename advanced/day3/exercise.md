# Label matching

Create a deployment with nginx or some other image. Make sure that the app is running on a node with
type = test.

# Scheduling apps to tainted nodes

There is a tainted node in the cluster: p-10-0-1-214.eu-west-1.compute.internal. It has two sorts of taints turned on.

```
key1=value1:NoExecute
key1=value2:NoSchedule
```

Create a deployment or a pod that runs on this node.
