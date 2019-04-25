# simple-microservice with golang

## on Docker version 18.06.1
### To create image
docker image build -t simple-app .
### To run
docker container run -p 80:80 simple-app:latest


## on Docker version 1.11.2
sudo docker build -t simple-app .
### to view images
sudo docker images
### to run the container
sudo docker run -p 80:80 simple-app:latest
