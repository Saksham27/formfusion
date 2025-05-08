# FormFusion

FormFusion is a mid-market form builder solution that bridges the gap between simple, free tools and complex, expensive enterprise solutions. The platform focuses on providing AI-powered form optimization, enhanced collaboration features, true no-code advanced logic, and a seamless mobile experience.

## Development Setup

### Prerequisites
- Docker Desktop installed on your machine
- Git for version control
- A code editor (VS Code recommended)

### Getting Started

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd formfusion
   ```

2. **Start the development container**
   ```bash
   docker compose up --build -d
   ```

3. **Access the container**
   ```bash
   docker exec -it formfusion-dev-1 bash
   ```

### Development Workflow

#### Local Development
- Edit files directly on your local machine using your preferred editor
- No need to start/stop container for editing files
- Git operations are performed on your local machine

#### Container Usage
Start the container when you need to:
- Run development commands (npm, dotnet, go)
- Start the application
- Run tests
- Use development tools installed in the container

#### Typical Development Session
1. Start your day:
   ```bash
   docker compose up -d
   ```

2. During development:
   - Edit files locally
   - Run commands in container when needed
   - Container stays running in background

3. End your day:
   ```bash
   docker compose down
   ```

### Development Container Commands

- **Start the container:**
  ```bash
  docker compose up -d
  ```

- **Stop the container:**
  ```bash
  docker compose down
  ```

- **View container logs:**
  ```bash
  docker compose logs
  ```

- **Rebuild container (after Dockerfile changes):**
  ```bash
  docker compose up --build -d
  ```

### Project Structure
```
formfusion/
├── .devcontainer/     # Development container configuration
├── docs/             # Project documentation
├── Dockerfile        # Container definition
├── docker-compose.yml # Container orchestration
└── README.md         # This file
```

### Technology Stack
- Frontend: React with Next.js
- Backend: ASP.NET Core (C#)
- Microservices: Go (Golang)
- Database: PostgreSQL
- Caching: Redis
- File Storage: Azure Blob Storage/AWS S3

## Contributing
1. Create a new branch for your feature
2. Make your changes
3. Submit a pull request

## License
[Add your license information here]
