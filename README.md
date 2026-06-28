# Purchasing & Inventory Management System (Backend API)

Backend API untuk sistem Purchasing dan Manajemen Inventaris Barang, dibangun menggunakan **Go (Golang)**, **Fiber v2 (Web Framework)**, **GORM (ORM)**, dan database **MySQL**.

---

## Tech Stack

- **Language**: Go (Golang) v1.24+
- **Web Framework**: [Fiber v2](https://gofiber.io/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: MySQL / MariaDB
- **Auth**: JWT (JSON Web Tokens)
- **ID Generation**: ULID (Universally Unique Lexicographically Sortable Identifier)

---

## Struktur Project

Struktur project mengikuti pola modular yang rapi di bawah folder `app/`:

```text
task-be/
├── app/
│   ├── db/
│   │   └── connection.go       # Koneksi DB & Auto-Migration Model
│   ├── helper/
│   │   ├── blacklisttoken/     # Blacklist token handler (Logout)
│   │   ├── hashpassword.go     # BCrypt hashing utilitas
│   │   ├── message.go          # Standard return messages
│   │   ├── timehelper.go       # Formatting & manipulasi waktu (UTC)
│   │   └── token.go            # JWT claims generator & validator
│   ├── model/
│   │   ├── base.go             # BaseModel (ULID) & DateTime (timestamps)
│   │   ├── item.go             # Model database Item
│   │   ├── pusrchasing.go      # Model database Purchasing (Header)
│   │   ├── pusrchasingdetail.go# Model database Detail Purchasing
│   │   ├── supplier.go         # Model database Supplier
│   │   └── user.go             # Model database User
│   ├── service/
│   │   ├── base.go             # Custom JSON response helper
│   │   ├── validator.go        # Custom Bind & Validate body payload
│   │   ├── auth/               # Service Login & Logout
│   │   ├── item/               # Service Manajemen Item
│   │   ├── middleware/         # Auth & Role-Based Access Middleware
│   │   ├── purchasing/         # Service Dashboard & Transaksi Pembelian
│   │   ├── supplier/           # Service Manajemen Supplier
│   │   └── user/               # Service User CRUD & Profile
│   └── run.go                  # Setup Fiber app, CORS, & Registrasi Router
├── go.mod
├── go.sum
├── main.go                     # Entry point utama aplikasi
└── .env                        # File konfigurasi environment
```

---

## Kebijakan Hak Akses (RBAC)

Sistem membedakan hak akses berdasarkan klaim `role` di dalam JWT token:

| Fitur / Endpoint        | Method         | Route                   | Admin | Staff (Non-Admin) |
| :---------------------- | :------------- | :---------------------- | :---: | :---------------: |
| **Auth** (Login)        | `POST`         | `/auth`                 |  ✅   |        ✅         |
| **Auth** (Logout)       | `GET`          | `/auth/logout`          |  ✅   |        ✅         |
| **Dashboard**           | `GET`          | `/purchasing/dashboard` |  ✅   |        ✅         |
| **User CRUD**           | `*`            | `/user/*`               |  ✅   |  ❌ (Forbidden)   |
| **My Profile**          | `GET/PUT`      | `/user/my-profile`      |  ✅   |        ✅         |
| **Supplier** (Read)     | `GET`          | `/supplier` & `/:id`    |  ✅   |        ✅         |
| **Supplier** (Write)    | `POST/PUT/DEL` | `/supplier/*`           |  ✅   |  ❌ (Forbidden)   |
| **Item** (Read)         | `GET`          | `/item` & `/:id`        |  ✅   |        ✅         |
| **Item** (Write)        | `POST/PUT/DEL` | `/item/*`               |  ✅   |  ❌ (Forbidden)   |
| **Purchasing** (Read)   | `GET`          | `/purchasing` & `/:id`  |  ✅   |        ✅         |
| **Purchasing** (Create) | `POST`         | `/purchasing`           |  ✅   |        ✅         |
| **Purchasing** (Write)  | `PUT/DELETE`   | `/purchasing/:id`       |  ✅   |  ❌ (Forbidden)   |

---

## Cara Setup & Menjalankan Aplikasi

### 1. Prasyarat

Pastikan Anda sudah menginstal:

- [Go (Golang) v1.24+](https://go.dev/dl/)
- MySQL Server yang sedang aktif.

### 2. Salin Konfigurasi Environment (`.env`)

Buat file `.env` di root direktori dengan contoh konfigurasi berikut:

```env
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=your_db_host
DB_PORT=your_db_port
DB_NAME=your_db_name
TOKEN_SECRET=your_token_secret
LIFE_SPAN=your_life_span
```

### 3. Import Database

Import database dari file `purchasing_db.sql` ke MySQL Server Anda.

username: admin
password: admin123

### 4. Jalankan Aplikasi

Backend menggunakan `AutoMigrate` dari GORM, sehingga tabel database akan dibuat secara otomatis saat pertama kali dijalankan.

```bash
go run main.go
```

Aplikasi akan berjalan secara default di port `:8080` (atau sesuai konfigurasi env).

---

## Fitur Khusus

### 1. Filter Tanggal Dashboard

Endpoint `GET /purchasing/dashboard` mendukung filter opsional `startDate` dan `endDate` (`YYYY-MM-DD`):

- Jika tidak dikirim oleh client, data otomatis diambil dari **tanggal terbaru sampai 10 hari ke belakang**.
- Query data dan kalkulasi ringkasan stok diselesaikan secara efisien langsung dalam satu kali proses query data di level database.

### 2. Keamanan User Profile

Ubah password via `/user/my-profile` (PUT) mewajibkan pencocokan `old_password` untuk menjamin keamanan akun dari pembajakan sesi.
