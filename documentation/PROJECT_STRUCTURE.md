# Project Structure - Task Management Application

## 📂 Complete Folder Structure

```
task-management/
│
├── backend/                           # Backend API (Golang)
│   ├── database/
│   │   └── database.go               # PostgreSQL connection & config
│   │
│   ├── handlers/
│   │   └── task_handler.go           # Task CRUD logic & business rules
│   │
│   ├── models/
│   │   └── task.go                   # Task model, validation, migration
│   │
│   ├── routes/
│   │   └── routes.go                 # API route definitions
│   │
│   ├── main.go                       # Application entry point
│   ├── go.mod                        # Go module dependencies
│   ├── go.sum                        # Go dependency checksums
│   ├── .env                          # Environment variables (not in git)
│   ├── .env.example                  # Example environment variables
│   ├── .gitignore                    # Git ignore rules
│   └── README.md                     # Backend documentation
│
├── frontend/                          # Frontend App (React)
│   ├── public/                       # Static assets
│   │
│   ├── src/
│   │   ├── components/
│   │   │   ├── TaskForm.jsx         # Reusable form for create/edit
│   │   │   ├── TaskItem.jsx         # Task card component
│   │   │   └── Modal.jsx            # Modal wrapper component
│   │   │
│   │   ├── services/
│   │   │   └── api.js               # Axios API client & endpoints
│   │   │
│   │   ├── App.jsx                  # Main application component
│   │   ├── main.jsx                 # React entry point
│   │   └── index.css                # Global styles + Tailwind imports
│   │
│   ├── index.html                    # HTML entry point
│   ├── package.json                  # NPM dependencies & scripts
│   ├── package-lock.json             # NPM lock file
│   ├── vite.config.js               # Vite configuration
│   ├── tailwind.config.js           # TailwindCSS configuration
│   ├── postcss.config.js            # PostCSS configuration
│   ├── .gitignore                   # Git ignore rules
│   └── README.md                    # Frontend documentation
│
├── README.md                         # Main project documentation
└── PROJECT_STRUCTURE.md             # This file

```

## 🏗️ Architecture Overview

### Backend Architecture (Clean Architecture)

```
Request Flow:
┌─────────┐     ┌────────┐     ┌──────────┐     ┌──────────┐
│ Client  │────▶│ Routes │────▶│ Handlers │────▶│ Database │
└─────────┘     └────────┘     └──────────┘     └──────────┘
                    │                │                │
                    │                │                │
                    │                ▼                │
                    │           ┌────────┐            │
                    └──────────▶│ Models │◀───────────┘
                                └────────┘
```

**Layers:**
1. **Routes Layer** (`routes/`): Mendefinisikan endpoint URLs dan method
2. **Handlers Layer** (`handlers/`): Business logic dan request/response handling
3. **Models Layer** (`models/`): Data structures, validation, database schema
4. **Database Layer** (`database/`): Database connection dan configuration

### Frontend Architecture (Component-Based)

```
Component Hierarchy:
┌──────────────────────────────────────┐
│              App.jsx                  │
│  (State Management & API Calls)      │
│                                       │
│  ┌────────────────────────────────┐  │
│  │        Modal.jsx               │  │
│  │  ┌──────────────────────────┐  │  │
│  │  │     TaskForm.jsx         │  │  │
│  │  │  (Create/Edit Form)      │  │  │
│  │  └──────────────────────────┘  │  │
│  └────────────────────────────────┘  │
│                                       │
│  ┌────────────────────────────────┐  │
│  │  TaskItem.jsx (x N)            │  │
│  │  (Task Cards in Grid)          │  │
│  └────────────────────────────────┘  │
└──────────────────────────────────────┘
            ▲
            │
    ┌───────┴────────┐
    │  services/     │
    │    api.js      │
    │  (API Client)  │
    └────────────────┘
```

## 📋 File Responsibilities

### Backend Files

#### `main.go`
- Application entry point
- Environment setup
- Database connection initialization
- Server startup
- CORS configuration

#### `database/database.go`
- PostgreSQL connection
- Database configuration dari environment variables
- Connection pool management

#### `models/task.go`
- Task struct definition
- Input validation structs
- Database migration
- Field validation rules
- Status enum definition

#### `handlers/task_handler.go`
- CreateTask: Handle POST request, validate input
- GetTasks: Retrieve all tasks from DB
- GetTask: Retrieve single task by ID
- UpdateTask: Update existing task
- DeleteTask: Delete task by ID
- Status validation helper

#### `routes/routes.go`
- API route registration
- Group routing (api/v1)
- Handler mapping
- Health check endpoint

### Frontend Files

#### `src/main.jsx`
- React application bootstrap
- Root component mounting
- Strict mode wrapper

#### `src/App.jsx`
- Main application logic
- State management (tasks, loading, errors, modals)
- API integration
- CRUD operations orchestration
- Notification system
- Modal state management

