# CI/CD Pipeline Documentation

This directory contains the GitHub Actions workflow for continuous integration and deployment of the Todo App.

## Pipeline Overview

The CI/CD pipeline consists of five main stages:

1. **Lint**: Code quality checks using golangci-lint
2. **Test**: Running Go unit tests
3. **Build**: Building and pushing Docker images to DockerHub
4. **Update Helm**: Updating Helm chart with new image tag
5. **Deploy**: Triggering ArgoCD sync for deployment

## Pipeline Triggers

The pipeline runs on:
- Push to `main` branch
- Pull requests to `main` branch

## Required Secrets

The following secrets need to be configured in your GitHub repository:

1. `DOCKERHUB_USERNAME`: Your DockerHub username
2. `DOCKERHUB_TOKEN`: Your DockerHub access token
3. `ARGOCD_USERNAME`: ArgoCD admin username
4. `ARGOCD_PASSWORD`: ArgoCD admin password

## Environment Variables

Configure these in your repository settings:

1. `REGISTRY`: Docker registry (default: docker.io)
2. `IMAGE_NAME`: Your DockerHub image name (e.g., username/repo)
3. `ARGOCD_SERVER`: Your ArgoCD server URL (e.g., argocd.example.com)
4. `ARGOCD_APP_NAME`: Your ArgoCD application name (e.g., todo-app)

## Pipeline Stages

### 1. Lint Stage
- Uses golangci-lint for static code analysis
- Ensures code quality and style consistency
- Runs on every push and pull request

### 2. Test Stage
- Runs Go unit tests
- Ensures code functionality
- Runs after successful linting

### 3. Build Stage
- Builds Docker image using Buildx
- Pushes image to DockerHub
- Tags images with:
  - `latest`
  - Short commit SHA
- Runs after successful testing

### 4. Update Helm Stage
- Updates image repository and tag in values.yaml
- Commits and pushes changes to repository
- Triggers ArgoCD sync

### 5. Deploy Stage
- Installs ArgoCD CLI
- Logs in to ArgoCD server
- Triggers application sync
- Waits for sync completion
- Only runs on pushes to main branch

## Manual Deployment

To deploy manually:

```bash
# Update image in values.yaml
sed -i "s|repository: .*|repository: docker.io/username/repo|g" helm/todo-app/values.yaml
sed -i "s|tag: .*|tag: <commit-sha>|g" helm/todo-app/values.yaml

# Push changes
git add helm/todo-app/values.yaml
git commit -m "Update image to docker.io/username/repo:<commit-sha>"
git push

# Trigger ArgoCD sync
argocd app sync todo-app --force --prune
argocd app wait todo-app --health
```

## Troubleshooting

Common issues and solutions:

1. **Build Failures**:
   - Check Dockerfile syntax
   - Verify DockerHub credentials
   - Ensure proper permissions for DockerHub
   - Check DockerHub rate limits

2. **ArgoCD Sync Failures**:
   - Verify ArgoCD credentials
   - Check ArgoCD server accessibility
   - Review application health status
   - Check resource quotas and limits

3. **Test Failures**:
   - Run tests locally first
   - Check test coverage
   - Verify test dependencies

## Best Practices

1. Always run tests locally before pushing
2. Keep dependencies updated
3. Monitor pipeline execution times
4. Review security best practices for GitHub Actions
5. Regularly backup important configurations
6. Use ArgoCD's auto-sync feature carefully
7. Implement proper RBAC for ArgoCD
8. Monitor ArgoCD application health
9. Use DockerHub access tokens instead of passwords
10. Implement proper Docker image tagging strategy 