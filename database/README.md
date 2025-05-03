# 📚 Go Book API with MongoDB

This project is a simple RESTful API built with **Go**, **Gin**, and **MongoDB**, allowing users to create, read, update, and delete book records.

---

## 🚀 Features

- Create, Read, Update, and Delete (CRUD) operations
- MongoDB integration using the official Go driver
- Docker support for MongoDB
- Clean project structure

---

## 📦 Tech Stack

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [MongoDB](https://www.mongodb.com/)
- [MongoDB Go Driver](https://go.mongodb.org/mongo-driver)

---

## 🧪 API Endpoints

| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| `GET`  | `/api/books`     | List all books      |
| `POST` | `/api/books`     | Create a new book   |
| `PUT`  | `/api/books/:id` | Update a book       |
| `DELETE` | `/api/books/:id` | Delete a book       |
