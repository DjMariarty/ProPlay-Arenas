# Docker Setup

Инструкция по запуску проекта в Docker.

## Требования

- Docker Desktop (или Docker Engine + Docker Compose)
- Минимум 4GB свободной RAM

## Быстрый старт

1. **Запустить все сервисы:**
```bash
docker-compose up --build
```

2. **Запустить в фоновом режиме:**
```bash
docker-compose up -d --build
```

3. **Остановить все сервисы:**
```bash
docker-compose down
```

4. **Остановить с удалением данных:**
```bash
docker-compose down -v
```

## Порты сервисов

- **Gateway**: `http://localhost:8085` (главная точка входа)
- User Service: `http://localhost:8080`
- Venue Service: `http://localhost:8082`
- Payment Service: `http://localhost:8084`
- Reservation Service: `http://localhost:8081`
- Kafka: `localhost:9092`

## Базы данных PostgreSQL

- User DB: `localhost:5433`
- Venue DB: `localhost:5434`
- Payment DB: `localhost:5435`
- Reservation DB: `localhost:5436`

## Переменные окружения

Создайте файл `.env` в корне проекта (опционально):

```env
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
```

Если файл не создан, будет использовано значение по умолчанию.

## Полезные команды

### Просмотр логов
```bash
# Все сервисы
docker-compose logs -f

# Конкретный сервис
docker-compose logs -f gateway
docker-compose logs -f user-service
```

### Пересборка конкретного сервиса
```bash
docker-compose up -d --build gateway
```

### Проверка статуса
```bash
docker-compose ps
```

### Очистка
```bash
# Удалить все контейнеры и volumes
docker-compose down -v

# Удалить неиспользуемые образы
docker image prune -a
```

## Структура сервисов

- **gateway** - API Gateway (порт 8085)
- **user-service** - Сервис пользователей (порт 8080)
- **venue-service** - Сервис площадок (порт 8082)
- **payment-service** - Сервис платежей (порт 8084)
- **reservation-service** - Сервис бронирований (порт 8081)
- **kafka** - Message broker (порт 9092)
- **zookeeper** - Координатор для Kafka (порт 2181)
- **user-db, venue-db, payment-db, reservation-db** - PostgreSQL базы данных

## Troubleshooting

### Проблемы с сетью при загрузке образов

Если возникают проблемы с загрузкой образов из Docker Hub:

```bash
# Загрузить образы по одному
docker pull alpine:latest
docker pull golang:1.25-alpine
docker pull postgres:15
docker pull confluentinc/cp-zookeeper:7.5.3
docker pull confluentinc/cp-kafka:7.5.3

# Затем собрать без загрузки
docker-compose build --pull=never
docker-compose up -d
```

### Проблемы с правами доступа

Убедитесь, что Docker Desktop запущен и имеет необходимые права.

### Проблемы с портами

Если порты заняты, измените их в `docker-compose.yaml` в секции `ports`.
