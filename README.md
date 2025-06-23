# 📝 Todo List API with Golang & Gin

Đây là một project API đơn giản viết bằng [Go](https://golang.org/) sử dụng framework [Gin](https://github.com/gin-gonic/gin), kết nối tới MySQL để thực hiện các thao tác **CRUD (Create, Read, Update, Delete)** cho các item trong Todo List.

---

## 📦 Tech Stack

- Golang
- Gin Web Framework
- GORM (ORM cho MySQL)
- MySQL
- dotenv (quản lý biến môi trường)

---

## ⚙️ Biến môi trường

Ứng dụng sử dụng biến môi trường để cấu hình kết nối database:

```env
DB_CONN_STR=user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
