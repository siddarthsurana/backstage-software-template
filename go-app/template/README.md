# Simple Todo API

A simple RESTful API built with Go and Gin framework, similar to Flask in Python. This API provides basic CRUD operations for managing todo items.

## Prerequisites

- Go 1.21 or higher
- Git
- Docker (optional)
- Kind (Kubernetes in Docker) (optional)

## Installation

### Local Installation

1. Clone the repository:
```bash
git clone <your-repo-url>
cd go-app
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Docker Installation

1. Build the Docker image:
```bash
docker build -t go-todo-api .
```

2. Run the container:
```bash
docker run -p 8080:8080 go-todo-api
```

The server will start on `http://localhost:8080`

### Kubernetes (Kind) Installation with Ingress

1. Install Kind if you haven't already:
```bash
brew install kind  # For macOS
```

2. Create a Kind cluster with Ingress support:
```bash
kind create cluster --name todo-cluster --config kind-config.yaml
```

3. Deploy the NGINX Ingress Controller:
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

4. Wait for the Ingress controller to be ready:
```bash
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```

5. Build and load the Docker image into Kind:
```bash
docker build -t go-todo-api .
kind load docker-image go-todo-api:latest --name todo-cluster
```

6. Deploy the application and Ingress:
```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/ingress.yaml
```

7. Add the host to your hosts file:
```bash
echo "127.0.0.1 todo.local" | sudo tee -a /etc/hosts
```

The API will be available at:
- `http://todo.local:30080/health`
- `http://todo.local:30080/todos`
- `http://todo.local:30080/todos/1`

### Backstage Integration

This project is integrated with Backstage, a developer portal that provides a unified interface for managing software components. The integration includes:

1. **Component Registration**
   - The project is registered in Backstage as a software component
   - Metadata includes repository information, ownership, and tech stack details

2. **CI/CD Integration**
   - GitHub Actions workflows are visible in Backstage
   - Build and deployment status can be monitored
   - ArgoCD deployment information is available

3. **Documentation**
   - API documentation is accessible through Backstage
   - Technical specifications and architecture details are maintained
   - Links to relevant resources and documentation

4. **Dependencies**
   - External dependencies are tracked
   - Security vulnerabilities are monitored
   - License compliance is maintained

To view this component in Backstage:
1. Navigate to your Backstage instance
2. Search for "todo-api" or "go-todo-app"
3. Access component details, CI/CD status, and documentation

## API Endpoints

### Health Check
Check if the API is running.

```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "status": "ok"
}
```

### Get All Todos
Retrieve all todo items.

```bash
curl http://localhost:8080/todos
```

Response:
```json
[
  {
    "id": "1",
    "title": "Learn Go",
    "completed": false
  },
  {
    "id": "2",
    "title": "Build API",
    "completed": false
  }
]
```

### Get a Specific Todo
Retrieve a single todo item by its ID.

```bash
curl http://localhost:8080/todos/1
```

Response:
```json
{
  "id": "1",
  "title": "Learn Go",
  "completed": false
}
```

### Create a New Todo
Create a new todo item.

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{
    "id": "3",
    "title": "Test API",
    "completed": false
  }'
```

Response:
```json
{
  "id": "3",
  "title": "Test API",
  "completed": false
}
```

## Error Handling

The API returns appropriate HTTP status codes and error messages:

- `200 OK`: Successful request
- `201 Created`: Resource successfully created
- `400 Bad Request`: Invalid request body
- `404 Not Found`: Resource not found

Example error response:
```json
{
  "error": "Todo not found"
}
```

## Project Structure

```
.
├── main.go      # Main application code
├── go.mod       # Go module file
└── README.md    # This file
```
## License

This project is licensed under the MIT License - see the LICENSE file for details. 