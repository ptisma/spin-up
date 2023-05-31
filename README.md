# spin-up
Some ideas: full on native k8 app with ingress, domain and ssl cert. Web server gets requests, triggers another services which spins/builds VM (locally or on some cloud vendor). Implement metrics, logging, k8 patterns and so on. Tech to try out: Prometheus, Grafana, Packer, maybe some service mesh etc.



## api-server
API server has been built following the service pattern. We want to achieve graceful shutdown, dependency injection, circuit breaking and more. Some of the design patterns used: Strategy, Singleton, Composite, Decorator.
Docs:
https://irahardianto.github.io/service-pattern-go/

https://gabrieltanner.org/blog/collecting-prometheus-metrics-in-golang/

https://www.robustperception.io/configuring-prometheus-with-docker/


Authentication is implemented using Auth0.

curl --request GET url http:localhost:8080/health --header 'authorization: Bearer <TOKEN>'



https://auth0.com/docs/quickstart/backend/golang/interactive

## packer

Provide a git repo URL, git repo creds and a path to build file.


https://learn.microsoft.com/en-us/azure/virtual-machines/windows/build-image-with-packer


## mongodb

We use the mongodb database to docker image job data


