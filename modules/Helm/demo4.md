# Using subcharts

Let's add  a Redis to the chart we created.

We'll need to add Redis as a dependency. Let's search the Bitnami repo for it:
```sh
helm search repo redis
```

Then add this block to your Chart.yaml:
```yaml
dependencies:
  - name: redis
    version: 17.11.3
    repository: https://charts.bitnami.com/bitnami 
```

Then run `helm dependency update <your chart name>`. Helm will pull the chart as a tgz into charts/

Now you'll need to use the settings we figured out earlier to configure Redis to run in standalon mode
with no authentication, since the app is a bit daft that way.

Add the following to your values.yaml:
```yaml
redis:
  defaultStorageClass: ebs-sc
  auth:
    enabled: false
  architecture: standalone
  master:
    containerPorts:
      redis: 6379
    resources:
      limits:
        cpu: 100m
        memory: 64Mi
      requests:
        cpu: 100m
        memory: 64Mi
```

Check that Redis is installed as well. Next we'll need to configure our app to use Redis and we'll do that via Environment variables.

Add the following block to your Deployment.yaml in correct place and indent:
```yaml
          volumeMounts:
            {{- toYaml . | nindent 12 }}
          {{- end }}
          env:
            - name: REDIS
              value: go-app-redis-headless
            - name: REDIS_PORT
              value: "6379"          
```
You can use go-app-redis-master as well. Fo-app is the name of your release, so change it accordingly.

Test the installation first. If the endpoint /helloworld is working as intended you could "helmify" the values next. This way your service is reachable without you having to figure out what the name of your release is.

```yaml
          env:
            - name: REDIS
              value: {{ .Release.Name }}-redis-headless
            - name: REDIS_PORT
              value: "{{ .Values.redis.master.containerPorts.redis }}"
```

Then let's add a flag for Redis to enable and disable it. We'll need to edit Chart.yaml as well.

Add this to values.yaml:
```yaml
charts:
  redis:
    enabled: false
```

Upgrade and see what happens to Redis.

Change the value to `true` and re-upgrade.