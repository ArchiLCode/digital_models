# Go WebSocket Messenger

A real-time messaging application built with Go, PostgreSQL, and WebSocket.

## Features

- Real-time messaging using WebSocket
- Multiple chat rooms support
- Message history
- Private and group chats
- PostgreSQL database for persistent storage
- Docker support with automatic migrations

## Prerequisites

- Docker and Docker Compose
- Make sure you have the following environment variables set in your `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=messenger
SERVER_PORT=8080
```

## Setup with Docker

1. Clone the repository

2. Create `.env` file with the required environment variables (see above)

3. Build and start the containers:
```bash
docker-compose up --build
```

The application will be available at `http://localhost:8080`

## Manual Setup (without Docker)

1. Install Go 1.21 or later and PostgreSQL 12 or later
2. Create a PostgreSQL database:
```sql
CREATE DATABASE messenger;
```

3. Run the database migrations (schema.sql)
4. Install dependencies:
```bash
go mod download
```

5. Start the server:
```bash
go run cmd/main.go
```

## API Endpoints

- `POST /api/users` - Create a new user
- `POST /api/chats` - Create a new chat
- `GET /api/chats/{chatID}/messages` - Get messages from a specific chat
- `GET /api/users/{userID}/chats` - Get all chats for a user
- `WS /ws/{userID}/{chatID}` - WebSocket endpoint for real-time messaging

## WebSocket Message Format

Messages sent through WebSocket should follow this format:

```json
{
    "type": "chat_message",
    "content": {
        "chat_id": 1,
        "user_id": 1,
        "content": "Hello, World!"
    }
}
```

## Example Usage

1. Create a user:
```bash
curl -X POST http://localhost:8080/api/users -d '{"username": "john", "password": "secret"}'
```

2. Create a chat:
```bash
curl -X POST http://localhost:8080/api/chats -d '{"name": "General", "is_private": false}'
```

3. Connect to WebSocket:
```javascript
const ws = new WebSocket('ws://localhost:8080/ws/1/1');
ws.onmessage = (event) => {
    console.log('Received:', JSON.parse(event.data));
};
```

## Docker Commands

- Build and start containers:
```bash
docker-compose up --build
```

- Stop containers:
```bash
docker-compose down
```

- View logs:
```bash
docker-compose logs -f
```

- Rebuild a specific service:
```bash
docker-compose up -d --build app
```

## License

MIT 