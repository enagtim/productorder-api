# Order API

## Описание
Order API — это сервис для управления заказами. Поддерживает аутентификацию пользователей, создание и получение заказов.

## Стек технологий
- **Golang** (net/http, gorm)
- **PostgreSQL**
- **JWT (github.com/golang-jwt/jwt/v5)**
- **Docker** (для тестирования и локальной разработки)

## Установка и запуск

### 1. Клонирование репозитория
```sh
git clone https://github.com/enagtim/product-order-api.git
cd product-order-api
```

### 2. Настройка переменных окружения
Создайте файл `.env` и добавьте в него:
```env
DSN=host=localhost user=user password=user dbname=order-api port=5432 sslmode=disable
SECRET=/2+XnmJGz1j3ehIVI/5P9kl+CghrE3DcS7rnT+qar5w=
```

### 3. Запуск сервера
```sh
go run main.go
```

## API эндпоинты

### Авторизация пользователя
**POST** `/auth`
```json
{
  "phone": "83263346912",
}
```
Ответ:
```json
{
  "token": "your-jwt-token"
}
```

### Создание заказа
**POST** `/orders`
**Headers:**
```
Authorization: Bearer your-jwt-token
```
```json
{
  "productsIDs": [1, 2, 3]
}
```

### Получение заказа по ID
**GET** `/orders/{id}`
**Headers:**
```
Authorization: Bearer your-jwt-token
```
Ответ:
```json
{
  "id": 1,
  "userId": 10,
  "products": [
    {"id": 1, "name": "Product 1", "price": 100},
    {"id": 2, "name": "Product 2", "price": 150}
  ]
}
```

## Docker

### Запуск базы данных через Docker
```sh
docker-compose up -d
```

### Остановка контейнеров
```sh
docker-compose down
```