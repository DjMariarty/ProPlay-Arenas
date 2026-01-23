# API Documentation

Базовый URL: `http://localhost:8085/api`

Все запросы проходят через Gateway на порту 8085.

## Аутентификация

Большинство endpoints требуют JWT токен в заголовке:
```
Authorization: Bearer <your_token>
```

### Публичные endpoints (не требуют авторизации):
- `POST /api/auth/register` - Регистрация
- `POST /api/auth/login` - Вход
- `GET /api/venues` - Список площадок
- `GET /api/venues/:id` - Детали площадки
- `GET /api/venue-types` - Типы площадок

---

## 1. Аутентификация (Auth)

### Регистрация
```http
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "full_name": "John Doe"
}
```

**Примечание:** Поле `full_name` обязательно. Пароль должен быть минимум 8 символов.

**Ответ:**
```json
{
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  },
  "error": null
}
```

### Вход
```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Ответ:**
```json
{
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  },
  "error": null
}
```

---

## 2. Пользователи (Users)

### Получить информацию о текущем пользователе
```http
GET /api/users/me
Authorization: Bearer <token>
```

### Обновить информацию о текущем пользователе
```http
PUT /api/users/me
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "John Updated",
  "email": "newemail@example.com"
}
```

### Стать владельцем площадки
```http
POST /api/users/me/become-owner
Authorization: Bearer <token>
Content-Type: application/json

{
  "phone": "+1234567890"
}
```

### Получить площадки пользователя
```http
GET /api/users/:id/venues
Authorization: Bearer <token>
```

---

## 3. Площадки (Venues)

### Получить список площадок
```http
GET /api/venues?district=Центральный&venue_type=Конференц-зал&hour_price=5000&page=1&limit=10
```

**Query параметры:**
- `district` - район (опционально)
- `venue_type` - тип площадки (опционально)
- `hour_price` - цена за час (опционально)
- `is_active` - активна ли площадка (опционально, true/false)
- `owner_id` - ID владельца (опционально)
- `page` - номер страницы (по умолчанию 1)
- `limit` - количество на странице (по умолчанию 10, максимум 100)

**Ответ:**
```json
{
  "venues": [...],
  "total": 100,
  "page": 1,
  "limit": 10
}
```

### Получить детали площадки
```http
GET /api/venues/:id
```

### Создать площадку
```http
POST /api/venues
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Конференц-зал А",
  "description": "Просторный зал для мероприятий",
  "venue_type": "Конференц-зал",
  "district": "Центральный",
  "address": "ул. Примерная, д. 1",
  "hour_price": 5000,
  "capacity": 50,
  "is_active": true
}
```

### Обновить площадку
```http
PUT /api/venues/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Обновленное название",
  "hour_price": 6000
}
```

### Удалить площадку
```http
DELETE /api/venues/:id
Authorization: Bearer <token>
```

### Получить расписание площадки
```http
GET /api/venues/:id/schedule
```

### Обновить расписание площадки
```http
PUT /api/venues/:id/schedule
Authorization: Bearer <token>
Content-Type: application/json

{
  "schedule": {
    "monday": {"start": "09:00", "end": "18:00"},
    "tuesday": {"start": "09:00", "end": "18:00"},
    ...
  }
}
```

### Проверить доступность площадки
```http
GET /api/venues/:id/availability?start_time=2026-01-25T10:00:00Z&end_time=2026-01-25T12:00:00Z
```

**Query параметры:**
- `start_time` - начало периода (ISO 8601)
- `end_time` - конец периода (ISO 8601)

### Получить бронирования площадки
```http
GET /api/venues/:id/bookings
Authorization: Bearer <token>
```

---

## 4. Типы площадок (Venue Types)

### Получить список типов площадок
```http
GET /api/venue-types
```

**Ответ:**
```json
[
  "Конференц-зал",
  "Коворкинг",
  "Спортивный зал",
  ...
]
```

---

## 5. Бронирования (Bookings)

### Создать бронирование
```http
POST /api/booking
Authorization: Bearer <token>
Content-Type: application/json

{
  "venue_id": 1,
  "start_time": "2026-01-25T10:00:00Z",
  "end_time": "2026-01-25T12:00:00Z",
  "notes": "Корпоративное мероприятие"
}
```

### Получить бронирование по ID
```http
GET /api/bookings/:id
```

### Получить все бронирования текущего пользователя
```http
GET /api/bookings
Authorization: Bearer <token>
```

### Обновить бронирование
```http
PUT /api/bookings/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "start_time": "2026-01-25T11:00:00Z",
  "end_time": "2026-01-25T13:00:00Z",
  "notes": "Обновленные заметки"
}
```

### Отменить бронирование
```http
POST /api/bookings/:id/cancel
Authorization: Bearer <token>
Content-Type: application/json

{
  "reason": "Изменение планов"
}
```

### Получить сводку бронирования (агрегированные данные)
```http
GET /api/bookings/:id/summary
Authorization: Bearer <token>
```

**Ответ включает:**
- Информацию о бронировании
- Информацию о площадке
- Информацию о платеже

```json
{
  "booking": {...},
  "venue": {...},
  "payment": {...},
  "venue_error": null,
  "payment_error": null
}
```

---

## 6. Платежи (Payments)

### Создать платеж
```http
POST /api/payments
Authorization: Bearer <token>
Content-Type: application/json

{
  "booking_id": "123",
  "amount": 10000,
  "payment_method": "card"
}
```

### Получить историю платежей
```http
GET /api/payments
Authorization: Bearer <token>
```

### Получить платеж по ID
```http
GET /api/payments/:id
Authorization: Bearer <token>
```

### Получить платеж по ID бронирования
```http
GET /api/bookings/:id/payment
Authorization: Bearer <token>
```

### Создать возврат
```http
POST /api/payments/:id/refund
Authorization: Bearer <token>
Content-Type: application/json

{
  "reason": "Отмена бронирования",
  "amount": 10000
}
```

---

## Примеры использования с curl

### Регистрация и получение токена
```bash
curl -X POST http://localhost:8085/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "full_name": "John Doe"
  }'
```

### Получение списка площадок
```bash
curl http://localhost:8085/api/venues?page=1&limit=10
```

### Создание бронирования
```bash
curl -X POST http://localhost:8085/api/booking \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{
    "venue_id": 1,
    "start_time": "2026-01-25T10:00:00Z",
    "end_time": "2026-01-25T12:00:00Z"
  }'
```

### Получение сводки бронирования
```bash
curl http://localhost:8085/api/bookings/123/summary \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## Коды ответов

- `200 OK` - Успешный запрос
- `201 Created` - Ресурс создан
- `400 Bad Request` - Неверный запрос
- `401 Unauthorized` - Требуется авторизация
- `403 Forbidden` - Доступ запрещен
- `404 Not Found` - Ресурс не найден
- `500 Internal Server Error` - Внутренняя ошибка сервера
- `502 Bad Gateway` - Сервис недоступен

---

## Примечания

1. Все даты и время должны быть в формате ISO 8601 (например: `2026-01-25T10:00:00Z`)
2. JWT токен получается при регистрации или входе
3. Токен должен передаваться в заголовке `Authorization: Bearer <token>`
4. Gateway автоматически перенаправляет запросы к соответствующим микросервисам
5. Некоторые endpoints могут требовать дополнительных прав (например, владелец площадки)
