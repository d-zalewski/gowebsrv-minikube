# Simple golang web server running on minikube

Mini project showing how to run very basic golang webserver in Kubernetes (minikube)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

# Requirements

Minikube requires that VT-x/AMD-v virtualization is enabled in BIOS. To check that this is enabled on OSX / macOS run:

sysctl -a | grep machdep.cpu.features | grep VMX

If there's output, you're good!

# Prerequisites

- kubectl
- docker (for Mac)
- minikube
- virtualbox

```
brew update && brew install kubectl && brew cask install docker minikube virtualbox
```

# Start
```
minikube start
```

```
    This can take a while, expected output:
    Starting local Kubernetes v1.13.2 cluster...
    Starting VM...
    Getting VM IP address...
    Moving files into cluster...
    Setting up certs...
    Connecting to cluster...
    Setting up kubeconfig...
    Stopping extra container runtimes...
    Starting cluster components...
    Verifying kubelet health ...
    Verifying apiserver health ...
    Kubectl is now configured to use the cluster.
    Loading cached images from config file.
```

# Check k8s
```
kubectl get nodes
```  
Should output something like:

```
NAME       STATUS   ROLES    AGE   VERSION
minikube   Ready    master   88s   v1.13.2
```
    
# Use minikube's built-in docker daemon:
```
eval $(minikube docker-env)
```
Add this line to `.bash_profile` or `.zshrc` or ... if you want to use minikube's daemon by default (or if you do not want to set this every time you open a new terminal).

You can revert back to the host docker daemon by running:
```
eval $(docker-machine env -u)
```

# Enable ingress controller
```
minikube addons enable ingress
```

# Clone git repo with Dockerfile and k8s deployment
```
git clone git@github.com:d-zalewski/gowebsrv-minikube.git
cd gowebsrv-minikube
```

# Build go app image
```
docker build -t mygoapp:latest docker-image
```

# Deploy mygoapp (pods,service,ingress)
```
kubectl apply -f k8s/mygoapp.yaml
```

Should output something like:

```
deployment.apps/mygoapp-deployment created
service/mygoapp-service created
ingress.extensions/mygoapp-ingress created
```

```
kubectl get pods
```

```
NAME                                  READY   STATUS    RESTARTS   AGE
mygoapp-deployment-6465ffbbcd-9q27w   1/1     Running   0          10s
mygoapp-deployment-6465ffbbcd-ntqc7   1/1     Running   0          10s
mygoapp-deployment-6465ffbbcd-xh5c7   1/1     Running   0          10s
```

```
kubectl get service
```

```
NAME              TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
kubernetes        ClusterIP   10.96.0.1      <none>        443/TCP   22m
mygoapp-service   ClusterIP   10.97.17.131   <none>        80/TCP    45s
```

```
kubectl get ingress
```

```
NAME              HOSTS   ADDRESS     PORTS   AGE
mygoapp-ingress   *       10.0.2.15   80      80s
```

# Browse the go app
```
curl http://$(minikube ip)
```

You should see something like (container IDs will be changing on every few requests as the service is load balanced):

Hello World! Running on container ID mygoapp-deployment-6465ffbbcd-ntqc7%