#### `src/components/TaskForm.jsx`
**Props:**
- `onSubmit`: Function to handle form submission
- `onCancel`: Function to close form
- `initialData`: Pre-fill data for editing

**Responsibilities:**
- Form rendering
- Client-side validation
- Input handling
- Error display

#### `src/components/TaskItem.jsx`
**Props:**
- `task`: Task data object
- `onEdit`: Callback for edit action
- `onDelete`: Callback for delete action
- `onStatusChange`: Callback for status update

**Responsibilities:**
- Display task information
- Status badge rendering
- Quick status change dropdown
- Action buttons (Edit/Delete)

#### `src/components/Modal.jsx`
**Props:**
- `isOpen`: Boolean to show/hide modal
- `onClose`: Callback to close modal
- `title`: Modal header title
- `children`: Modal content

**Responsibilities:**
- Modal overlay and backdrop
- ESC key handling
- Body scroll lock
- Click outside to close

#### `src/services/api.js`
- Axios instance configuration
- API base URL setup
- Task CRUD API methods:
  - `getAllTasks()`
  - `getTask(id)`
  - `createTask(data)`
  - `updateTask(id, data)`
  - `deleteTask(id)`

## 🔄 Data Flow

### Create Task Flow

```
1. User clicks "Add Task" button
   └─▶ App.jsx opens modal (setIsModalOpen(true))

2. User fills TaskForm and submits
   └─▶ TaskForm validates input
       └─▶ Calls onSubmit callback

3. App.jsx.handleCreateTask
   └─▶ Calls taskAPI.createTask(data)
       └─▶ api.js sends POST to backend

4. Backend receives request
   └─▶ routes.go matches POST /api/v1/tasks
       └─▶ task_handler.CreateTask processes
           └─▶ Validates input
               └─▶ Saves to database
                   └─▶ Returns success response

5. Frontend receives response
   └─▶ Shows success notification
       └─▶ Closes modal
           └─▶ Refreshes task list
```

### Update Status Flow

```
1. User changes status dropdown on TaskItem
   └─▶ TaskItem.handleStatusChange fires
       └─▶ Calls onStatusChange(taskId, newStatus)

2. App.jsx.handleStatusChange
   └─▶ Calls taskAPI.updateTask(id, {status: newStatus})
       └─▶ api.js sends PUT to backend

3. Backend processes update
   └─▶ Validates status value
       └─▶ Updates database
           └─▶ Returns updated task

4. Frontend updates UI
   └─▶ Shows notification
       └─▶ Refreshes task list
```

## 🎯 Design Patterns Used

### Backend
1. **Repository Pattern**: Database layer abstraction
2. **Handler Pattern**: Request handling separation
3. **Model-View-Controller (MVC)**: Adapted for API
4. **Dependency Injection**: Handler receives DB instance

### Frontend
1. **Component Composition**: Reusable components
2. **Container/Presentation**: App.jsx (Container), TaskItem (Presentation)
3. **Render Props**: Modal with children prop
4. **Controlled Components**: Forms with state
5. **Service Layer**: API abstraction

## 📊 Database Schema

```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date DATE NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT check_title_length CHECK (LENGTH(title) >= 3),
    CONSTRAINT check_status CHECK (status IN ('pending', 'in-progress', 'completed'))
);

CREATE INDEX idx_tasks_status ON tasks(status);
CREATE INDEX idx_tasks_due_date ON tasks(due_date);
```

## 🔐 Environment Variables

### Backend (.env)
```
DB_HOST=localhost          # Database host
DB_PORT=5432              # PostgreSQL port
DB_USER=postgres          # Database user
DB_PASSWORD=password      # Database password
DB_NAME=task_management   # Database name
DB_SSLMODE=disable        # SSL mode
SERVER_PORT=8080          # API server port
```

### Frontend
No environment variables needed for development.
For production, set API base URL in `src/services/api.js`

## 🚀 Deployment Considerations

### Backend Deployment
- Set environment variables on server
- Use production database
- Enable SSL for database connection
- Use reverse proxy (nginx)
- Set proper CORS origins

### Frontend Deployment
- Update API_BASE_URL to production backend
- Build with `npm run build`
- Serve `dist/` folder with static file server
- Configure proper redirects for SPA routing

## 📝 Naming Conventions

### Backend (Go)
- Files: `snake_case.go`
- Packages: `lowercase`
- Structs: `PascalCase`
- Functions: `PascalCase` (exported), `camelCase` (private)
- Variables: `camelCase`
- Constants: `PascalCase` or `SCREAMING_SNAKE_CASE`

### Frontend (JavaScript/React)
- Files: `PascalCase.jsx` (components), `camelCase.js` (utilities)
- Components: `PascalCase`
- Functions: `camelCase`
- Constants: `SCREAMING_SNAKE_CASE`
- CSS Classes: `kebab-case` (Tailwind utilities)

---

**Last Updated:** October 2025