# API Metrics

API Metrics es un proyecto GoLang diseñado para gestionar métricas de dispositivos a través de una API RESTful.

## Introducción

API Metrics es una solución flexible y escalable para gestionar métricas de dispositivos utilizando una API RESTful. Proporciona una interfaz simple y eficiente para registrar dispositivos, almacenar y recuperar métricas, y autenticar usuarios

## Características

- Registro de dispositivos: Los usuarios pueden registrar nuevos dispositivos en la plataforma.
- Gestión de métricas: Los dispositivos pueden enviar métricas y recuperar métricas históricas.
- Autenticación de usuarios: Se utiliza JSON Web Tokens (JWT) para autenticar y autorizar usuarios.

## Requisitos Previos

Antes de comenzar, asegúrate de tener instalados los siguientes requisitos:

- Go 1.15 o superior
- PostgreSQL
- Git

## Instalación

Para configurar y ejecutar el proyecto localmente, sigue estos pasos:

1. Clona el repositorio:

```bash
git clone https://github.com/tuusuario/apiMetrics.git

2. Instala las dependencias

```bash
go mod tidy

3. Configura las variables de entorno creando un archivo .env y estableciendo las configuraciones para la vase de datos que se utiliza(Postgres).

## Uso

Para ejecutar el proyecto, utiliza el siguiente comando:

```bash
go run main.go

## Endpoints

La API expone los siguientes endpoints:

    POST /devices/register: Registra un nuevo dispositivo (requiere autenticación).
    POST /devices/:id/metrics: Guarda métricas para un dispositivo específico.
    GET /devices/:id/metrics: Obtiene las últimas métricas para un dispositivo específico.
    GET /devices/:id/metrics/history: Obtiene las métricas históricas para un dispositivo específico.
    POST /login: Para autenticacion de usuario, devuelve un token JWT.
    POST /signup: Registra un nuevo usuario.
    GET /logout: Cierra sesión e invalida el token JWT.

