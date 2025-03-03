# Packaging and publishing a package

Test your package with:
```sh
helm lint
```

Then you can compress it into a tgz:
```
helm package mynewchart --version 1.0.0
```

And push it into an artifact registry:
```sh
helm push mynewchart-1.0.0.tgz oci://europe-north1-docker.pkg.dev/tuomaspal-sandbox/training/  
```

You could then use this chart to install the app with:
```sh
helm install demo oci://europe-north1-docker.pkg.dev/tuomaspal-sandbox/training/mynewchart --version 1.0.0 -f ownValues.yaml
```

It's sometimes easier to use a values file instead of overriding everything on the command line, especially if Helm is expecting so-called tables of values.