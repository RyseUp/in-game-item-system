# 🎮 In-Game Item Management System

This project is a server-side system for managing in-game items, tracking player inventory, and processing transactions. It supports both gRPC and REST APIs via Connect-Go, uses PostgreSQL for data storage, and RabbitMQ for asynchronous transaction logging and event processing. The project is fully containerized using Docker Compose, so you (or any interviewer) can easily set it up and run it locally.

---

## 📌 Features

- **Item Management**
    - Create, retrieve, update, and delete in-game items.
- **Inventory Management**
    - Track player inventory.
    - Update item quantities safely (with concurrency control and prevention of negative balances).
    - Record inventory update history (via InventoryRecord).
- **Transaction Handling**
    - Log each transaction (e.g., purchase, use) with pre- and post-update balances.
    - Publish transaction events asynchronously to RabbitMQ for background processing.
- **RabbitMQ Integration**
    - Offload non-critical processing (logging) from API responses.
- **Containerized Deployment**
    - Use Docker Compose to run PostgreSQL, RabbitMQ, and the Go service seamlessly.

---

## 📌 Prerequisites

Before running the project, ensure you have the following installed on your machine:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Git](https://git-scm.com/)

*No additional software is required since everything runs inside Docker.*

---

## 🚀 Getting Started

### 1️⃣ Clone the Repository

Open your terminal and run:

```sh
git clone https://github.com/YOUR_GITHUB_USERNAME/in-game-item-system.git
cd in-game-item-system

````
### 2️⃣ Configure Environment Variables

### 3️⃣ Run the project using Docker Compose

```sh
docker-compose up --build -d
```
### 4️⃣ Test the APIs using Postman
```sh
http://localhost:50051/api.inventory.v1.InventoryAPI/UserUseItemInInventory
{
  "user_id": "user_001",
  "item_id": "item_001",
  "quantity": 5
}
```