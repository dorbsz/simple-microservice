# simple-microservice with golang

## Docker
### To create image
sudo docker build -t simple-microservice .
### To run the container detached
sudo docker run -d -p 80:80 simple-microservice:latest

### to view images
sudo docker images

### to view running containers
docker container ls

## Kubernetes
### To create the objects, we always start with the pod
kubectl apply -f config/kubernetes/simple-microservice-pod.yaml

### Apply the NodePort
kubectl apply -f config/kubernetes/simple-microservice-port.yaml
