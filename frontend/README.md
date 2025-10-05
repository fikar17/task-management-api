# Task Management App - Frontend

Frontend aplikasi Task Management menggunakan React v18, Vite, dan TailwindCSS.

## Tech Stack

- **React** 18.2.0
- **Vite** - Build tool
- **TailwindCSS** - Styling
- **Axios** - HTTP client

## Struktur Project

```
frontend/
├── public/
├── src/
│   ├── components/
│   │   ├── TaskForm.jsx      # Form untuk create/edit task
│   │   ├── TaskItem.jsx      # Komponen card task
│   │   └── Modal.jsx         # Modal component
│   ├── services/
│   │   └── api.js            # API service layer
│   ├── App.jsx               # Main application component
│   ├── main.jsx              # Entry point
│   └── index.css             # Global styles + Tailwind
├── index.html
├── package.json
├── vite.config.js
├── tailwind.config.js
└── postcss.config.js
```

## Prerequisites

- Node.js 16+ atau lebih tinggi
- npm atau yarn
- Backend API harus berjalan di `http://localhost:8080`

## Installation & Setup

### 1. Install Dependencies

```bash
npm install
```

atau dengan yarn:

```bash
yarn install
```

### 2. Run Development Server

```bash
npm run dev
```

atau dengan yarn:

```bash
yarn dev
```

Aplikasi akan berjalan di `http://localhost:3000`

### 3. Build for Production

```bash
npm run build
```

Output build akan ada di folder `dist/`

### 4. Preview Production Build

```bash
npm run preview
```

## Features

### 1. Task List
- Menampilkan semua tasks dalam bentuk card grid
- Responsive layout (1 kolom mobile, 2 kolom tablet, 3 kolom desktop)
- Status badge dengan warna berbeda
- Informasi lengkap: title, description, due date, status

### 2. Create Task
- Form modal untuk membuat task baru
- Validasi:
  - Title: required, minimum 3 karakter
  - Due date: required, format date
- Status default: "pending"

### 3. Edit Task
- Form modal untuk edit task
- Pre-filled dengan data existing
- Validasi sama dengan create

### 4. Delete Task
- Konfirmasi sebelum delete
- Notifikasi success/error

### 5. Update Status
- Dropdown pada setiap task card
- Update realtime
- Status: pending, in-progress, completed

### 6. UI/UX Features
- Loading spinner saat fetch data
- Error message jika backend tidak tersedia
- Success/error notifications
- Responsive design
- Empty state dengan call-to-action
- Modal dengan ESC key support
- Smooth transitions dan hover effects

## API Configuration

API base URL dikonfigurasi di `src/services/api.js`:

```javascript
const API_BASE_URL = 'http://localhost:8080/api/v1';
```

Jika backend berjalan di port lain, ubah URL tersebut.

## Component Structure

### App.jsx
Main component yang mengelola:
- State management
- API calls
- Task CRUD operations
- Modal state
- Notifications

### TaskForm.jsx
Reusable form component untuk:
- Create task
- Edit task
- Form validation
- Error handling

### TaskItem.jsx
Card component untuk menampilkan:
- Task information
- Status badge
- Quick status change
- Edit & Delete buttons

### Modal.jsx
Reusable modal component dengan:
- Backdrop overlay
- ESC key support
- Body scroll lock
- Smooth animations

## Styling

Project menggunakan TailwindCSS dengan utility-first approach:
- Responsive breakpoints: `sm:`, `md:`, `lg:`
- Custom colors: Blue (primary), Red (danger), Green (success)
- Consistent spacing dan typography
- Hover states dan transitions

## Browser Support

- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)

## Troubleshooting

### Backend Connection Error

Jika muncul error "Failed to fetch tasks":
1. Pastikan backend running di `http://localhost:8080`
2. Check CORS configuration di backend
3. Cek network tab di browser developer tools

### Port Already in Use

Jika port 3000 sudah terpakai, Vite akan otomatis mencari port available. Atau ubah di `vite.config.js`:

```javascript
export default defineConfig({
  server: {
    port: 3001, // ganti port
  },
})
```

### Styling Not Applied

Jika TailwindCSS tidak muncul:
1. Pastikan `index.css` di-import di `main.jsx`
2. Check `tailwind.config.js` content paths
3. Restart dev server

## Development Tips

### Hot Module Replacement (HMR)
Vite mendukung HMR, perubahan akan langsung terlihat tanpa refresh

### React DevTools
Install React DevTools extension untuk debugging:
- Chrome: [React Developer Tools](https://chrome.google.com/webstore/detail/react-developer-tools/fmkadmapgofadopljbjfkapdkoienihi)
- Firefox: [React DevTools](https://addons.mozilla.org/en-US/firefox/addon/react-devtools/)

### VS Code Extensions (Recommended)
- ES7+ React/Redux/React-Native snippets
- Tailwind CSS IntelliSense
- ESLint
- Prettier

## License

MIT