---
name: crud-module-generator
description: Gunakan skill ini ketika user meminta membuat service CRUD baru, atau fitur baru yang melibatkan fetch data, get data, tambah, edit, hapus di project task-be ini. Skill ini memastikan konsistensi arsitektur dan pola kode yang sudah ada.
---

# CRUD Module Generator — Task BE

Skill ini membantu membuat service CRUD baru yang **konsisten** dengan arsitektur dan coding pattern project `task-be` (Go, Fiber, GORM, MySQL).

## Tech Stack Project

| Teknologi | Versi/Detail |
| --------- | ------------ |
| Language  | Go (Golang)  |
| Framework | Fiber v2     |
| ORM       | GORM         |
| Database  | MySQL        |
| Auth      | JWT Middleware |

---

## Arsitektur Project

Seluruh file source code backend berada di dalam folder `app/`:

```
app/
├── db/
│   └── connection.go       # Koneksi DB & Auto-Migration
├── helper/
│   ├── blacklisttoken/     # Blacklist JWT token repository & model
│   ├── hashpassword.go     # Password hashing utilities
│   ├── message.go          # Standard response messages (Create, Update, Delete)
│   ├── timehelper.go       # UTC time helper
│   └── token.go            # JWT token utilities
├── model/
│   ├── base.go             # BaseModel (ULID) & DateTime (timestamps)
│   ├── item.go             # Contoh model Item
│   └── <module>.go         # Definisikan struct Model database baru di sini
├── service/
│   ├── base.go             # generic JSON handler helper
│   ├── validator.go        # generic BindAndValidate body validator
│   ├── middleware/         # Auth Middleware
│   ├── supplier/           # Contoh module CRUD
│   │   ├── repository.go   # Data Access / Query GORM
│   │   ├── request.go      # Request Payload Validator Struct
│   │   ├── response.go     # Response DTO Struct
│   │   ├── router.go       # Fiber Routing & Endpoints
│   │   └── service.go      # Business Logic
│   └── <module>/           # Folder modul baru yang akan dibuat
│       ├── repository.go
│       ├── request.go
│       ├── response.go
│       ├── router.go
│       └── service.go
├── run.go                  # Setup & Inisialisasi Fiber & Route Register
└── main.go (root)          # Entry point utama (memanggil app.Run())
```

---

## Pola Kode yang WAJIB Diikuti

### 1. Definisi Model Database (`app/model/<module>.go`)
Setiap model database baru wajib:
- Meng-embed `model.BaseModel` untuk generate ID otomatis menggunakan **ULID** (pada `BeforeCreate` hook).
- Meng-embed `model.DateTime` untuk field `CreatedAt`, `UpdatedAt`, dan `DeletedAt` otomatis (soft delete support).
- Menentukan nama table secara eksplisit lewat method `TableName() string`.

*Contoh:*
```go
package model

type <Module> struct {
	BaseModel
	Name        string    `gorm:"type:varchar(100)"`
	Description string    `gorm:"type:varchar(255)"`
	DateTime
}

func (<Module>) TableName() string {
	return "<table_name>"
}
```

### 2. Auto-Migration Model Baru (`app/db/connection.go`)
Setelah membuat file model, daftarkan model tersebut ke dalam pemanggilan `AutoMigrate` pada fungsi `init()` di `app/db/connection.go`.

```go
func init() {
	db := Connection()
	err := db.AutoMigrate(
		model.Supplier{},
		model.Item{},
		// ... model lainnya
		model.<Module>{}, // Tambahkan model baru di sini
	)
	if err != nil {
		panic(err)
	}
}
```

### 3. Skema Request & Validation (`app/service/<module>/request.go`)
Gunakan tag `validate` dari `github.com/go-playground/validator`. Gunakan format penamaan `<module>Request` secara `camelCase` / `pascalCase`.

```go
package <module>

type <module>Request struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
```

### 4. Skema Response DTO (`app/service/<module>/response.go`)
Gunakan model custom untuk Response DTO agar detail internal database tidak bocor. Embed `model.BaseModel` dan `model.DateTime`. Tentukan `TableName` agar GORM dapat memetakan hasil query langsung ke struct response.

```go
package <module>

import "task-be/app/model"

type <module>Response struct {
	model.BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	model.DateTime
}

func (<module>Response) TableName() string {
	return "<table_name>"
}
```

### 5. Repository Pattern dengan GORM (`app/service/<module>/repository.go`)
Pisahkan interface dan implementation struct. Berikan akses DB GORM.

