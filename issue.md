# Issue: Setup & Inisialisasi Project Fullstack (Go Echo, Vue, MySQL, Bun)

## 1. Ringkasan Proyek
Membangun pondasi proyek fullstack baru di repositori ini yang memisahkan arsitektur backend dan frontend secara rapi, siap untuk dikembangkan lebih lanjut.

### Tech Stack:
- **Package Manager / Runtime Frontend**: Bun
- **Frontend**: Vue 3 (Vite + Vue)
- **Backend**: Go (Framework: Echo + GORM)
- **Database**: MySQL
- **Orkestrasi Lingkungan**: Docker Compose untuk MySQL lokal

---

## 2. Struktur Direktori
```text
play-vibe-code/
├── backend/               # Service Go (Echo + GORM)
│   ├── cmd/api/           # Entry point aplikasi (main.go)
│   ├── internal/          # Config, Handler, Model, Database connection
│   ├── go.mod
│   └── go.sum
├── frontend/              # Web Client (Vue 3 via Bun)
│   ├── src/               # UI components, services, styling
│   ├── package.json
│   └── bun.lockb
├── docker-compose.yml     # Service MySQL 8 lokal
├── .env.example           # Variabel environment umum
├── README.md              # Panduan menjalankan proyek
└── issue.md               # Dokumen spesifikasi dan roadmap
```

---

## 3. Rencana Kerja (High-Level Milestones)

### Fase 1: Inisialisasi Struktur & Database
- [x] Buat folder `backend/` dan `frontend/`.
- [x] Sediakan konfigurasi environment (`.env.example`) untuk koneksi MySQL.
- [x] Sediakan file `docker-compose.yml` untuk menjalankan instance MySQL lokal.

### Fase 2: Backend Setup (Go + Echo + MySQL)
- [x] Inisialisasi Go Module di dalam folder `backend/`.
- [x] Install dependensi Echo (`labstack/echo/v4`) dan driver MySQL (`gorm.io/driver/mysql`).
- [x] Buat konfigurasi koneksi database MySQL dengan mekanisme health check.
- [x] Buat endpoint dasar:
  - `GET /health` : Mengecek status server dan koneksi database.
  - `GET /api/items`, `POST /api/items`, `DELETE /api/items/:id` : Endpoint CRUD.
- [x] Pasang middleware: Logger, Recover, dan CORS.

### Fase 3: Frontend Setup (Vue via Bun)
- [x] Inisialisasi project Vue 3 baru di folder `frontend/` menggunakan Bun.
- [x] Kelola seluruh dependensi dan script frontend menggunakan **Bun** (`bun install`, `bun dev`, `bun run build`).
- [x] Setup HTTP client di `src/services/api.js` mengarah ke backend Go.
- [x] Buat antarmuka responsif modern untuk status koneksi (`/health`) dan CRUD items.

### Fase 4: Integrasi & Validasi Alur
- [x] Komunikasi Frontend (Vue) dan Backend (Echo) berjalan tanpa kendala CORS.
- [x] Validasi data flow: Frontend -> Request API -> Backend Echo -> MySQL -> Response -> Render Vue.

### Fase 5: Dokumentasi Developer
- [x] Tulis instruksi lengkap di `README.md`.

---

## 4. Kriteria Keberhasilan (Acceptance Criteria)
- [x] Backend Go Echo dapat berjalan dan terhubung ke MySQL.
- [x] Endpoint `GET /health` mengembalikan status `200 OK` dan status koneksi database.
- [x] Frontend Vue berhasil di-scaffold dan dapat dijalankan serta di-build menggunakan Bun (`bun dev` / `bun run build`).
- [x] Halaman frontend berhasil mengambil data dari API backend via HTTP request.
- [x] Tersedia panduan lengkap untuk menjalankan seluruh service secara lokal di `README.md`.
