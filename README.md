# REST API Monitoring System

A lightweight and extensible **REST API monitoring service** built with **Go**, **PostgreSQL**, and **Docker Compose**.  
This system periodically checks API endpoints, logs their response status and time, and sends **Telegram alerts** when an endpoint is down or returns an unexpected status code.

---

## Features

- **Automated scheduler** that checks endpoints at custom intervals  
- **PostgreSQL database** to store endpoint configurations and check logs  
- **RESTful API** for managing endpoints  
- **Telegram bot integration** for real-time alerts  
- **Dockerized setup** for easy deployment  
- Optional **live reload** during development with [Air](https://github.com/cosmtrek/air)

---

## Tech Stack

| Component      | Technology        |
|----------------|------------------|
| Language       | Go (Golang)  |
| Database       | PostgreSQL 15 |
| Deployment     | Docker & Docker Compose |
| Notification   | Telegram Bot API |

---

## Project Structure

```
api-monitoring/
│
├── config/           # Load environment variables
├── controllers/      # Auth and Endpoint handler
├── database/         # Database initialization
├── middleware/       # Bridge between applications
├── models/           # GORM models (Endpoint, CheckLog, etc.)
├── services/         # Scheduler, Telegram, and core logic
├── utils/            # Helper functions
├── main.go           # Application entry point
├── Dockerfile        # Build instructions for the app
├── docker-compose.yml# Multi-container setup (app + db)
└── README.md         # Documentation
```

---

## Configuration

Create a `.env` file in the project root:

```env
PORT=8080
DATABASE_URL=postgres://postgres:postgres@db:5432/apimonitor?sslmode=disable
JWT_SECRET=your_jwt_secret
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
TELEGRAM_CHAT_ID=your_telegram_chat_id
```

### Getting Telegram Bot Token
1. Open Telegram → search for [@BotFather](https://t.me/BotFather)  
2. Type `/newbot` and follow the instructions  
3. Copy the generated **bot token**

### For A Personal Telegram Chat ID
1. Open Telegram → search for `@userinfobot`
2. Type `/start`  
3. You’ll get your **Chat ID**

### For A Group Telegram Chat ID
1. Add your bot to the group. 
2. Send a message in that group (any message).
3. Then visit this URL in your browser (replace YOUR_BOT_TOKEN with your real bot token): https://api.telegram.org/botYOUR_BOT_TOKEN/getUpdates
4. You’ll get a JSON response like this:
```json
{
  "ok": true,
  "result": [
    {
      "message": {
        "chat": {
          "id": -1001234567890,
          "title": "My Group",
          "type": "supergroup"
        }
      }
    }
  ]
}
```
The value of "id" (like -1001234567890) is your group chat ID (group chat IDs always start with -100).

---

## Running the Project

### 1. Build and run using Docker Compose
```bash
docker-compose up --build
```

This will:
- Build the Go application
- Start PostgreSQL
- Run migrations (if defined)
- Launch the monitoring scheduler

### 2. Verify running containers
```bash
docker ps
```

You should see:
```
api-monitoring-app-1   Up ...
api-monitoring-db-1    Up ...
```

### 3. View logs
```bash
docker logs api-monitoring-app-1
```

---

## Development with Live Reload

To automatically rebuild and reload when Go files change:

1. Install [Air](https://github.com/cosmtrek/air):
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

2. Update `Dockerfile`:
   ```dockerfile
   RUN go install github.com/cosmtrek/air@latest
   CMD ["air"]
   ```

3. Run again:
   ```bash
   docker-compose up --build
   ```

Now the app reloads automatically whenever you edit `.go` files

---

## API Endpoints

| Method | Endpoint        | Description            |
|--------|-----------------|------------------------
| POST   | `/api/register` | Register new user     |
| GET    | `/api/login`    | Login with JWT        |
| POST   | `/api/endpoints` | Add new monitored URL  |
| GET    | `/api/endpoints` | List all endpoints     |
| GET    | `/api/endpoints/:id/logs` | List all logs depends on endpoint id   |

---

## How It Works

1. Scheduler runs every **10 seconds** (customizable)
2. Each endpoint is checked asynchronously:
   - Send HTTP request
   - Measure response time
   - Compare status code with expected
3. Results are logged in the database
4. Telegram alert is sent if:
   - The endpoint is down
   - Status code doesn’t match expectation

---

## Example Telegram Alert

**Message when endpoint is down:**
```
<b>API DOWN</b>
URL: https://api.example.com/health
Error: connection timeout
Response Time: 1234 ms
```

**Message when unexpected status:**
```
<b>Unexpected Status</b>
URL: https://api.example.com/health
Got: 500 Expected: 200
Response Time: 980 ms
```

---

## Troubleshooting

| Issue | Solution |
|-------|-----------|
| `.env` not found | Make sure it exists and mapped in `docker-compose.yml` |
| App container stops immediately | Check `docker logs api-monitoring-app-1` for error |
| Telegram alert not received | Verify your bot is added to chat and chat ID is correct |
| Database not reachable | Ensure service name `db` matches in `DATABASE_URL` |

---

## Author

**Neville Jeremy**  
Informatics Student at Universitas Sanata Dharma, Yogyakarta  
Focus: Backend Engineering and System Architecture  
GitHub: [@nepile](https://github.com/nepile)

---

## License

This project is licensed under the [MIT License](LICENSE).

---

### If you like this project, give it a star on GitHub!

> _“Monitor smarter. Alert faster. Sleep better.”_
