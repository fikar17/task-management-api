# Quick Start Guide

Panduan singkat untuk menjalankan aplikasi Task Management.

## ⚡ TL;DR (For Experienced Developers)

```bash
# 1. Setup Database
createdb task_management

# 2. Backend Setup
cd backend
cp .env.example .env
# Edit .env dengan database credentials
go mod download
go run main.go

# 3. Frontend Setup (terminal baru)
cd frontend
npm install
npm run dev

# Open http://localhost:3000
```

---

## 📦 Prerequisites

Install terlebih dahulu:
- **Go** 1.21+ → https://golang.org/dl/
- **Node.js** 16+ → https://nodejs.org/
- **PostgreSQL** 12+ → https://www.postgresql.org/download/

---

## 🚀 5-Minute Setup

### Step 1: Clone Repository

```bash
git clone <repository-url>
cd task-management
```

### Step 2: Database

```bash
# Create database
psql -U postgres -c "CREATE DATABASE task_management;"
```

Atau gunakan pgAdmin atau tool GUI lainnya.

### Step 3: Backend

```bash
cd backend

# Copy environment config
cp .env.example .env

# Edit .env (gunakan text editor favorit)
# Minimal yang perlu diubah:
# DB_PASSWORD=your_postgres_password

# Install dependencies dan run
go mod download
go run main.go
```

✅ Backend running di `http://localhost:8080`

### Step 4: Frontend

**Buka terminal baru** (jangan tutup backend):

```bash
cd frontend

# Install dependencies
npm install

# Run development server
npm run dev
```

✅ Frontend running di `http://localhost:3000`

### Step 5: Test

1. Buka browser: `http://localhost:3000`
2. Click "Add Task"
3. Isi form dan submit
4. Task muncul di list ✅

---

## 🔧 Configuration

### Backend (.env)

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password    # ← Ubah ini
DB_NAME=task_management
DB_SSLMODE=disable
SERVER_PORT=8080
```

### Frontend

Tidak perlu konfigurasi. API URL sudah di-set ke `http://localhost:8080`

---

## 🧪 Quick Test

Test backend dengan curl:

```bash
# Create task
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Task",
    "description": "Testing the API",
    "due_date": "2025-12-31",
    "status": "pending"
  }'

# Get all tasks
curl http://localhost:8080/api/v1/tasks
```

---

## 🛠️ Helper Scripts

### Linux/macOS

```bash
# Make script executable
chmod +x run.sh

# Run backend only
./run.sh backend

# Run frontend only
./run.sh frontend

# Run both
./run.sh both

# Setup database
./run.sh setup-db
```

### Windows

```cmd
REM Run backend only
run.bat backend

REM Run frontend only
run.bat frontend

REM Run both (opens 2 windows)
run.bat both

REM Setup database
run.bat setup-db
```

---

## 📱 Ports

| Service | Port | URL |
|---------|------|-----|
| Backend API | 8080 | http://localhost:8080 |
| Frontend | 3000 | http://localhost:3000 |
| PostgreSQL | 5432 | localhost:5432 |

---

## ❌ Common Issues & Quick Fixes

### Backend won't start

**Error:** `connection refused`

```bash
# Check PostgreSQL is running
sudo systemctl status postgresql  # Linux
brew services list                # macOS

# Start if not running
sudo systemctl start postgresql   # Linux
brew services start postgresql@15 # macOS
```

### Port already in use

**Error:** `address already in use`

```bash
# Change port in backend/.env
SERVER_PORT=8081

# Or kill process
# Linux/macOS
lsof -ti:8080 | xargs kill -9

# Windows
netstat -ano | findstr :8080
taskkill /PID <PID> /F
```

### Frontend can't connect

**Error:** `Network Error`

1. Pastikan backend running: `curl http://localhost:8080/health`
2. Check CORS di browser console
3. Restart both services

### Database not found

```bash
# List databases
psql -U postgres -l

# Create if missing
psql -U postgres -c "CREATE DATABASE task_management;"
```

---

## 📚 API Endpoints

### Base URL
```
http://localhost:8080/api/v1
```

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /tasks | Create task |
| GET | /tasks | Get all tasks |
| GET | /tasks/:id | Get single task |
| PUT | /tasks/:id | Update task |
| DELETE | /tasks/:id | Delete task |

### Example: Create Task

```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Complete project",
    "description": "Finish the task management app",
    "due_date": "2025-12-31",
    "status": "pending"
  }'
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

---

## 🎨 Features

### ✅ What You Can Do

- **Create** new tasks with title, description, due date, and status
- **View** all tasks in a beautiful card grid
- **Edit** existing tasks via modal form
- **Delete** tasks with confirmation
- **Change status** quickly via dropdown (pending → in-progress → completed)
- **Responsive** design works on mobile, tablet, and desktop
- **Validation** prevents invalid data
- **Notifications** show success/error messages

### 🎯 Task Status

- 🟡 **Pending** - Task belum dimulai
- 🔵 **In Progress** - Task sedang dikerjakan  
- 🟢 **Completed** - Task selesai

---

## 🔍 Troubleshooting Commands

```bash
# Check Go version
go version

# Check Node version
node --version

# Check PostgreSQL version
psql --version

# Check if backend is running
curl http://localhost:8080/health

# Check if frontend is running
curl http://localhost:3000

# View backend logs
cd backend && go run main.go

# View frontend logs
cd frontend && npm run dev

# Database connection test
psql -U postgres -d task_management -c "SELECT version();"
```

---

## 📖 More Documentation

Butuh info lebih detail?

- **[README.md](README.md)** - Overview lengkap project
- **[SETUP_GUIDE.md](SETUP_GUIDE.md)** - Setup detail step-by-step
- **[PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md)** - Struktur folder dan arsitektur
- **[backend/README.md](backend/README.md)** - Backend documentation
- **[frontend/README.md](frontend/README.md)** - Frontend documentation
- **[SUBMISSION_CHECKLIST.md](SUBMISSION_CHECKLIST.md)** - Checklist untuk submission

---

## 🆘 Need Help?

1. **Check logs** di terminal
2. **Browser console** (F12) untuk frontend errors
3. **Review [SETUP_GUIDE.md](SETUP_GUIDE.md)** untuk detailed troubleshooting
4. **Google error message** untuk solusi umum

---

## 🎉 You're All Set!

Jika aplikasi sudah running:
- ✅ Backend di http://localhost:8080
- ✅ Frontend di http://localhost:3000
- ✅ Bisa create, read, update, delete tasks
- ✅ Siap untuk development!

**Happy coding! 🚀**