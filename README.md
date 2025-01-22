# Bot Blocked Ticket Handler

This script searches for tickets with the type `bot_blocked` and closes them with the tag "Uninformative".

## Setup

To run the script, create a `.env` file in the root directory and specify the following environment variables:

```ini
KEY="Basic token"          # Obtainable from the Infobip account
AGENT_ID="11111222222333334444" # ID of the agent with supervisor rights
TAG="Uninformative"        # Tag with which the ticket will be closed
INTERVAL=60                # Interval in seconds
URL="https://your_server.api.infobip.com/ccaas/1/conversations" # URL of the Infobip server
```

## Building the Project (Windows)

To build the project for Linux on a Windows machine, use the following command:

```bash
$Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go build -o bot_blocked_handler cmd/main.go
```

## Running the Script

After building, you can run the script using:

```bash
./bot_blocked_handler
```

## Running with Docker Compose

To run the binary in a Docker container using Docker Compose, follow these steps:

1. **Create a Dockerfile** in the root directory:

    ```dockerfile
    FROM golang:alpine

    WORKDIR /app

    COPY bot_blocked_handler /app/bot_blocked_handler
    COPY .env /app/.env

    CMD ["./bot_blocked_handler"]
    ```

2. **Create a `docker-compose.yml` file** in the root directory:

    ```yaml
    version: '3.8'

    services:
      bot_handler:
        build: .
        container_name: bot_blocked_handler
        restart: unless-stopped
        environment:
          - KEY=${KEY}
          - AGENT_ID=${AGENT_ID}
          - TAG=${TAG}
          - INTERVAL=${INTERVAL}
          - URL=${URL}
    ```

3. **Deploy the container** on your server:

    - Ensure Docker and Docker Compose are installed on your server.
    - Run the following command to start the service:

    ```bash
    docker-compose up -d
    ```

    This command will build the Docker image and start the container in detached mode.

## Code Structure

The project is organized into several packages to improve maintainability:

- **cmd/main.go**: The entry point of the application.
- **internal/config**: Handles loading and accessing environment variables.
- **internal/conversations**: Manages conversation-related operations.
- **internal/messages**: Handles message-related operations.
- **internal/tags**: Manages tagging and closing of conversations.
- **internal/models**: Contains data models used across the application.

## Logging 

The script uses colored logs to indicate different types of messages. Examples include:

```go
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Error converting string to number ", err)
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Unmarshalling error: ", err)
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Error converting to JSON: ", err)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
