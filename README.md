# Project Metis 🛠️

[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](#)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-lightgrey)](./LICENSE)

Project Metis is a miniature, production-grade Internal Developer Platform (IDP) built to provide a massively simplified, opinionated "paved road" for application developers. It automates the complexity of building, deploying, and operating services at scale.

## The Problem
In a large engineering organization, managing the lifecycle of hundreds of microservices is complex. Developers face significant cognitive overhead dealing with raw Kubernetes YAML, CI/CD pipelines, observability tooling, and security policies. Metis solves this by providing a single, cohesive CLI and platform that automates best practices, allowing developers to focus on writing business logic.

## Core Features ✨
* ✅ **Declarative Deployments:** Define services in a simple `metis.yaml` and deploy to Kubernetes with a single `metisctl deploy` command.
* ✅ **Automated CI Pipeline:** On every `git push` to the main branch, a GitHub Actions workflow automatically builds and pushes a uniquely tagged container image to the GitHub Container Registry.
* ✅ **Dynamic Image Updates:** Use the `metisctl deploy --image <tag>` flag to deploy a specific, versioned container image, enabling seamless rollouts and rollbacks.
* ✅ **Professional CLI:** Built with Cobra & Viper, featuring build-time versioning, dynamic configuration from files and environment variables, and graceful shutdown handling.

## Example Workflow
This platform provides a complete "code-to-cluster" experience.

1.  A developer makes a code change to the `hello-world` application.
2.  They `git push` the commit to the `main` branch.
3.  The GitHub Actions pipeline automatically triggers, building a new image and pushing it to GHCR with a unique tag (e.g., `ghcr.io/your-username/hello-world:ci-a1b2c3d`).
4.  The developer (or an automated script) runs the deploy command with the new tag:
    ```bash
    ./bin/metisctl deploy --image ghcr.io/your-username/hello-world:ci-a1b2c3d
    ```
5.  Kubernetes performs a rolling update, and the new version of the application is now live.

## Getting Started
### Prerequisites
* Go (version 1.21+)
* Docker
* `kind`
* `kubectl`

### Build & Run
1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/your-username/project-metis.git](https://github.com/your-username/project-metis.git)
    cd project-metis
    ```
2.  **Build the `metisctl` binary:**
    ```bash
    make build
    ```
3.  **Run a command:**
    ```bash
    ./bin/metisctl --help
    ```

## Project Roadmap 🗺️
This project is being developed over a nine-month intensive program designed to build a complete, end-to-end developer platform.

* ✅ **Months 1-2: The "Paved Road"**: Declarative Deployment (`metisctl`, CI/CD, K8s Object Generation)
* ⬜ **Months 3-4: The "Nervous System"**: Centralized Observability (OpenTelemetry)
* ⬜ **Months 5-6: The "Immune System"**: Proactive Reliability (Chaos Engineering)
* ⬜ **Months 7-8: The "Secure Supply Chain"**: Policy as Code (OPA/Gatekeeper)
* ⬜ **Month 9: The "Distinguished Engineer Pitch"**: Integration & Evangelism

## Technology Stack
* **Language:** Go
* **Orchestration:** Kubernetes
* **CI/CD:** GitHub Actions
* **CLI Framework:** Cobra & Viper
* **API Libraries:** `client-go`, `gopkg.in/yaml.v3`
