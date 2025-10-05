# Task Management API - Backend

Backend API untuk aplikasi Task Management menggunakan Golang, Gin Framework, dan PostgreSQL.

## Tech Stack

- **Golang** 1.21+
- **Gin** - Web Framework
- **GORM** - ORM
- **PostgreSQL** - Database
- **godotenv** - Environment variable management

## Struktur Project

```
backend/
├── database/
│   └── database.go       # Database connection
├── handlers/
│   └── task_handler.go   # Task CRUD handlers
├── models/
│   └── task.go          # Task model & validation
├── routes/
│   └── routes.go        # API routes
├── main.go              # Entry point
├── go.mod               # Go modules
├── .env                 # Environment variables (create from .env.example)
└── .env.example         # Example environment variables
```

## Prerequisites

- Go 1.21 atau lebih tinggi
- PostgreSQL 12+
- Git

## Installation & Setup

### 1. Clone Repository

```bash
git clone <repository-url>
cd backend
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Setup Database

Buat database PostgreSQL baru:

```sql
CREATE DATABASE task_management;
```

### 4. Configure Environment Variables

Copy file `.env.example` menjadi `.env`:

```bash
cp .env.example .env
```

Edit file `.env` sesuai konfigurasi database Anda:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=task_management
DB_SSLMODE=disable

SERVER_PORT=8080
```

### 5. Run the Application

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## API Endpoints

Base URL: `http://localhost:8080/api/v1`

### Health Check

```
GET /health
```

### Tasks

#### Create Task

```
POST /api/v1/tasks
Content-Type: application/json

{
  "title": "Complete project",
  "description": "Finish the task management app",
  "due_date": "2025-12-31",
  "status": "pending"
}
```

**Response:**
```json
{
  "message": "Task created successfully",
  "data": {
    "id": 1,
    "title": "Complete project",
    "description": "Finish the task management app",
    "due_date": "2025-12-31T00:00:00Z",
    "status": "pending",
    "created_at": "2025-10-02T10:00:00Z",
    "updated_at": "2025-10-02T10:00:00Z"
  }
}
```

#### Get All Tasks

```
GET /api/v1/tasks
```

**Response:**
```json
{
  "message": "Tasks retrieved successfully",
  "data": [...]
}
```

#### Get Single Task

```
GET /api/v1/tasks/:id
```

#### Update Task

```
PUT /api/v1/tasks/:id
Content-Type: application/json

{
  "title": "Updated title",
  "status": "in-progress"
}
```

#### Delete Task

```
DELETE /api/v1/tasks/:id
```

## Data Model

### Task

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| id | integer | auto | Primary key |
| title | string | yes | Task title (min 3 chars) |
| description | string | no | Task description |
| due_date | date | yes | Due date (YYYY-MM-DD) |
| status | enum | yes | pending, in-progress, completed |
| created_at | timestamp | auto | Creation timestamp |
| updated_at | timestamp | auto | Update timestamp |

## Validation Rules

- **title**: Required, minimum 3 characters
- **due_date**: Required, must be valid date format (YYYY-MM-DD)
- **status**: Must be one of: `pending`, `in-progress`, `completed`

## Error Handling

API mengembalikan error dalam format:

```json
{
  "error": "Error message description"
}
```

Status codes:
- `200` - Success
- `201` - Created
- `400` - Bad Request
- `404` - Not Found
- `500` - Internal Server Error

## Development

### Run with Hot Reload (Optional)

Install air untuk hot reload:

```bash
go install github.com/cosmtrek/air@latest
air
```

### Build for Production

```bash
go build -o task-management-api
./task-management-api
```

## Testing API

Gunakan tools seperti:
- Postman
- Thunder Client (VS Code extension)
- curl

Contoh dengan curl:

```bash
# Create task
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Task","due_date":"2025-12-31","status":"pending"}'

# Get all tasks
curl http://localhost:8080/api/v1/tasks

# Get single task
curl http://localhost:8080/api/v1/tasks/1

# Update task
curl -X PUT http://localhost:8080/api/v1/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"completed"}'

# Delete task
curl -X DELETE http://localhost:8080/api/v1/tasks/1
```

## Troubleshooting

### Database Connection Error

Pastikan:
- PostgreSQL sudah berjalan
- Kredensial database di `.env` sudah benar
- Database `task_management` sudah dibuat

### Port Already in Use

Ubah `SERVER_PORT` di file `.env` ke port lain (misalnya 8081)

## License

MIT