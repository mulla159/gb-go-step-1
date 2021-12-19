# Пример распарсивания переменных окружения в структуру с помощью библиотеки [envconfig](https://github.com/kelseyhightower/envconfig)

# Добавляем переменные окружения
```Bash
export PORT=8080 DB_URL=postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable JAEGER_URL=http://jaeger:16686 SENTRY_URL=http://sentry:9000 KAFKA_BROKER=kafka:9092 SOME_APP_ID=testid SOME_APP_KEY=testkey
```

# Запускаем приложение
```Bash
go run ./main/main.go
```
