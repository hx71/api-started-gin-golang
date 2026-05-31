# API Starter with Gin & Golang (Service & Repository Pattern)

A REST API boilerplate built with the Gin framework using Clean Architecture principles (Service & Repository Pattern) in Golang.

## 🚀 Getting Started

### 1. Clone Project
```bash
git clone git@github.com:hx71/api-started-gin-golang.git
cd api-started-gin-golang
```

### 2. Setup Database
Create a database in your preferred DBMS (PostgreSQL is recommended).

### 3. Setup Configuration
Copy the `.env.example` file to create your local `.env` configuration:
```bash
cp .env.example .env
```
Ensure you update the database credentials inside the `.env` file to match your local setup.

### 4. Run the Project
You can easily start the development server using the Makefile:
```bash
make run
```

---

## 🌐 Endpoints

**Main API Endpoint:**
```
http://localhost:1234/api/v1
```

**Swagger API Documentation:**
```
http://localhost:1234/swagger/index.html
```

---

## 🧪 Testing

The project uses table-driven unit testing with a mock-driven approach for the Usecase/Service layer, ensuring that tests run extremely fast and isolated from external dependencies.

To execute the entire test suite:
```bash
make test
```