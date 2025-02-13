# Mewtwo-Docker

## Run the App in the Container

1. Create a Dockerfile.
2. Create a Docker Compose File.
3. Build the Docker image.
``` sh
docker build -t hello-gin .
```
4. Run the container.
``` sh
docker-compose up -d

# docker run -p 8080:8080 hello-gin
```
5. Access *http://localhost:8080/*

### Share the Container

1. Share the docker image via Docker Hub
2. Share the Image as a .tar file.
3. Share **docker-compose.yml** file.

## Move to Kubernetes

### Define the Kubernetes YAML Manifests

- postgres-deployment.yaml
- postgres-service.yaml
- gin-deployment.yaml
- gin-service.yaml

### Deploy to Kubernetes

1. Apply the YAML files:
``` sh
kubectl apply -f postgres-deployment.yaml
kubectl apply -f postgres-service.yaml
kubectl apply -f gin-deployment.yaml
kubectl apply -f gin-service.yaml
```

2. Check the running Pods
``` sh
kubectl get pods
```

3. Find the app's exposed port
``` sh
kubectl get service hello-gin
```

## Q&A

### Difference Between ``docker run`` and ``docker-compose up``?

- The ``docker run`` command is used to run a single container, ``docker-compose`` is used to manage multiple containers with a **docker-compose.yml** file.
- ``docker run``:
  - When you only need to run a single container.
  - When you don’t need complex configurations (volumes, multiple services, networks, etc.).
- ``docker-compose``:
  - When you have multiple containers (e.g., web + database).
  - When you want to manage configuration easily.
  - When you need to restart or scale services effortlessly.