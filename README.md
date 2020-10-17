# Go-Microservice-Starter

Go-Microservice-Starter is a starter project for developing microservices in GO using Micro + Echo + Nats + Kubernetes + Skaffold + Docker + Nginx

Use of Micro provides easy switch of Broker from Nats to Kafka, RabbitMQ etc. Similarly Echo can be switched to Gin

### Setup

- install skaffold
- install kubernetes/minikube
- install Docker
- install MongoDB
- install Nats

It containerizes the services and runs it in kubernetes. It creates a http application server and an internal grpc server.

- Each service has a http application server and 4 microservice components
  - GRPC Clients
  - GRPC Handlers
  - Publishers
  - Subscribers

Also possible to use Nginx as a reverse proxy using

```
    nginx -c nginx.dev.conf
    nginx -c nginx.prod.conf
```

### Dev Mode

Watches code changes, containerizes and deploys the service to kubernetes cluster

```
    skaffold dev
```

### Prod Mode

Builds the release build of the service and deploys the service to kubernetes cluster

```
    skaffold run
```
