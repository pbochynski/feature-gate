# Build installer image

```
docker build -t kyma-installer:latest -f kyma-installer/kyma.Dockerfile .
docker tag kyma-installer:latest pbochynski/kyma-installer:latest
docker push  pbochynski/kyma-installer:latest
```

# Install installer

```
kubectl create namespace kyma-installer
kubectl apply -f installation/installer.yaml
kubectl apply -f installation/installer-config.yaml

# Install knative version
kubectl apply -f installation/installer-knative.yaml

# Or legacy version
kubectl apply -f installation/installer-legacy.yaml
```

# Start installation
```
kubectl label installation/kyma-installation action=install
```

