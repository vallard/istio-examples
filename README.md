# Istio Examples

This is a work in progress and tries to come close to what Kelsey Hightower does in [this presentation on Istio](https://www.youtube.com/watch?v=6BYq6hNhceI). I use this to demonstrate Istio on CCP behind.

The demo goes as follows:

1. Show the go code and show that it creates a microservices application
2. Show the yaml files and show what deployments, services, resources they create. 
3. Create a CCP cluster with Istio
4. Show the Istio resources that are created. 
5. Make the Istio grafana into a LoadBalancer or as part of an ingress controller. 
6. Deploy the resources of this environment. 
7. Show that when calling the resource we can get a both versions of the backend. 
8. Add the istio rule to make it so we load balance the traffic. 


