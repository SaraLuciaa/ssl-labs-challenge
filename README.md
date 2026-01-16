# Reto SSL Labs

En el este repositorio encontrará la solucion propuesta para el reto.

## Configuración del Proyecto

### Variables de entorno

Crea un archivo `.env` en la raíz del proyecto con las siguientes variables:

```env
DB_URL=postgresql://usuario:contraseña@localhost:5432/nombre_bd?sslmode=disable
PORT=8080
ALLOWED_ORIGIN=http://localhost:5173
```

- **DB_URL**: Cadena de conexión a PostgreSQL
- **PORT**: Puerto en el que correrá el servidor (opcional, por defecto 8080)
- **ALLOWED_ORIGIN**: URL del frontend permitido para hacer peticiones CORS (ej: http://localhost:5173)

### Cómo correr el Proyecto

1. Instalar las dependencias
```bash
go mod download
```
2. Ejecuta el proyecto
```bash
go run main.go
```

El servidor iniciará en `http://localhost:PORT`.

## Endpoints

### POST /analysis
Inicia el análisis SSL/TLS de un host y retorna el ID del análisis creado. 

Dado que SSL Labs procesa las solicitudes de forma asíncrona, el sistema implementa un mecanismo de polling en segundo plano que actualiza automáticamente el estado del análisis hasta su finalización.

**Request:**
```json
{
  "host": "www.example.com"
}
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000"
}
```

### GET /analysis
Obtiene todos los análisis registrados, ordenados del más reciente al más antiguo.

**Response:**
```json
{
  "analyses": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "host": "www.example.com",
      "port": 443,
      "protocol": "http",
      "is_public": false,
      "status": "READY",
      "start_time": "2026-01-16T10:00:00Z",
      "test_time": "2026-01-16T10:05:00Z",
      "engine_version": "2.2.0",
      "criteria_version": "2009q",
      "last_checked_at": "2026-01-16T10:05:30Z",
      "endpoints": [
        {
          ...
        }
      ],
      "created_at": "2026-01-16T10:00:00Z",
      "updated_at": "2026-01-16T10:05:30Z"
    }
  ]
}
```

### GET /analysis/:id
Obtiene los detalles de un análisis específico por su ID.

**Response:**
```json
{
  "analysis": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "host": "www.example.com",
    "port": 443,
    "protocol": "http",
    "is_public": false,
    "status": "READY",
    "start_time": "2026-01-16T10:00:00Z",
    "test_time": "2026-01-16T10:05:00Z",
    "engine_version": "2.2.0",
    "criteria_version": "2009q",
    "last_checked_at": "2026-01-16T10:05:30Z",
    "endpoints": [
      {
        "id": "660e8400-e29b-41d4-a716-446655440001",
        "ip_address": "192.0.2.1",
        "status_message": "Ready",
        "grade": "A+",
        "has_warnings": false,
        "is_exceptional": true
      }
    ],
    "created_at": "2026-01-16T10:00:00Z",
    "updated_at": "2026-01-16T10:05:30Z"
  }
}
```

