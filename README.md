# Bot Blocked Ticket Handler

This script searches for tickets with the type `bot_blocked` and closes them with the tag "Uninformative".

## Setup

To run the script, create a `.env` file in the root directory and specify the following environment variables:

```ini
KEY="Basic token"          # Obtainable from the Infobip account
AGENT_ID="11111222222333334444" # ID of the agent with supervisor rights
TAG="Uninformative"        # Tag with which the ticket will be closed
INTERVAL=60                # Interval in seconds
```

## Building the Project

### Building on Windows

To build the project for Linux on a Windows machine, use the following command:

```bash
$Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go build -o bot_blocked_handler cmd/main.go
```

## Running the Script

After building, you can run the script using:

```bash
./bot_blocked_handler
```

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
