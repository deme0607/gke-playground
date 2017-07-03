# gke-playground

## API

### build container

```
$ gcloud container builds submit --config=api/build.yml api
```

### create Deployment

```
$ kubectl create -f api/deployment.yml
```

### create Service

```
$ kubectl create -f api/service.yml
```

### create Ingress

```
$ kubectl create -f api/ingress.yml
```
