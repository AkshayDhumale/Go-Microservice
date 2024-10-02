# Golang Deployment on Kubernetes

## Demonstration of deploying a Go backend service.

This repository showcases the way to setting up, configuring, and deploying a scalable Go-based backend on Kubernetes. Explore the code to learn about deployment strategies, containerization with Docker and Kubernetes orchestration.

![Logo](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png)

## Documentation

- [Get started with Go](https://go.dev/doc/tutorial/getting-started)
- [docker](https://docs.docker.com/build/building/multi-stage/)
- [kubernetes](https://kubernetes.io/docs/concepts/workloads/)
- [wrk](https://github.com/wg/wrk)

## Prerequisites

- **Docker**: Installed and running.
- **Kubernetes**: A Kubernetes cluster set up and running. This can be a local setup using Minikube or a cloud-based Kubernetes service.
- **kubectl**: Command-line tool for interacting with your Kubernetes cluster.

## Deployment

**Steps :**

1. `DOCKER BUILD & TEST :`

   - Create docker image and test locally

2. `DEPLOY ON KUBERNETES`

   - use the created docker image to deploy the service on kubernetes

### 1. <u>DOCKER BUILD & TEST</u>

1. Navigate to the Source Directory :

```bash
  cd src/go-microservices
```

2. Build the Docker Image :

```bash
  docker build -t your-image-name:latest .
```

3. Tag the Docker Image with the dockerhub-repo-name :

```bash
  docker tag your-image-name:latest your-dockerhub-username/your-image-name:latest
```

**Run & Tests :**

```bash
  docker run -d -p 8080:80 your-dockerhub-username/your-image-name:latest
```

- access the url on `http://localhost:8080/`

### 2. <u>KUBERNETES DEPLOY & TEST</u>

1. Navigate to the kubernetes Directory :
   - **path** : `src/go-microservices/kubernetes`

```bash
  cd kubernetes
```

2. Apply the Deployment, Service, and HPA Files :

```bash
  kubectl apply -f go-deployment.yaml
  kubectl apply -f go-service.yaml
  kubectl apply -f go-hpa.yaml
```

3. Use Port Forward to Access the API :

```bash
  kubectl port-forward service/go-svc 8080:80
```

- access the url on `http://localhost:8080/`

## Testing HPA (Horizontal Pod Autoscaler)

### Prerequisites

1. **Enable Metric Server**:
   Ensure the metric server is enabled in your Kubernetes cluster:

- for **minikube**

```bash
  minikube addons enable metrics-server
```

2. Deploy a Simple Alpine Pod for Stress Testing:

```bash
  kubectl apply -f stress-pod.yaml
```

3. Install `wrk` in the stress Pod:

-

```bash
  kubectl exec -it stress-alpine -- sh
  apk update
  apk add wrk
```

4. Induce Stress by Applying Load on the Server URL:

   **_wrk - a HTTP benchmarking tool_**

   > wrk is a modern HTTP benchmarking tool capable of generating significant load when run on a single multi-core CPU. It combines a multithreaded design with scalable event notification systems such as epoll and kqueue.

- wrk -t10 -c20 -d100s http://your-service-url

```bash
  wrk -t10 -c20 -d100s http://go-svc/health
```

5. Monitor Pod Autoscaling:

```bash
  kubectl top pods
  kubectl get hpa
  kubectl get pods -w
```
