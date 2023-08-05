<h2>Скрипт ищет обращения с типом <b>bot_blocked</b> и закрывает их с тегом «Неинформативно»</h2>

<h3>Для запуска скрипта необходимо создать .env файл, в нем указать:</h3>

```Bash
KEY = "Basic токен" # Получить можно на аккаунте infobip
AGENT_ID = "11111222222333334444" # ID агента с правами супервайзера
TAG = "Неинформативно" # Тег с которым будет закрываться обращение
INTERVAL = 60 # В секундах
```

<h3>Билдим из под windows</h3>

```Bash
$Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go build .\main.go
```

<h3>Красим логи</h3>

```GO
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Ошибка преобразования строки в число ", err)
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Ошибка демаршалирования: ", err)
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Ошибка при преобразовании в JSON: ", err)
```
