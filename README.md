# Play Vibe Code

Fullstack starter application built with:
- **Runtime / Package Manager**: [Bun](https://bun.sh)
- **Frontend**: [Vue 3](https://vuejs.org) + [Vite](https://vitejs.dev)
- **Backend**: [Go](https://go.dev) + [Echo](https://echo.labstack.com)
- **Database**: [MySQL](https://www.mysql.com) (GORM)
- **Containerization**: [Docker Compose](https://docs.docker.com/compose/)

---

## 📁 Struktur Direktori

```text
play-vibe-code/
├── backend/               # Service Backend (Go + Echo + GORM)
│   ├── cmd/api/           # Entry point aplikasi (main.go)
│   ├── internal/          # Config, Database, Model, Handlers
│   ├── go.mod
│   └── go.sum
├── frontend/              # Web Client (Vue 3 + Bun)
│   ├── src/               # Komponen UI, Services, Design System
│   ├── package.json
│   └── bun.lockb
├── docker-compose.yml     # Service MySQL 8 lokal
├── .env.example           # Template environment variable
└── issue.md               # Spesifikasi & Rencana Proyek
```

---

## 🚀 Panduan Menjalankan Proyek

### 1. Menjalankan Database MySQL
Gunakan Docker Compose untuk menjalankan database lokal:

```bash
docker compose up -d
```

Service MySQL akan berjalan pada port `3306` dengan kredensial default dari `docker-compose.yml`:
- **Host**: `127.0.0.1:3306`
- **Database**: `appdb`
- **User**: `appuser`
- **Password**: `appsecret`

---

### 2. Menjalankan Backend (Go Echo)
Masuk ke direktori `backend/` dan jalankan server:

```bash
cd backend
go run ./cmd/api
```

Backend API akan aktif di `http://localhost:8080`.
- Health check: `GET http://localhost:8080/health`
- List items: `GET http://localhost:8080/api/items`
- Create item: `POST http://localhost:8080/api/items`

---

### 3. Menjalankan Frontend (Vue 3 via Bun)
Buka terminal baru, masuk ke direktori `frontend/`, lalu jalankan server pengembangan:

```bash
cd frontend
bun install
bun dev
```

Buka browser di `http://localhost:5173` untuk mengakses aplikasi web.

---

## 🛠️ API Endpoints

| Method | Endpoint | Deskripsi |
|---|---|---|
| `GET` | `/health` | Mengecek status server, uptime, dan koneksi MySQL |
| `GET` | `/api/items` | Mengambil seluruh daftar item dari database |
| `POST` | `/api/items` | Menambahkan item baru ke database |
| `DELETE` | `/api/items/:id` | Menghapus item berdasarkan ID |
