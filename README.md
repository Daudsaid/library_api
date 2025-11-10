# 📚 Library API (Golang + Gin)

A simple RESTful API for managing a library's books.  
Built in **Go** using the **Gin** framework.

This project demonstrates core backend concepts:

- Structs & JSON binding
- Routing & path parameters
- Query parameters
- Updating state in memory
- REST API design
- Handling success & error responses

---

## 🚀 Tech Stack

| Component | Technology |
|----------|------------|
| Language | Go (Golang) |
| Web Framework | Gin |
| Data Storage | In-Memory (Slice) |
| API Format | JSON |
| Testing Tool | cURL / HTTP Clients |

---

## 📂 Project Structure

library_api/
│ main.go
│ go.mod
│ go.sum
│ body.json (example input)

---

## 📡 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/books` | Get all books |
| GET | `/books/:id` | Get a book by ID |
| POST | `/books` | Create a new book |
| PATCH | `/checkout?id={id}` | Check out a book (quantity - 1) |
| PATCH | `/return?id={id}` | Return a book (quantity + 1) |

---

## 🧪 Example Usage

### Get all books
```bash
curl localhost:1880/books

✨ Author

Daud Abdi
Backend Developer (Go, SQL, Cloud)
London, UK

GitHub: https://github.com/Daudsaid
www.linkedin.com/in/daudabdi0506



