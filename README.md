# Task Management Application

Aplikasi manajemen tugas fullstack dengan Golang (Backend) dan React (Frontend).

## 📋 Deskripsi

Aplikasi Task Management sederhana yang memungkinkan user untuk:
- ✅ Membuat task baru
- 📝 Melihat daftar semua tasks
- ✏️ Edit task yang sudah ada
- 🗑️ Hapus task
- 🔄 Update status task (pending, in-progress, completed)

## 🛠️ Tech Stack

### Backend
- **Golang** 1.21+
- **Gin** Framework
- **GORM** ORM
- **PostgreSQL** Database

### Frontend
- **React** 18.2.0
- **Vite** Build Tool
- **TailwindCSS** Styling
- **Axios** HTTP Client

## 📁 Struktur Project

```
task-management/
├── backend/                 # Golang Backend API
│   ├── database/           # Database connection
│   ├── handlers/           # Request handlers
│   ├── models/             # Data models
│   ├── routes/             # API routes
│   ├── main.go            # Entry point
│   ├── go.mod             # Go dependencies
│   ├── .env               # Environment variables
│   └── README.md          # Backend documentation
│
├── frontend/               # React Frontend
│   ├── src/
│   │   ├── components/    # React components
│   │   ├── services/      # API services
│   │   ├── App.jsx        # Main component
│   │   └── main.jsx       # Entry point
│   ├── package.json       # Node dependencies
│   └── README.md          # Frontend documentation
│
└── README.md              # This file
```

## 🚀 Quick Start

### Prerequisites

1. **Golang** 1.21 atau lebih tinggi ([Download](https://golang.org/dl/))
2. **Node.js** 16+ dan npm ([Download](https://nodejs.org/))
3. **PostgreSQL** 12+ ([Download](https://www.postgresql.org/download/))
4. **Git** ([Download](https://git-scm.com/))

### Setup Database

1. Buat database PostgreSQL baru:

```sql
CREATE DATABASE task_management;
```

### Installation

#### 1. Clone Repository

```bash
git clone <repository-url>
cd task-management
```

#### 2. Setup Backend

```bash
# Masuk ke folder backend
cd backend

# Install dependencies
go mod download

# Copy environment variables
cp .env.example .env

# Edit .env dengan kredensial database Anda
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=postgres
# DB_PASSWORD=your_password
# DB_NAME=task_management
# DB_SSLMODE=disable
# SERVER_PORT=8080

# Run backend
go run main.go
```

Backend akan berjalan di `http://localhost:8080`

#### 3. Setup Frontend

Buka terminal baru:

```bash
# Masuk ke folder frontend
cd frontend

# Install dependencies
npm install

# Run frontend
npm run dev
```

Frontend akan berjalan di `http://localhost:3000`

### Verifikasi

1. Backend: Buka `http://localhost:8080/health` - harus return `{"status":"ok"}`
2. Frontend: Buka `http://localhost:3000` - aplikasi harus muncul
3. Coba create task baru untuk memastikan koneksi backend-frontend berfungsi

## 📚 API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Endpoints

#### Create Task
```http
POST /api/v1/tasks
Content-Type: application/json

{
  "title": "Complete project",
  "description": "Finish the task management app",
  "due_date": "2025-12-31",
  "status": "pending"
}
```

#### Get All Tasks
```http
GET /api/v1/tasks
```

#### Get Single Task
```http
GET /api/v1/tasks/:id
```

#### Update Task
```http
PUT /api/v1/tasks/:id
Content-Type: application/json

{
  "title": "Updated title",
  "status": "in-progress"
}
```

#### Delete Task
```http
DELETE /api/v1/tasks/:id
```

### Response Format

**Success:**
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

**Error:**
```json
{
  "error": "Error message"
}
```

## 🎨 Screenshots

### Task List
![Task List View - Grid layout dengan task cards]

### Create/Edit Task
![Modal form untuk create/edit task]

### Empty State
![Empty state dengan call-to-action]

## ✅ Features Checklist

### Backend ✅
- [x] Create Task endpoint
- [x] Read/List Tasks endpoint
- [x] Update Task endpoint
- [x] Delete Task endpoint
- [x] Input validation (title min 3 chars, due_date required)
- [x] PostgreSQL integration
- [x] Structured code (routes, handlers, models)
- [x] Error handling
- [x] CORS configuration
- [x] Auto migration

### Frontend ✅
- [x] List all tasks
- [x] Create new task
- [x] Edit existing task
- [x] Delete task
- [x] Change task status
- [x] Form validation
- [x] Loading states
- [x] Error handling
- [x] Responsive design
- [x] Notifications
- [x] Empty state
- [x] Modal for create/edit

## 🧪 Testing

### Backend Testing

Gunakan curl atau Postman:

```bash
# Create task
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Task","due_date":"2025-12-31","status":"pending"}'

# Get all tasks
curl http://localhost:8080/api/v1/tasks

# Update task
curl -X PUT http://localhost:8080/api/v1/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"completed"}'

# Delete task
curl -X DELETE http://localhost:8080/api/v1/tasks/1
```

### Frontend Testing

1. Buka aplikasi di browser
2. Test semua fitur:
   - Create task
   - Edit task
   - Delete task
   - Change status
   - Form validation

## 🔧 Development

### Backend Development

```bash
cd backend

# Run with hot reload (install air first)
go install github.com/cosmtrek/air@latest
air

# Format code
go fmt ./...

# Run tests (if available)
go test ./...
```

### Frontend Development

```bash
cd frontend

# Run dev server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

## 📦 Production Build

### Backend

```bash
cd backend
go build -o task-management-api
./task-management-api
```

### Frontend

```bash
cd frontend
npm run build
# Output ada di folder dist/
```

## 🐛 Troubleshooting

### Backend tidak bisa connect ke database

**Solusi:**
- Pastikan PostgreSQL sudah running
- Check kredensial di file `.env`
- Pastikan database `task_management` sudah dibuat
- Test koneksi: `psql -U postgres -d task_management`

### Frontend tidak bisa fetch data

**Solusi:**
- Pastikan backend running di port 8080
- Check CORS configuration di backend
- Buka browser console untuk lihat error
- Test API dengan curl atau Postman dulu

### Port already in use

**Backend:**
```bash
# Ubah SERVER_PORT di .env
SERVER_PORT=8081
```

**Frontend:**
```bash
# Ubah port di vite.config.js
server: { port: 3001 }
```

## 📝 Design Decisions

### Backend
- **Gin Framework**: Lightweight, fast, dan popular di Go community
- **GORM**: ORM yang powerful dengan auto-migration
- **PostgreSQL**: Reliable, feature-rich relational database
- **Clean Architecture**: Separated concerns (routes, handlers, models)

### Frontend
- **Vite**: Faster build times dibanding CRA
- **TailwindCSS**: Utility-first, consistent styling, no CSS file bloat
- **Functional Components**: Modern React best practice
- **Axios**: Better error handling dibanding fetch
- **Modal Pattern**: Better UX untuk forms

## 🔐 Security Notes

- Input validation di backend dan frontend
- SQL injection prevention dengan GORM prepared statements
- CORS properly configured
- Date format validation
- Error messages tidak expose sensitive info

## 📄 License

MIT License

## 👤 Author

Dibuat untuk Dunia Coding Fullstack Developer Test

## 🤝 Contributing

1. Fork repository
2. Create feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Open Pull Request

## 📞 Contact

Untuk pertanyaan atau feedback, silakan buat issue di repository ini.

---

**Happy Coding! 🚀**