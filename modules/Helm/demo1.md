# Helm repos

Add the Bitnami repo
```sh
helm repo add bitnami https://charts.bitnami.com/bitnami
```

List registered repo's
```sh
helm repo list
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

Extract pulled chart (adjust version number)
```sh
tar xvfz redis-20.10.0.tgz
```

Push is also available.