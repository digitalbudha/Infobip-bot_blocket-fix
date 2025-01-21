<h2>Script searches for tickets with type bot_blocked and closes them with the tag "Uninformative"</h2>

<h3>To run the script, create a .env file and specify the following:</h3>

```Bash
KEY = "Basic token" # Can be obtained from the infobip account
AGENT_ID = "11111222222333334444" # ID of the agent with supervisor rights
TAG = "Uninformative" # Tag with which the ticket will be closed
INTERVAL = 60 # In seconds
```

<h3>Building on Windows</h3>

```Bash
$Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go build .\main.go
```

<h6>Coloring logs</h6>

```GO
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Error converting string to number ", err)
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Unmarshalling error: ", err)
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Error converting to JSON: ", err)
```
