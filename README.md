# Fullstack & DevOps Engineer Technical Task

A simple RESTful API built with Golang, containerized using Docker (multi-stage build), and deployed via CI/CD using GitHub Actions.

## 🚀 Features & Implementation

- **Lightweight Application:** Built with Go standard library (`net/http`)
- **Reliability:** Implements graceful shutdown to prevent dropped requests
- **Optimized Docker Image:** Multi-stage build for smaller, secure image
- **Automated CI/CD:** GitHub Actions builds & pushes image on every push to `main`
- **Cloud Deployment:** Deployed on PaaS with auto-scaling (2 replicas) and HTTPS

## 🛠️ How to Run Locally

1. Clone repository:
   ```bash
   git clone https://github.com/kikukafandi/devops-go-task.git
   cd devops-go-task
   ```

2. Run with Docker Compose:

   ```bash
   docker-compose up -d --build
   ```

3. Access endpoint:

   ```text
   http://localhost:8080/health
   ```

## 🌐 Live URL (Production)

[https://ymyptadllgfn.ap-southeast-1.clawcloudrun.com/](https://ymyptadllgfn.ap-southeast-1.clawcloudrun.com/)
