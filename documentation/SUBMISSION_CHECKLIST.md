# Submission Checklist - Dunia Coding Fullstack Developer Test

Checklist untuk memastikan semua requirements terpenuhi sebelum submission.

---

## 📋 General Requirements

- [ ] Code clean, maintainable, dan well-structured
- [ ] React v18 digunakan untuk frontend
- [ ] Golang digunakan untuk backend
- [ ] Follow industry best practices (naming conventions, folder structure, error handling)
- [ ] Code di-upload ke GitHub private repository
- [ ] User `alfiankurniawan85@gmail.com` sudah ditambahkan sebagai collaborator
- [ ] README.md lengkap dengan instruksi jelas

---

## 🔧 Backend Requirements

### Framework & Database
- [ ] Menggunakan Gin framework
- [ ] Menggunakan PostgreSQL sebagai database
- [ ] GORM untuk ORM

### API Endpoints
- [ ] **POST** `/api/v1/tasks` - Create Task ✅
- [ ] **GET** `/api/v1/tasks` - List All Tasks ✅
- [ ] **GET** `/api/v1/tasks/:id` - Get Single Task ✅
- [ ] **PUT** `/api/v1/tasks/:id` - Update Task ✅
- [ ] **DELETE** `/api/v1/tasks/:id` - Delete Task ✅

### Data Model
Task memiliki fields:
- [ ] `id` (auto-generated) ✅
- [ ] `title` (string, required) ✅
- [ ] `description` (string, optional) ✅
- [ ] `due_date` (date, required) ✅
- [ ] `status` (enum: "pending", "in-progress", "completed") ✅

### Validation Rules
- [ ] `title`: required, min length 3 ✅
- [ ] `due_date`: required, must be valid date ✅
- [ ] `status`: must be one of valid enums ✅

### Code Structure
- [ ] Routes terpisah (`routes/`) ✅
- [ ] Handlers terpisah (`handlers/`) ✅
- [ ] Models terpisah (`models/`) ✅
- [ ] Database connection terpisah (`database/`) ✅

### Additional Features
- [ ] CORS configuration untuk frontend ✅
- [ ] Error handling yang proper ✅
- [ ] Database migration otomatis ✅
- [ ] Environment variables untuk config ✅

---

## ⚛️ Frontend Requirements

### Framework & Version
- [ ] React v18 ✅
- [ ] Functional components (no class components) ✅
- [ ] Hooks (useState, useEffect) ✅

### HTTP Client
- [ ] Axios untuk API calls ✅
- [ ] Centralized API service (`services/api.js`) ✅

### Features
- [ ] **List all tasks** dalam table/list view ✅
- [ ] **Add new task** menggunakan form ✅
- [ ] **Edit task** (inline atau modal) ✅
- [ ] **Delete task** dengan konfirmasi ✅
- [ ] **Change task status** ✅

### UI/UX
- [ ] Task list menampilkan: title, description, due date, status ✅
- [ ] Simple dan clean UI ✅
- [ ] TailwindCSS untuk styling ✅
- [ ] Responsive design ✅
- [ ] Loading states ✅
- [ ] Error handling dan notifikasi ✅

### Components
- [ ] TaskForm component (reusable untuk create/edit) ✅
- [ ] TaskItem component (display task card) ✅
- [ ] Modal component ✅
- [ ] App component (main logic) ✅

---

## 📦 Repository & Documentation

### GitHub Repository
- [ ] Private repository created
- [ ] Code uploaded to GitHub
- [ ] `.gitignore` files untuk backend dan frontend
- [ ] No sensitive data (`.env`) di repository
- [ ] Clean commit history

### Collaborator
- [ ] Add `alfiankurniawan85@gmail.com` as collaborator
  - Go to: Settings → Collaborators → Add people
  - Enter: alfiankurniawan85@gmail.com
  - Send invitation

### Documentation Files

#### Root Level
- [ ] `README.md` - Main documentation ✅
- [ ] `PROJECT_STRUCTURE.md` - Struktur project ✅
- [ ] `SETUP_GUIDE.md` - Setup lengkap ✅
- [ ] `SUBMISSION_CHECKLIST.md` - This file ✅
- [ ] `.gitignore` - Root gitignore ✅
- [ ] `run.sh` - Helper script (Linux/macOS) ✅
- [ ] `run.bat` - Helper script (Windows) ✅

#### Backend
- [ ] `backend/README.md` - Backend documentation ✅
- [ ] `backend/.env.example` - Example environment variables ✅
- [ ] `backend/.gitignore` - Backend gitignore ✅
- [ ] All source code files ✅

#### Frontend
- [ ] `frontend/README.md` - Frontend documentation ✅
- [ ] `frontend/.gitignore` - Frontend gitignore ✅
- [ ] All source code files ✅

### README.md Must Explain

#### Installation
- [ ] How to install dependencies (backend) ✅
- [ ] How to install dependencies (frontend) ✅
- [ ] Database setup instructions ✅
- [ ] Environment variables configuration ✅

#### Running Locally
- [ ] How to run backend ✅
- [ ] How to run frontend ✅
- [ ] Port information ✅
- [ ] Prerequisites listed ✅

