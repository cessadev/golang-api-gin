# go-api-gin

**Descripción:**  
Esta es una API RESTful construida en Golang utilizando el framework `Gin`. La API permite la autenticación y gestión de usuarios, conectándose a una base de datos MongoDB. La estructura sigue buenas prácticas de separación de responsabilidades y utiliza `docker-compose` para facilitar el despliegue de MongoDB.

## Características

- Autenticación de usuario con JWT.
- Gestión de usuarios (registro y obtención).
- Conexión a MongoDB para almacenamiento de datos.
- Arquitectura limpia con separación de handlers, servicios y repositorios.
- Uso de `Gin` como framework HTTP con middlewares de recuperación y registro.
- Configuración de base de datos en archivo dedicado.

## Tecnologías

- **Golang**
- **Gin** - Framework para el enrutamiento HTTP.
- **MongoDB** - Base de datos NoSQL.
- **Docker y Docker Compose** - Para ejecutar MongoDB en un contenedor.
- **JWT** - Para autenticación basada en tokens.

## Requisitos

- **Golang** 1.23 o superior
- **Docker** y **Docker Compose** (para MongoDB)

## Instalación y Configuración

1. **Clonar el repositorio**:

   ```bash
   git clone https://github.com/cessadev/golang-api-gin.git
   cd golang-api-gin
   ```

2. **Configurar Docker Compose para MongoDB**:

   Crea un archivo `docker-compose.yml` en la raíz del proyecto:

   ```yaml
   version: "3.8"
   services:
     mongodb:
       image: mongo
       container_name: mongodb
       ports:
         - "27017:27017"
       environment:
         MONGO_INITDB_ROOT_USERNAME: root
         MONGO_INITDB_ROOT_PASSWORD: example
   ```

3. **Ejecutar MongoDB**:

   Inicia el contenedor de MongoDB:

   ```bash
   docker-compose up --build
   ```

4. **Configurar la conexión a la base de datos**:

   En `config/database.go`, se utiliza esta cadena de conexión para MongoDB:

   ```go
   clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
   ```

5. **Ejecutar la API**:

   ```bash
   go run cmd/api/main.go
   ```

   Si la conexión es exitosa, deberías ver en la consola:

   ```
   Conectado a MongoDB
   [GIN-debug] Listening and serving HTTP on :8080
   ```

## Endpoints de la API

### Autenticación

- **Login**  
  `POST /auth/login`

  **Body** (JSON):
  ```json
  {
    "email": "example@example.com",
    "password": "yourpassword"
  }
  ```

  **Respuesta**:
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

### Usuarios

- **Crear usuario**  
  `POST /users`

  **Body** (JSON):
  ```json
  {
    "name": "Jorge Mahp",
    "email": "jmahp12345@example.com",
    "password": "yourpassword"
  }
  ```

  **Respuesta**:
  ```json
  {
    "id": "60c72b2f5f1b2c001c8e4c3e",
    "name": "Jorge Mahp",
    "email": "jmahp12345@example.com"
  }
  ```

- **Obtener usuario por ID**  
  `GET /users/:id`

  **Respuesta**:
  ```json
  {
    "id": "60c72b2f5f1b2c001c8e4c3e",
    "name": "Jorge Mahp",
    "email": "jmahp12345@example.com"
  }
  ```

## Seguridad y Modo de Producción

Para poner esta API en producción:

1. Cambia el modo de `debug` a `release`:

   ```go
   gin.SetMode(gin.ReleaseMode)
   ```

2. Configura proxies de confianza si tu API está detrás de un balanceador de carga o proxy.

3. **Variables de Entorno**: Reemplaza credenciales sensibles y configuraciones de MongoDB con variables de entorno.

## Licencia

Este proyecto está bajo la licencia MIT. Puedes hacer uso del código para tus propios proyectos.