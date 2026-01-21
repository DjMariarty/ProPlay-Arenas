# Payment Service API Examples

## 1. Создать платеж (POST /payments, mock -> completed)

```bash
curl -X POST http://localhost:8084/payments \
  -H "Content-Type: application/json" \
  -d '{
    "booking_id": "550e8400-e29b-41d4-a716-446655440000",
    "user_id": "550e8400-e29b-41d4-a716-446655440001",
    "amount": 50000,
    "currency": "RUB",
    "method": "card"
  }'
```

---

## 2. Статус платежа (GET /payments/:id)

```bash
curl http://localhost:8084/payments/1
```

---

## 3. История платежей (GET /payments?user_id=...)

```bash
curl "http://localhost:8084/payments?user_id=550e8400-e29b-41d4-a716-446655440001&limit=10&offset=0"
```

---

## 4. Запросить возврат (POST /payments/:id/refund)

```bash
curl -X POST http://localhost:8084/payments/1/refund \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 25000,
    "reason": "Отмена брони"
  }'
```

---

## 5. Платеж по брони (GET /bookings/:id/payment)

```bash
curl "http://localhost:8084/bookings/550e8400-e29b-41d4-a716-446655440000/payment"
```