#### Additional Info
- [ ] Tech stack documented ✅
- [ ] API endpoints documented ✅
- [ ] Features list ✅
- [ ] Project structure explained ✅
- [ ] Troubleshooting section ✅

---

## ✅ Acceptance Criteria Verification

### Backend ✅
- [x] All CRUD operations work correctly
- [x] Validation prevents invalid input
- [x] Code is structured (routes, handlers, models)
- [x] PostgreSQL integration works
- [x] Error handling implemented
- [x] CORS configured

### Frontend ✅
- [x] Displays task list (title, description, due date, status)
- [x] User can create task
- [x] User can edit task
- [x] User can delete task
- [x] User can update status
- [x] Simple and clean UI
- [x] Form validation works
- [x] Responsive design

### Repository & Documentation ✅
- [x] Code uploaded to GitHub private repository
- [x] Collaborator added
- [x] README.md explains installation and running

---

## 🚀 Pre-Submission Testing

### Backend Tests
```bash
cd backend

# Test health endpoint
curl http://localhost:8080/health

# Test create task
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Test","due_date":"2025-12-31","status":"pending"}'

# Test get tasks
curl http://localhost:8080/api/v1/tasks

# Test update task
curl -X PUT http://localhost:8080/api/v1/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"completed"}'

# Test delete task
curl -X DELETE http://localhost:8080/api/v1/tasks/1
```

Results:
- [ ] Health check returns OK
- [ ] Create task works and returns task with ID
- [ ] Get tasks returns array
- [ ] Update task modifies data
- [ ] Delete task removes task

### Frontend Tests
Open `http://localhost:3000` and test:

- [ ] Page loads without errors
- [ ] Can see "Add Task" button
- [ ] Clicking "Add Task" opens modal
- [ ] Can fill and submit form
- [ ] New task appears in list
- [ ] Can edit existing task
- [ ] Can delete task (with confirmation)
- [ ] Can change status via dropdown
- [ ] Notifications appear
- [ ] No console errors

### Integration Tests
- [ ] Frontend successfully connects to backend
- [ ] CRUD operations work end-to-end
- [ ] Data persists in database
- [ ] Validation errors display properly
- [ ] Loading states show during API calls

---

## 📝 Final Steps Before Submission

### 1. Code Quality
```bash
# Backend - Format code
cd backend
go fmt ./...

# Frontend - Check for issues
cd frontend
npm run build  # Should build without errors
```

### 2. Clean Up
- [ ] Remove any debug console.log statements
- [ ] Remove any test/dummy data
- [ ] Remove unused imports
- [ ] Remove commented code
- [ ] Check for TODO comments

### 3. Git Commands
```bash
# Ensure everything is committed
git status

# Add all files
git add .

# Commit
git commit -m "Complete Task Management App - Backend (Go/Gin/PostgreSQL) + Frontend (React/Vite/TailwindCSS)"

# Push to GitHub
git push origin main
```

### 4. GitHub Repository Settings
- [ ] Repository is set to Private
- [ ] Collaborator `alfiankurniawan85@gmail.com` added and invitation sent
- [ ] Repository has proper description
- [ ] README.md is visible on repository homepage

### 5. Final Verification
- [ ] Clone repository to fresh directory
- [ ] Follow README.md instructions
- [ ] Verify app runs successfully
- [ ] Test all features work

---

## 📧 Submission Email Template

Subject: **Fullstack Developer Programming Test Submission - [Your Name]**

Body:
```
Halo,

Saya telah menyelesaikan Fullstack Developer Programming Test untuk Dunia Coding.

Repository Details:
- Repository URL: [Your GitHub Repository URL]
- Branch: main
- Collaborator: alfiankurniawan85@gmail.com (invitation sent)

Tech Stack:
- Backend: Golang + Gin Framework + PostgreSQL + GORM
- Frontend: React 18 + Vite + TailwindCSS + Axios

Repository includes:
✅ Complete backend API dengan CRUD operations
✅ Complete frontend React application
✅ Comprehensive documentation (README.md, SETUP_GUIDE.md)
✅ Database migrations
✅ Form validation
✅ Error handling
✅ Responsive UI

Terima kasih atas kesempatannya.

Best regards,
[Your Name]
```

---

## ⚠️ Common Mistakes to Avoid

- [ ] ❌ Forgetting to set repository to Private
- [ ] ❌ Not adding collaborator
- [ ] ❌ Committing `.env` file
- [ ] ❌ Not testing after cloning repository
- [ ] ❌ Missing README.md or incomplete documentation
- [ ] ❌ Not handling errors properly
- [ ] ❌ Missing validation
- [ ] ❌ Not using React 18
- [ ] ❌ Not using functional components
- [ ] ❌ Backend not using specified framework (Gin)
- [ ] ❌ Not using PostgreSQL

---

## 🎉 Completion

When all items are checked:
- ✅ Your submission is complete!
- ✅ Code quality is high
- ✅ Documentation is comprehensive
- ✅ All requirements are met
- ✅ Ready for review

**Good luck with your submission! 🚀**