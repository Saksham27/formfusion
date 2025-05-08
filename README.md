# FormFusion

FormFusion is a mid-market form builder solution that bridges the gap between simple, free tools and complex, expensive enterprise solutions. The platform focuses on providing AI-powered form optimization, enhanced collaboration features, true no-code advanced logic, and a seamless mobile experience.

## Technology Stack

- **Frontend**: React with Next.js, TypeScript, Tailwind CSS, Aceternity UI
- **Backend**: ASP.NET Core (C#), Entity Framework Core
- **Microservices**: Go (Golang) with Gin
- **Database**: PostgreSQL
- **Caching**: Redis
- **File Storage**: Azure Blob Storage/AWS S3
- **DevOps**: Docker, GitHub Actions

## Development Setup

### Prerequisites

- [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed on your machine
- [Git](https://git-scm.com/downloads) for version control
- A code editor ([VS Code](https://code.visualstudio.com/) recommended)

### Getting Started

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd formfusion
   ```

2. **Start all services**
   ```bash
   docker-compose up
   ```
   This will start all services defined in the docker-compose.yml file:
   - Frontend (Next.js) - http://localhost:3000
   - Backend (ASP.NET Core) - http://localhost:5000
   - Realtime Service (Go) - http://localhost:8080
   - PDF Service (Go) - http://localhost:8081
   - Analytics Service (Go) - http://localhost:8082
   - PostgreSQL - localhost:5432
   - Redis - localhost:6379
   - PgAdmin - http://localhost:5050

3. **Selective Service Startup**
   
   To run only specific services (e.g., just database and backend):
   ```bash
   docker-compose up postgres backend
   ```

### Project Structure

```
formfusion/
├── docker-compose.yml        # Service orchestration
├── frontend/                 # Next.js frontend application
│   ├── Dockerfile.frontend   # Frontend container definition
│   ├── package.json          # Node dependencies
│   └── ... (Next.js files)
├── backend/                  # ASP.NET Core backend API
│   ├── Dockerfile.backend    # Backend container definition
│   ├── FormFusion.csproj     # .NET project file
│   └── ... (ASP.NET Core files)
├── microservices/            # Go microservices
│   ├── realtime/             # Real-time collaboration service
│   │   ├── Dockerfile.go     # Go container definition
│   │   ├── main.go           # Entry point
│   │   └── ... (Go files)
│   ├── pdf-service/          # PDF generation service
│   │   ├── Dockerfile.go
│   │   └── ... (Go files)
│   └── analytics/            # Analytics processing service
│       ├── Dockerfile.go
│       └── ... (Go files)
└── README.md                 # This file
```

### Development Workflow

#### Code-First Database Approach

This project uses a Code-First approach with Entity Framework Core:
- Database schema is defined by C# model classes
- Migrations are generated from model changes
- Database is automatically updated when backend service starts

To create a new migration after model changes:
```bash
# Inside the backend container
dotnet ef migrations add <MigrationName>
```

#### Local Development

- Edit files directly on your local machine using your preferred editor
- Files are mounted into containers via Docker volumes
- Changes are reflected immediately due to hot-reloading configurations:
  - Frontend uses Next.js dev server
  - Backend uses dotnet watch
  - Go services use Air for hot reloading

#### Common Docker Commands

- **Start all services in the background:**
  ```bash
  docker-compose up -d
  ```

- **View logs from all services:**
  ```bash
  docker-compose logs -f
  ```

- **View logs from specific service:**
  ```bash
  docker-compose logs -f backend
  ```

- **Execute command in a service container:**
  ```bash
  docker-compose exec backend bash
  ```

- **Stop all services:**
  ```bash
  docker-compose down
  ```

- **Rebuild services after Dockerfile changes:**
  ```bash
  docker-compose build
  docker-compose up -d
  ```

### Accessing Services

- **Frontend application**: http://localhost:3000
- **Backend API**: http://localhost:5000
- **API Documentation**: http://localhost:5000/swagger
- **PgAdmin database interface**: http://localhost:5050
  - Email: admin@formfusion.com
  - Password: admin

### Running Database Migrations

Database migrations are automatically applied on startup. If you need to run them manually:

```bash
# Inside the backend container
docker-compose exec backend bash
dotnet ef database update
```

## Core Features (In Development)

- AI-powered form optimization
- Enhanced collaboration features
- True no-code advanced logic builder
- Seamless mobile experience
- Advanced analytics dashboard
- Integration ecosystem

## Contributing

1. Create a new branch for your feature
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes, following our code style guidelines

3. Commit with clear, descriptive messages
   ```bash
   git commit -m "Add feature: brief description"
   ```

4. Push your branch and submit a pull request
   ```bash
   git push origin feature/your-feature-name
   ```

5. Wait for code review and address any feedback

## License

[Add your license information here]