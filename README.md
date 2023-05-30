# spin-up
Some ideas: full on native k8 app with ingress, domain and ssl cert. Web server gets requests, triggers another services which spins/builds VM (locally or on some cloud vendor). Implement metrics, logging, k8 patterns and so on. Tech to try out: Prometheus, Grafana, Packer, maybe some service mesh etc.



## api-server
API server has been built following the service pattern. We want to achieve graceful shutdown, dependency injection, circuit breaking and more. Some of the design patterns used: Strategy, Singleton, Composite, Decorator.
Docs:
https://irahardianto.github.io/service-pattern-go/

https://gabrieltanner.org/blog/collecting-prometheus-metrics-in-golang/

https://www.robustperception.io/configuring-prometheus-with-docker/


Authentication is implemented using Auth0.

curl --request GET url http:localhost:8080/health --header 'authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlJhME51NEtJWHJMUmNuZ2d0ZFdBaSJ9.eyJpc3MiOiJodHRwczovL2Rldi1hMHR5YWFoY3ZndnN6YnFqLnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJNWThvU0ExVk1lYkN2bGtMcTRLaFRPbXZac0ZJMVkzeUBjbGllbnRzIiwiYXVkIjoiaHR0cHM6Ly9kZXYtYTB0eWFhaGN2Z3ZzemJxai51cy5hdXRoMC5jb20vYXBpL3YyLyIsImlhdCI6MTY4NTI4MTMwMiwiZXhwIjoxNjg1MzY3NzAyLCJhenAiOiJNWThvU0ExVk1lYkN2bGtMcTRLaFRPbXZac0ZJMVkzeSIsImd0eSI6ImNsaWVudC1jcmVkZW50aWFscyJ9.O65Sx5O41_3o6LUAKhTptAUqSKx-lZts8wl5GgP9Kfv9KQEjwv_ZJES3cUXruEsfwrEs3hAGzz_sRftjOTMeOYFJLKYJ0FusdKJ8BA5Ll2kouTiG3W5Wfcwbu8XPA4b6kmof-IOizLm_h5KwdKDs1OgWLfLzqk9xvos3FsY6Jog-toyFGfsXRMie5ZG7sqmTvyeRzdwj02aVjbcvQLEu9M8V5ZxQuR-54IYNs3g7lN64Oi-I79s-rWN1m_g6VpdeQ3W0ehHwaxanYjbLJq85VqGEBzfJtdKyVa9Gz3DwPCNYHKPqJYpveecKifDjw0ikV7rtyDUNcULrkLtF71jgXQ'



https://auth0.com/docs/quickstart/backend/golang/interactive

## packer

Provide a git repo URL, git repo creds and a path to build file.


https://learn.microsoft.com/en-us/azure/virtual-machines/windows/build-image-with-packer


## mongodb

We use the mongodb database to docker image job data


