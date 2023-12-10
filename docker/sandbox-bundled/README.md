# Nebula Deployment Sandbox

First make images
```
ytong@Yees-MBP:~/go/src/github.com/nebulaclouds/nebula/docker/sandbox-ultra [nebula-sandbox] (cicd-sandbox-lite) $ make images
```

then build the k3s image.
```
ytong@Yees-MBP:~/go/src/github.com/nebulaclouds/nebula/docker/sandbox-ultra [] (cicd-sandbox-lite) $ docker buildx build --file images/dockerfiles/k3s/Dockerfile --platform linux/arm64,linux/amd64 --push --tag ghcr.io/nebulaclouds/nebula-sandbox-lite:ultra7 .
```
