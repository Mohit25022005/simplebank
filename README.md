# 💰 Simple Bank API (Fintech Ledger System)

A backend system built in **Golang** that simulates a real-world banking/ledger platform.  
It supports account management, money transfers, authentication, and scalable deployment.

---

## 🚀 Features

- 🏦 Create and manage bank accounts  
- 💸 Transfer money between accounts (ACID transactions)  
- 🔐 Secure authentication (JWT / PASETO)  
- 🧾 Double-entry ledger system (transfers + entries)  
- ⚡ REST API using Gin  
- 🧪 Unit & integration testing with high coverage  
- 🐳 Dockerized environment  
- ☁️ Deployment-ready (AWS, Kubernetes)  
- 🔁 Background workers with Redis (async tasks)  

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)  
- **Database:** PostgreSQL  
- **Query Tool:** SQLC  
- **Web Framework:** Gin  
- **Auth:** JWT / PASETO  
- **Containerization:** Docker, Docker Compose  
- **CI/CD:** GitHub Actions  
- **Deployment:** AWS (ECR, RDS, EKS)  
- **Async Tasks:** Redis  

---

## 📁 Project Structure

```
.
├── api/            # HTTP handlers (Gin)
├── db/             # SQL queries & generated code (sqlc)
├── util/           # Utility functions (hashing, tokens)
├── worker/         # Background workers (Redis)
├── cmd/            # Entry point of application
├── migrations/     # Database migrations
├── docker/         # Docker configs
└── tests/          # Unit & integration tests
```

---

## ⚙️ Setup & Installation

### 1. Clone the repo
```bash
git clone https://github.com/yourusername/simple-bank.git
cd simple-bank
```

### 2. Start services (Postgres + Redis)
```bash
docker-compose up -d
```

### 3. Run migrations
```bash
make migrateup
```

### 4. Run the server
```bash
make server
```

---

## 🔑 Environment Variables

Create a `.env` file:

```env
DB_SOURCE=postgresql://user:password@localhost:5432/simple_bank?sslmode=disable
TOKEN_SYMMETRIC_KEY=your-secret-key
ACCESS_TOKEN_DURATION=15m
```

---

## 📡 API Endpoints (Sample)

| Method | Endpoint          | Description              |
|--------|------------------|--------------------------|
| POST   | /users           | Create user              |
| POST   | /users/login     | Login user               |
| POST   | /accounts        | Create account           |
| GET    | /accounts/:id    | Get account              |
| POST   | /transfers       | Transfer money           |

---

## 🧠 Key Concepts Implemented

- Database Transactions & Isolation Levels  
- Deadlock Handling & Prevention  
- Double-entry Accounting System  
- Authentication & Authorization Middleware  
- Clean Architecture (Separation of concerns)  
- Scalable Deployment with Kubernetes  

---

## 🧪 Running Tests

```bash
make test
```

---

## 🐳 Docker

Build image:
```bash
docker build -t simple-bank .
```

Run container:
```bash
docker run -p 8080:8080 simple-bank
```

---

## ☁️ Deployment

- Docker image pushed to AWS ECR  
- PostgreSQL hosted on AWS RDS  
- Kubernetes deployment on AWS EKS  
- CI/CD via GitHub Actions  

---

## 📈 Future Improvements

- Rate limiting  
- GraphQL support  
- Multi-currency support  
- Monitoring & observability  

---

## 🙌 Acknowledgements

Inspired by real-world fintech systems and backend engineering practices.

---

---

## ⭐ If you found this useful

Give it a star ⭐ — it helps!