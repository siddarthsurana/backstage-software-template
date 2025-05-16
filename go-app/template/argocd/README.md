# ArgoCD Integration with Todo App Helm Chart

This directory contains the ArgoCD configuration for deploying the Todo App using Helm.

## Prerequisites

1. Kubernetes cluster with ArgoCD installed
2. Helm chart repository accessible to ArgoCD
3. Proper RBAC permissions for ArgoCD

## Setup Instructions

### Standard Installation

1. Install ArgoCD (if not already installed):
```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

### Lightweight Installation (Development/Testing)

For development or testing environments, you can use the lightweight configuration:

1. Install ArgoCD with lightweight configuration:
```bash
kubectl create namespace argocd
helm repo add argo https://argoproj.github.io/argo-helm
helm install argocd argo/argo-cd -n argocd -f argocd/values-lightweight.yaml
```

2. Get the ArgoCD admin password:
```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

3. Port-forward to access ArgoCD UI:
```bash
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

4. Access ArgoCD UI at `https://localhost:8080` (username: admin, password: from step 2)

5. Apply the ArgoCD Application manifest:
```bash
kubectl apply -f argocd/application.yaml
```

## Lightweight Configuration Features

The lightweight configuration (`values-lightweight.yaml`) includes:

- Reduced resource requests and limits
- Disabled unnecessary components (Dex, Notifications, ApplicationSet)
- Minimal RBAC configuration
- Disabled persistence and HA features
- Disabled metrics collection
- Simplified security settings

This configuration is suitable for:
- Development environments
- Testing setups
- Local Kubernetes clusters
- Resource-constrained environments

## Configuration

The `application.yaml` file contains the following key configurations:

- **Source**: Points to your Helm chart repository
- **Destination**: Deploys to the default namespace
- **Sync Policy**: 
  - Automated sync enabled
  - Prune enabled (removes resources not in the chart)
  - Self-heal enabled (automatically corrects drift)
  - Creates namespace if it doesn't exist

## Customizing Values

To customize the Helm values for different environments:

1. Create environment-specific value files (e.g., `values-dev.yaml`, `values-prod.yaml`)
2. Update the `valueFiles` section in `application.yaml` to point to the desired values file

## Monitoring

- Access the ArgoCD UI to monitor the deployment status
- Check application health and sync status
- View resource tree and events

## Troubleshooting

Common issues and solutions:

1. **Sync Failed**: Check the ArgoCD logs and events for specific error messages
2. **Authentication Issues**: Verify repository access credentials
3. **Resource Conflicts**: Ensure no manual changes are made to resources managed by ArgoCD

## Best Practices

1. Use GitOps workflow for all changes
2. Keep Helm values in version control
3. Use separate ArgoCD Applications for different environments
4. Regularly backup ArgoCD configuration 