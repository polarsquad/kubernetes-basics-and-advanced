# Creating charts

Run:
```sh
helm create mynewchart
```

Take a look at the content. You can try installing it with the default options and see how it blows up.

# Editing and customizing the chart

Let's make some initial edits be able to actually run the chart in our cluster. You'll need to edit the resource settings to be able to run it in your namespace.

## Get the chart into a runnable state

Comment out the array first: {}

Then uncomment the actual default limits, these are just fine now.

If you want to enable ingress for example look for the relevant settings. Enable the ingress and change the host to match
your namespace + our cluster host.

Now you can try installing it. Do a dry run first.

Then do the actual installation. Try out the service and ingress. Once again, remember to be patient.

## Change the actual app image

Let's change nginx to our Go app. Use this as the image:
```
europe-north1-docker.pkg.dev/tuomaspal-sandbox/training/hello-http:0.0.2
```

Change image.repository and image.tag to match.

Try updating.

Our app is running in port 8080 by design so we'll need to make a couple of changes. We could add a new value and use that!

Let's add the following to values.yaml:
```yaml
application:
  port: 8080
```

Now we'll need to map the value to the correct variables in the manifests. Open Deployment.yaml and edit it. Then make the corresponding change in Service.yaml.

You'll need to use `{{ .Values.application.port }}`.

## Securing the app

Edit values.yaml and uncomment the securityContext part. It should look something like this:
```yaml
securityContext: 
  capabilities:
    drop:
    - ALL
  readOnlyRootFilesystem: true
  runAsNonRoot: true
  runAsUser: 1000
```

The default Nginx image requires filesystem access, but our Go app does not. We can remove all unnecessary permissions and force the app to run as nonroot (like it had any choice).

