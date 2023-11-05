# HPS Internship Project

This project is part of my internship project at HPS. It consists of 4 services: frontend, backend, parser, and Redis for caching. The backend and parser services are written in Go and communicate with each other using gRPC. The client communicates directly with the backend through the frontend.

## DevOps

This project is focused on DevOps and is meant to be deployed on a Kubernetes cluster. The following tools and technologies are used for DevOps:

- Terraform with S3 backend for infrastructure provisioning (DigitalOcean) & Kubernetes cluster setup (Nginx Ingress Controller, Cert-Manager, ArgoCD and Prometheus Operator + Grafana)
- GitHub Actions for continuous integration
- Docker Hub for storing Docker images
- ArgoCD for continuous deployment
- Prometheus Operator + Grafana for monitoring

The project follows the GitOps pattern, which means that the entire deployment process is managed through Git and changes are automatically deployed to the cluster when changes are pushed to the Git repository.

## Getting Started

To get deploy this project, follow these steps:

1. Clone the repository
2. Jump into the `terraform` directory and run `./tf-apply.sh` to initialize Terraform, provision the infrastructure, setup the Kubernetes cluster, install the ArgoCD application and add this repository as a source for the application.
3. Assign the outputted IP address of the nginx-ingress-controller service to your prod and staging domain names on your DNS provider.
4. Using the generated `config` file, run this command to get the password for the ArgoCD admin user: `kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d`
5. Port forward the ArgoCD server to your local machine: `kubectl port-forward svc/argocd-server -n argocd 3000:443` and login to the ArgoCD dashboard using the username `admin` and the password from the previous step.
6. Sync the ArgoCD production application to deploy the project to the cluster.

## Usage

This project is using semantic versioning. The version number is stored in the kubernetes manifests file in the kubernetes/production directory. The version number is automatically incremented by the CI pipeline when a new tag is pushed to the repository.

The project is deployed to two environments: staging and production. The staging environment is deployed to the staging domain name, while the production environment is deployed to the production domain name.

The staging environment is automatically synced and deployed when a new commit is pushed to the `main` branch, while the production environment is not automatically synced and deployed when a new tag is pushed. This is done to prevent accidental deployments to the production environment.
