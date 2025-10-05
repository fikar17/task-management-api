# Setup Guide - Task Management Application

Panduan lengkap untuk setup dan menjalankan aplikasi Task Management dari awal.

## 📋 Table of Contents

1. [Prerequisites](#prerequisites)
2. [Database Setup](#database-setup)
3. [Backend Setup](#backend-setup)
4. [Frontend Setup](#frontend-setup)
5. [Verification](#verification)
6. [Common Issues](#common-issues)

---

## Prerequisites

### 1. Install Golang

**Windows:**
1. Download dari https://golang.org/dl/
2. Jalankan installer
3. Verifikasi: `go version`

**macOS:**
```bash
brew install go
go version
```

**Linux:**
```bash
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```

### 2. Install Node.js

**Windows/macOS:**
Download dari https://nodejs.org/ (LTS version)

**Linux:**
```bash
curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash -
sudo apt-get install -y nodejs
```

Verifikasi:
```bash
node --version
npm --version
```

### 3. Install PostgreSQL

**Windows:**
Download installer dari https://www.postgresql.org/download/windows/

**macOS:**
```bash
brew install postgresql@15
brew services start postgresql@15
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

Verifikasi:
```bash
psql --version
```

### 4. Install Git

**Windows:**
Download dari https://git-scm.com/download/win

**macOS:**
```bash
brew install git
```

**Linux:**
```bash
sudo apt install git
```

Verifikasi:
```bash
git --version
```

---

## Database Setup

### 1. Akses PostgreSQL

**Linux/macOS:**
```bash
sudo -u postgres psql
```

**Windows:**
Buka "SQL Shell (psql)" dari Start Menu

### 2. Create Database

```sql
-- Create database
CREATE DATABASE task_management;

-- Verify
\l

-- Connect to database
\c task_management

-- Exit
\q
```

### 3. Create User (Optional)

Jika ingin user khusus selain postgres:

```sql
CREATE USER taskuser WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE task_management TO taskuser;
```

### 4. Test Connection

```bash
psql -U postgres -d task_management -c "SELECT version();"
```

---

## Backend Setup

### 1. Clone/Create Project

```bash
# Clone repository
git clone <repository-url>
cd task-management

# Atau create folder baru
mkdir task-management
cd task-management
mkdir backend
cd backend
```

### 2. Initialize Go Module

```bash
go mod init task-management-api
```

### 3. Create Project Structure

```bash
mkdir database handlers models routes
```

### 4. Create Files

Copy semua file backend yang sudah saya berikan:
- `main.go`
- `database/database.go`
- `models/task.go`
- `handlers/task_handler.go`
- `routes/routes.go`
- `go.mod`
- `.env.example`
- `.gitignore`

### 5. Install Dependencies

```bash
go mod download
```

atau manual:

```bash
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
go get github.com/lib/pq
```

### 6. Configure Environment

```bash
# Copy example
cp .env.example .env

# Edit .env
nano .env  # atau text editor favorit Anda
```

File `.env`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_NAME=task_management
DB_SSLMODE=disable

SERVER_PORT=8080
```

### 7. Run Backend

```bash
go run main.go
```

Output yang diharapkan:
```
Database connected successfully!
[GIN-debug] Listening and serving HTTP on :8080
```

### 8. Test Backend

Buka browser atau terminal baru:

```bash
# Health check
curl http://localhost:8080/health

# Create task
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Task","due_date":"2025-12-31","status":"pending"}'

# Get tasks
curl http://localhost:8080/api/v1/tasks
```

---

## Frontend Setup

### 1. Create Frontend Project

```bash
# Dari root project
cd ..
mkdir frontend
cd frontend
```

### 2. Create Files

Copy semua file frontend yang sudah saya berikan:
- `package.json`
- `vite.config.js`
- `tailwind.config.js`
- `postcss.config.js`
- `index.html`
- `.gitignore`

### 3. Create Source Structure

```bash
mkdir -p src/components src/services
```

Copy files:
- `src/main.jsx`
- `src/App.jsx`
- `src/index.css`
- `src/components/TaskForm.jsx`
- `src/components/TaskItem.jsx`
- `src/components/Modal.jsx`
- `src/services/api.js`

### 4. Install Dependencies

```bash
npm install
```

Jika ada error, coba:
```bash
npm install --legacy-peer-deps
```

### 5. Verify Package Installation

```bash
npm list --depth=0
```

Should show:
```
├── axios@1.6.2
├── react@18.2.0
├── react-dom@18.2.0
├── tailwindcss@3.3.6
├── vite@5.0.8
└── ...
```

### 6. Run Frontend

```bash
npm run dev
```

Output yang diharapkan:
```
  VITE v5.0.8  ready in 500 ms

  ➜  Local:   http://localhost:3000/
  ➜  Network: use --host to expose
```

### 7. Open Browser

Buka `http://localhost:3000`

Aplikasi harus muncul dengan UI Task Management.

---

## Verification

### Complete System Check

#### 1. Backend Health

```bash
curl http://localhost:8080/health
```

Expected: `{"status":"ok"}`

#### 2. Database Connection

```bash
curl http://localhost:8080/api/v1/tasks
```

Expected: `{"message":"Tasks retrieved successfully","data":[]}`

#### 3. Create Test Task

```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title":"Test Task",
    "description":"Testing the API",
    "due_date":"2025-12-31",
    "status":"pending"
  }'
```

Expected: Success response dengan task data

#### 4. Frontend Integration

1. Buka http://localhost:3000
2. Click "Add Task" button
3. Fill form:
   - Title: "My First Task"
   - Description: "Testing the app"
   - Due Date: Select any future date
   - Status: "pending"
4. Click "Create Task"
5. Task harus muncul di list
6. Coba edit, delete, dan change status

#### 5. Full CRUD Test Checklist

- [ ] Create task berhasil
- [ ] Task muncul di list
- [ ] Edit task berhasil
- [ ] Delete task berhasil (dengan konfirmasi)
- [ ] Change status berhasil
- [ ] Form validation bekerja (title < 3 chars)
- [ ] Date validation bekerja
- [ ] Notification muncul
- [ ] Loading state terlihat saat fetch

---

## Common Issues

### Issue 1: Backend - Database Connection Failed

**Error:**
```
Failed to connect to database: connection refused
```

**Solutions:**

1. Check PostgreSQL is running:
```bash
# Linux/macOS
sudo systemctl status postgresql

# Windows
# Check Services app untuk PostgreSQL

# macOS with Homebrew
brew services list
```

2. Start PostgreSQL if not running:
```bash
# Linux
sudo systemctl start postgresql

# macOS
brew services start postgresql@15

# Windows
# Start dari Services app
```

3. Verify credentials in `.env`:
```bash
# Test connection
psql -U postgres -d task_management
```

4. Check database exists:
```sql
\l
```

### Issue 2: Backend - Port Already in Use

**Error:**
```
listen tcp :8080: bind: address already in use
```

**Solutions:**

1. Find process using port:
```bash
# Linux/macOS
lsof -i :8080

# Windows
netstat -ano | findstr :8080
```

2. Kill the process:
```bash
# Linux/macOS
kill -9 <PID>

# Windows
taskkill /PID <PID> /F
```

3. Or change port in `.env`:
```env
SERVER_PORT=8081
```

### Issue 3: Frontend - Cannot Fetch Data

**Error in Browser Console:**
```
Failed to fetch
CORS error
Network Error
```

**Solutions:**

1. Ensure backend is running:
```bash
curl http://localhost:8080/health
```

2. Check API URL in `src/services/api.js`:
```javascript
const API_BASE_URL = 'http://localhost:8080/api/v1';
```

3. Verify CORS in backend `main.go`:
```go
config.AllowOrigins = []string{"http://localhost:3000"}
```

4. Check browser console for exact error

5. Test API directly with curl:
```bash
curl http://localhost:8080/api/v1/tasks
```

### Issue 4: Frontend - npm install Fails

**Error:**
```
npm ERR! ERESOLVE unable to resolve dependency tree
```

**Solutions:**

1. Clear npm cache:
```bash
npm cache clean --force
```

2. Delete node_modules and package-lock.json:
```bash
rm -rf node_modules package-lock.json
npm install
```

3. Use legacy peer deps:
```bash
npm install --legacy-peer-deps
```

4. Use different Node version (try LTS):
```bash
nvm install --lts
nvm use --lts
```

### Issue 5: Backend - go.mod Issues

**Error:**
```
package not found
module not found
```

**Solutions:**

1. Clean module cache:
```bash
go clean -modcache
```

2. Tidy dependencies:
```bash
go mod tidy
```

3. Download dependencies:
```bash
go mod download
```

4. Verify go.mod file exists and is correct

### Issue 6: Database - Migration Errors

**Error:**
```
ERROR: relation "tasks" already exists
```

**Solutions:**

1. Drop and recreate table:
```sql
DROP TABLE IF EXISTS tasks CASCADE;
```

2. Or use fresh database:
```sql
DROP DATABASE task_management;
CREATE DATABASE task_management;
```

3. Restart backend to run migrations

### Issue 7: Frontend - Blank Page

**Solutions:**

1. Check browser console for errors

2. Verify all files are in correct location:
```bash
ls src/
# Should show: App.jsx, main.jsx, index.css, components/, services/
```

3. Check main.jsx imports App.jsx correctly

4. Restart dev server:
```bash
# Ctrl+C to stop
npm run dev
```

### Issue 8: TailwindCSS Not Working

**Solutions:**

1. Verify tailwind.config.js content paths:
```javascript
content: [
  "./index.html",
  "./src/**/*.{js,ts,jsx,tsx}",
],
```

2. Check index.css has Tailwind directives:
```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

3. Verify main.jsx imports index.css:
```javascript
import './index.css'
```

4. Restart dev server

### Issue 9: Backend - GORM Auto-Migration Failed

**Error:**
```
Failed to run migrations
```

**Solutions:**

1. Check database permissions:
```sql
GRANT ALL PRIVILEGES ON DATABASE task_management TO postgres;
```

2. Manually create table:
```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date DATE NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

3. Check PostgreSQL version compatibility (12+)

### Issue 10: Git Push Issues

**Error:**
```
Large files detected
```

**Solutions:**

1. Ensure .gitignore is in place:
```bash
# Backend
cat backend/.gitignore

# Frontend  
cat frontend/.gitignore
```

2. Remove node_modules from git:
```bash
git rm -r --cached frontend/node_modules
git commit -m "Remove node_modules"
```

3. Remove vendor from git:
```bash
git rm -r --cached backend/vendor
git commit -m "Remove vendor"
```

---

## Development Tips

### Backend Development

#### Hot Reload with Air

Install Air:
```bash
go install github.com/cosmtrek/air@latest
```

Create `.air.toml` in backend folder:
```toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./tmp/main ."
  bin = "tmp/main"
  full_bin = "./tmp/main"
  include_ext = ["go", "tpl", "tmpl", "html"]
  exclude_dir = ["assets", "tmp", "vendor"]
  delay = 1000
  stop_on_error = true

[log]
  time = true

[color]
  main = "magenta"
  watcher = "cyan"
  build = "yellow"
  runner = "green"
```

Run with hot reload:
```bash
air
```

#### Debug with Delve

Install Delve:
```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

Debug:
```bash
dlv debug
```

### Frontend Development

#### VS Code Extensions

Install recommended extensions:
- ES7+ React/Redux/React-Native snippets
- Tailwind CSS IntelliSense
- ESLint
- Prettier - Code formatter

Create `.vscode/extensions.json`:
```json
{
  "recommendations": [
    "dsznajder.es7-react-js-snippets",
    "bradlc.vscode-tailwindcss",
    "dbaeumer.vscode-eslint",
    "esbenp.prettier-vscode"
  ]
}
```

#### React DevTools

Install browser extension untuk debugging:
- Chrome: React Developer Tools
- Firefox: React Developer Tools

#### Quick Commands

```bash
# Install single package
npm install axios

# Update dependencies
npm update

# Check for outdated packages
npm outdated

# Audit security issues
npm audit

# Fix security issues
npm audit fix
```

---

## Performance Optimization

### Backend

1. **Add Database Indexes:**
```sql
CREATE INDEX idx_tasks_status ON tasks(status);
CREATE INDEX idx_tasks_due_date ON tasks(due_date);
CREATE INDEX idx_tasks_created_at ON tasks(created_at);
```

2. **Connection Pooling:**
Already handled by GORM, but can configure:
```go
sqlDB, _ := db.DB()
sqlDB.SetMaxIdleConns(10)
sqlDB.SetMaxOpenConns(100)
```

3. **Add Pagination:**
```go
// In handler
page := c.DefaultQuery("page", "1")
limit := c.DefaultQuery("limit", "10")
```

### Frontend

1. **Code Splitting:**
```javascript
const TaskForm = lazy(() => import('./components/TaskForm'));
```

2. **Memoization:**
```javascript
const MemoizedTaskItem = memo(TaskItem);
```

3. **Debounce Search:**
```javascript
const debouncedSearch = useMemo(
  () => debounce((query) => fetchTasks(query), 300),
  []
);
```

---

## Production Deployment

### Backend Deployment

1. **Build Binary:**
```bash
go build -o task-management-api
```

2. **Use Environment Variables:**
Never commit `.env` to git

3. **Run with PM2 or systemd:**
```bash
# PM2
pm2 start ./task-management-api --name task-api

# systemd
sudo systemctl start task-api
```

### Frontend Deployment

1. **Build for Production:**
```bash
npm run build
```

2. **Serve with nginx:**
```nginx
server {
    listen 80;
    server_name your-domain.com;
    
    root /var/www/task-management/dist;
    index index.html;
    
    location / {
        try_files $uri $uri/ /index.html;
    }
}
```

3. **Update API URL:**
Change in `src/services/api.js`:
```javascript
const API_BASE_URL = 'https://api.your-domain.com/api/v1';
```

---

## Testing

### Backend Testing

Create test file `handlers/task_handler_test.go`:
```go
func TestCreateTask(t *testing.T) {
    // Test implementation
}
```

Run tests:
```bash
go test ./...
```

### Frontend Testing

Install testing libraries:
```bash
npm install --save-dev @testing-library/react @testing-library/jest-dom vitest
```

Create test file `src/App.test.jsx`:
```javascript
import { render, screen } from '@testing-library/react';
import App from './App';

test('renders task management title', () => {
  render(<App />);
  const titleElement = screen.getByText(/Task Management/i);
  expect(titleElement).toBeInTheDocument();
});
```

Run tests:
```bash
npm test
```

---

## Next Steps

After successful setup:

1. **Add Features:**
   - Search functionality
   - Filter by status
   - Sort by date
   - Task categories
   - User authentication

2. **Improve UI:**
   - Add animations
   - Better mobile responsiveness
   - Dark mode
   - Task priority colors

3. **Backend Improvements:**
   - Add pagination
   - Implement caching
   - Add logging
   - API rate limiting

4. **Documentation:**
   - API documentation (Swagger)
   - Component documentation (Storybook)
   - User guide

---

## Getting Help

Jika masih ada masalah:

1. Check logs:
   - Backend: Terminal output
   - Frontend: Browser console (F12)

2. Verify versions:
```bash
go version
node --version
psql --version
```

3. Google error message
4. Check Stack Overflow
5. Read documentation:
   - Gin: https://gin-gonic.com/docs/
   - React: https://react.dev/
   - TailwindCSS: https://tailwindcss.com/docs

---

**Good luck with your setup! 🚀**