```go
package <module>

import (
	"task-be/app/model"
	"gorm.io/gorm"
)

type Repository interface {
	create(data model.<Module>) error
	update(id string, name string, description string) error
	delete(id string) error
	getAll() []<module>Response
	getById(id string) <module>Response
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) create(data model.<Module>) error {
	return r.db.Create(&data).Error
}

func (r *repository) update(id string, name string, description string) error {
	return r.db.Where("id = ?", id).Model(&model.<Module>{}).Updates(map[string]interface{}{
		"name":        name,
		"description": description,
	}).Error
}

func (r *repository) delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.<Module>{}).Error
}

func (r *repository) getAll() []<module>Response {
	var res []<module>Response
	r.db.Find(&res)
	return res
}

func (r *repository) getById(id string) <module>Response {
	var res <module>Response
	r.db.Where("id = ?", id).Find(&res)
	return res
}
```

### 6. Service / Business Logic (`app/service/<module>/service.go`)
Koordinasikan repository, lakukan mapping dari Request payload ke database Model, dan panggil repository methods.

```go
package <module>

import "task-be/app/model"

type Service interface {
	create(req <module>Request) error
	update(id string, req <module>Request) error
	delete(id string) error
	getAll() []<module>Response
	getById(id string) <module>Response
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) create(req <module>Request) error {
	return s.repository.create(model.<Module>{
		Name:        req.Name,
		Description: req.Description,
	})
}

func (s *service_) update(id string, req <module>Request) error {
	return s.repository.update(id, req.Name, req.Description)
}

func (s *service_) delete(id string) error {
	return s.repository.delete(id)
}

func (s *service_) getAll() []<module>Response {
	return s.repository.getAll()
}

func (s *service_) getById(id string) <module>Response {
	return s.repository.getById(id)
}
```

### 7. Routing dengan Fiber (`app/service/<module>/router.go`)
- Gunakan `service.BindAndValidate(c, &req)` untuk parser & validation (fungsi helper global).
- Gunakan `service.JSON(c, err, helper.GetMessage.Create/Update/Delete)` untuk standard format response actions.
- Gunakan parameter `auth` dan `admin` middleware (`fiber.Handler`) untuk mengontrol hak akses API. Terapkan `auth` pada grouping utama router, dan `admin` pada endpoint spesifik write (POST/PUT/DELETE) yang hanya boleh diakses Admin.

```go
package <module>

import (
	"task-be/app/helper"
	"task-be/app/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB, auth fiber.Handler, admin fiber.Handler) {
	Repository := NewRepository(db)
	service_ := NewService(Repository)
	route := app.Group("/<module>", auth)

	route.Post("", admin, func(c *fiber.Ctx) error {
		var req <module>Request
		if service.BindAndValidate(c, &req) {
			return nil
		}
		err := service_.create(req)
		return service.JSON(c, err, helper.GetMessage.Create)
	})

	route.Put("/:id", admin, func(c *fiber.Ctx) error {
		id := c.Params("id")
		var req <module>Request
		if service.BindAndValidate(c, &req) {
			return nil
		}
		err := service_.update(id, req)
		return service.JSON(c, err, helper.GetMessage.Update)
	})

	route.Delete("/:id", admin, func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := service_.delete(id)
		return service.JSON(c, err, helper.GetMessage.Delete)
	})

	route.Get("", func(c *fiber.Ctx) error {
		data := service_.getAll()
		return c.Status(200).JSON(data)
	})

	route.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		data := service_.getById(id)
		return c.Status(200).JSON(data)
	})
}
```

---

## Langkah Membuat Modul Baru

Ketika user meminta modul CRUD baru (misal: `category`), ikuti langkah-langkah berikut:

1. **Buat Model Database**: Buat file `app/model/category.go` dan definisikan struct `Category` beserta field, tag gorm, dan method `TableName()`.
2. **Daftarkan Model**: Daftarkan `model.Category{}` ke `AutoMigrate` pada `app/db/connection.go`.
3. **Buat Folder Modul**: Buat folder `app/service/category`.
4. **Implementasikan File Modul**:
   - `app/service/category/request.go`
   - `app/service/category/response.go`
   - `app/service/category/repository.go`
   - `app/service/category/service.go`
   - `app/service/category/router.go`
5. **Daftarkan Route**: Di `app/run.go`, import package category dan panggil `category.Router(app, db, middleware.Auth(), middleware.IsAdmin())` (atau sesuaikan hak akses sesuai kebutuhan bisnis).

---

## Aturan Penting

1. **Jangan** melanggar layering arsitektur (Router -> Service -> Repository -> Database).
2. **Selalu** gunakan `service.BindAndValidate` untuk validasi JSON payload.
3. **Selalu** gunakan `helper.GetMessage.Create/Update/Delete` dengan `service.JSON` untuk action response yang konsisten.
4. **Gunakan** ULID secara otomatis untuk primary key data baru lewat `model.BaseModel`.
5. **Jangan** hardcode response error database ke user secara telanjang; gunakan format standard helper.
