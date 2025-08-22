# 🔐 PBAC – Policy Based Access Control (Go + Fiber + MySQL)

A simple **Policy Based Access Control (PBAC)** system built with:

* [Go](https://golang.org/)
* [Fiber](https://gofiber.io/) (fast HTTP web framework)
* [GORM](https://gorm.io/) (ORM for Go, with AutoMigrate)
* [MySQL](https://www.mysql.com/)

---

## 🚀 Features

* Define **Users**, **Resources**, and **Actions**
* Create **Policies** that either `allow` or `deny` specific actions on resources for users
* Access control check endpoint (`/check`)
* Auto database migration (no need to manually create tables)
* Lightweight and easy to extend

---

## 📂 Project Structure

```
pbac/
├── db/
│   └── db.go          # Database connection + AutoMigrate
├── models/
│   └── models.go      # GORM models (User, Resource, Action, Policy)
├── pbac/
│   └── access.go      # CheckAccess logic
├── main.go            # Fiber HTTP routes
├── go.mod
└── README.md
```

---

## ⚡ Setup & Run

### 1. Install dependencies

```bash
go mod tidy
```

### 2. MySQL setup

Create database and user (if not already):

```sql
CREATE DATABASE pbac;
CREATE USER 'user'@'localhost' IDENTIFIED BY '';
GRANT ALL PRIVILEGES ON pbac.* TO 'user'@'localhost';
FLUSH PRIVILEGES;
```

### 3. Run the app

```bash
go run .
```

If successful:

```
✅ DB connected & migrated
```

---

## 📡 API Endpoints

### Users

* `POST /users` → create user
* `GET /users` → list users

### Resources

* `POST /resources` → create resource
* `GET /resources` → list resources

### Actions

* `POST /actions` → create action
* `GET /actions` → list actions

### Policies

* `POST /policies` → create policy
* `GET /policies` → list policies

### Access Check

* `POST /check` → check if a user can perform an action on a resource

**Example request:**

```json
POST /check
{
  "user_id": 1,
  "resource_id": 2,
  "action_id": 3
}
```

**Example response:**

```json
{ "allowed": true }
```

---

## 🛠️ Future Improvements

* Role-based policies (RBAC on top of PBAC)
* Condition-based rules (time, IP, etc.)
* JWT authentication for users

---

👉 Do you want me to also include **example `curl` commands** in the README so someone can test all the routes right away?
