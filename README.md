<b>Скрипт ищет обращения с типом <U>BOT_BLOCKED</U> и закрывает их с тегом «Неинформативно»</b>
<br>Для запуска скрипта необходимо создать .env файл, в нем указать:

```Bash
KEY = "Basic токен" # Получить можно на аккаунте infobip
AGENT_ID = "11111222222333334444" # ID агента с правами супервайзера
TAG = "Неинформативно" # Тег с которым будет закрываться обращение
INTERVAL = 60 # В секундах
```
#
$Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go build .\main.go

#
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Ошибка преобразования строки в число ", err)
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Ошибка демаршалирования: ", err)
log.Println(" — " + color.RedString("WARNING") + " >>> " + "Ошибка при преобразовании в JSON: ", err)