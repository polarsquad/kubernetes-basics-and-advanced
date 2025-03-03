# Install an app

Try installing Redis first by running:
```
helm install redis-test bitnami/redis --set auth.enabled=false --set architecture=standalone --dry-run
```

Then do a proper installation:
```
helm install redis-test bitnami/redis --set auth.enabled=false --set architecture=standalone --set defaultStorageClass=ebs-sc
```

Inspect the release and figure out why it didn't work. You can use `kubectl` tools here. 

You can delete the installation with
```sh
helm delete redis-test
```

Install Redis with the following parameters
```sh
helm install redis-test bitnami/redis --set auth.enabled=false --set architecture=standalone --set master.resources.limits.cpu=100m --set master.resources.limits.memory=64Mi --set master.resources.requests.cpu=100m --set master.resources.requests.memory=64Mi --set defaultStorageClass=ebs-sc
```

Or create an ownValues.yaml or something and insert the values there like this:
```yaml
auth.enabled: false
architecture: standalone
defaultStorageClass: ebs-sc
master:
  resources:
    limits:
      cpu: 100m
      memory: 64Mi
    request:
      cpu: 100m
      memory: 64Mi
```

Then install Redis with:
```sh
helm install redis-test bitnami/redis -f ownValues.yaml
```

You can also run:
```sh
helm upgrade -i redis-test bitnami/redis --set auth.enabled=false --set architecture=standalone --set master.resources.limits.cpu=100m --set master.resources.limits.memory=64Mi --set master.resources.requests.cpu=100m --set master.resources.requests.memory=64Mi --set defaultStorageClass=ebs-sc
```

# Upgrading

Let's increase the CPU limit a bit. You can either use the full command used previously and just edit one bit _or_
you can use `--reuse-values`.

```sh
helm upgrade redis-test bitnami/redis --set auth.enabled=false --set architecture=standalone --set master.resources.limits.cpu=150m --set master.resources.limits.memory=64Mi --set master.resources.requests.cpu=100m --set master.resources.requests.memory=64Mi --set defaultStorageClass=ebs-sc
```

```sh
helm upgrade redis-test bitnami/redis --reuse-values --set master.resources.limits.cpu=150m
```

# Rollback

Check history with:
```sh
helm history redis-test
```

Then do the rollback to the previous version:
```sh
helm rollback redis-test 1
```
The rollback is stored in the history as well. Check:
```sh
helm history redis-test
```

# Completion

When your done, delete the release.

```sh
helm delete redis-test
```

Remember to delete the PVC as well!