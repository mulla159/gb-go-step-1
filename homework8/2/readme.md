# Пример распарсивания переменных окружения

# Добавляем переменные окружения
```Bash
export PORT=8080 JAEGER_URL=http://jaeger:16686 SENTRY_URL=http://sentry:9000
```

# Запускаем приложение
```Bash
go run ./main/main.go
```
