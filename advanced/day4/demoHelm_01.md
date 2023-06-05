# Helm repos

Add the Bitnami repo

```sh
helm repo add bitnami https://charts.bitnami.com/bitnami
```

Search for Redis
```sh
helm search repo redis
```

Look at the default values:
```sh
helm show values bitnami/redis
```

You can download a remote chart:
```sh
helm pull bitnami/redis
```

Push is also available.