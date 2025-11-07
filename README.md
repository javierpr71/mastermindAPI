# Mastermind API

![Go Version](https://img.shields.io/badge/Go-1.17+-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)

A REST API implementation of the classic Mastermind game built with Go, Redis, and Docker.

## About Mastermind

Mastermind is a code-breaking game where one player (the codemaker) creates a secret code, and another player (the codebreaker) tries to guess it. In this digital version:

- The API generates a random 4-color secret code using colors: **R**ed, **B**lue, **Y**ellow, **G**reen, **W**hite, **O**range
- Players have up to 10 rounds to guess the correct code
- After each guess, the API provides feedback:
  - **Blacks**: Number of colors that are correct and in the right position
  - **Whites**: Number of colors that are correct but in the wrong position
- Win by getting 4 blacks (all colors correct and in correct positions)

## Features

- RESTful API with JSON responses
- Persistent game state using Redis
- Docker support for easy deployment
- Comprehensive API documentation
- Postman collection for testing

## Prerequisites

- **Go 1.17+** (for local development)
- **Redis instance** (for data persistence)
- **Docker** (optional, for containerized deployment)
- **Make** (for build automation)

## Quick Start

### Using Docker (Recommended)

1. **Clone the repository:**
   ```bash
   git clone https://github.com/javierpr71/mastermindAPI.git
   cd mastermindAPI
   ```

2. **Set up environment variables:**
   ```bash
   cp .env.example .env
   # Edit .env with your settings
   ```

3. **Run with Docker:**
   ```bash
   make docker
   docker run -p 8080:8080 --env-file .env mastermind:latest
   ```

### Local Development

1. **Install dependencies:**
   ```bash
   go mod download
   ```

2. **Start Redis locally:**
   ```bash
   # Using Docker
   docker run -d -p 6379:6379 redis:alpine
   
   # Or install Redis locally
   # macOS: brew install redis && brew services start redis
   # Ubuntu: sudo apt-get install redis-server
   ```

3. **Configure environment variables:**
   ```bash
   export PORT=8080
   export REDIS=localhost:6379
   ```

4. **Run the application:**
   ```bash
   make run
   # or: go run main.go
   ```

## Configuration

Configure the following environment variables:

| Variable | Description | Example | Required |
|----------|-------------|---------|----------|
| `PORT` | Port where the API will listen | `8080` | Yes |
| `REDIS` | Redis instance URL in format `<host>:<port>` | `localhost:6379` | Yes |

## Building

### Build for Current Platform
```bash
make build
```

### Build for All Platforms
```bash
make build-all
```
This creates binaries in the `bin/` folder for:
- `mastermind-darwin` (macOS)
- `mastermind-linux` (Linux)
- `mastermind-win.exe` (Windows)

### Clean Build Artifacts
```bash
make clean
```

## API Endpoints

Base URL: `http://localhost:8080`

### 1. Create New Game

**Endpoint:** `POST /newgame`

**Description:** Creates a new Mastermind game with a randomly generated secret code.

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000"
}
```

**Status Codes:**
- `200 OK`: Game created successfully
- `500 Internal Server Error`: Failed to create game

---

### 2. Make a Guess

**Endpoint:** `POST /round`

**Description:** Submit a guess for an existing game.

**Request Body:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "guess": "RBYG"
}
```

**Parameters:**
| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | string | Yes | UUID of the game |
| `guess` | string | Yes | 4-character guess using colors R,B,Y,G,W,O |

**Response:**
```json
{
  "round": 1,
  "guess": "RBYG",
  "whites": 2,
  "blacks": 1
}
```

**Status Codes:**
- `200 OK`: Guess processed successfully
- `400 Bad Request`: Invalid request format
- `404 Not Found`: Game not found
- `500 Internal Server Error`: Processing error

---

### 3. Get Game Status

**Endpoint:** `GET /status/{id}`

**Description:** Retrieve the current status of a game including all previous rounds.

**Path Parameters:**
| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | string | Yes | UUID of the game |

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "rounds": [
    {
      "round": 1,
      "guess": "RBYG",
      "whites": 2,
      "blacks": 1
    }
  ],
  "result": ""
}
```

**Game End Results:**
- `"You Win!"`: Player guessed the correct code
- `"You Loose!"`: Player used all 10 rounds without success
- `""`: Game is still in progress

**Status Codes:**
- `200 OK`: Status retrieved successfully
- `404 Not Found`: Game not found
- `500 Internal Server Error`: Retrieval error

## Game Rules

1. **Secret Code**: 4 colors from R(ed), B(lue), Y(ellow), G(reen), W(hite), O(range)
2. **Guessing**: Submit 4-character strings (e.g., "RBYG")
3. **Feedback**:
   - **Blacks**: Correct color in correct position
   - **Whites**: Correct color in wrong position
4. **Winning**: Get 4 blacks (all positions correct)
5. **Losing**: Use all 10 rounds without winning

## Testing

A Postman collection is included in `./tests/postman/MasterMind.postman_collection.json` for easy API testing.

### Example Game Flow

1. **Start a new game:**
   ```bash
   curl -X POST http://localhost:8080/newgame
   ```

2. **Make your first guess:**
   ```bash
   curl -X POST http://localhost:8080/round \
     -H "Content-Type: application/json" \
     -d '{"id":"your-game-id","guess":"RBYG"}'
   ```

3. **Check game status:**
   ```bash
   curl http://localhost:8080/status/your-game-id
   ```

## Deployment

### Docker Deployment

1. **Build the Docker image:**
   ```bash
   docker build -t mastermind:latest .
   ```

2. **Run the container:**
   ```bash
   docker run -d \
     -p 8080:8080 \
     -e PORT=8080 \
     -e REDIS=your-redis-host:6379 \
     mastermind:latest
   ```

### Binary Deployment

1. **Build for target platform:**
   ```bash
   make build-all
   ```

2. **Deploy the appropriate binary:**
   - Copy binary from `bin/` folder to target server
   - Set environment variables
   - Run the binary

## Development

### Project Structure
```
├── main.go              # Application entry point
├── controllers/         # HTTP handlers and business logic
├── models/             # Data structures
├── repository/         # Data access layer
├── handler/            # Request handlers
├── middlewares/        # HTTP middlewares
├── responses/          # Response utilities
├── tests/              # Test files and Postman collection
├── Dockerfile          # Docker configuration
├── Makefile           # Build automation
└── README.md          # This file
```

### Adding New Features

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

