# El repositorio cuenta con 2 proyectos , el directorio awesomeProject hace referencia al agente y el proyecto apiMetrics a la apirest
# Agent Project

Este proyecto de agente está diseñado para recopilar información sobre el uso de CPU, memoria y espacio en disco de un dispositivo y enviar estos datos a un servidor remoto.

## Requisitos

- Go (Golang)
- Dependencias del proyecto: `github.com/shirou/gopsutil`, `gopkg.in/yaml.v2`
- Estar ejecutando el proyecto apiMetrics , ya que consume la api para recabar informacion de las metricas
- Tener instalada una base de datos postgres (Se puede ocupar una version Docker)

## Instalación

1. Clona este repositorio:

```
git clone https://github.com/tu-usuario/agent-project.git
```

2. Instala las dependencias:

```
go mod tidy
```

## Uso

El proyecto hace uso de un archivo registerdevices.yaml configurado en la raíz del proyecto. Este archivo contine información sobre los dispositivos que el agente debe monitorear, incluidos los ID de dispositivo y las métricas a recopilar.

1. Ejecuta el agente:

```
go run main.go
```

El agente recopilará información sobre el uso de CPU, memoria y espacio en disco según la configuración proporcionada en el archivo YAML y enviará estos datos al servidor especificado (En este caso un base de datos postgress en la tabla monitor_details)

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
- PostgreSQL (se debe crear una base de datos de nombre metrics, las tablas son creadas automaticamente exepto metrics), en ella se guardara informacion en 4 tablas:
  a) devices
  b) metrics(para efecto del ejercicio no se ocupara ya que serviria mas cuando el backend este integrado aun front para en lugar de mostrar el id muestre el nombre (se adjunta el script)
  c) monitor_details
  d) users
- Git

## Instalación

Para configurar y ejecutar el proyecto localmente, sigue estos pasos:

1. Clona el repositorio:

```
git clone https://github.com/tuusuario/apiMetrics.git
```

2. Instala las dependencias

```
go mod tidy
```

3. Configura las variables de entorno creando un archivo .env y estableciendo las configuraciones para la vase de datos que se utiliza(Postgres).

## Uso

Para ejecutar el proyecto, utiliza el siguiente comando:

```
go run main.go
```

## Endpoints

La API expone los siguientes endpoints:

a) POST /devices/register: Registra un nuevo dispositivo (requiere autenticación, haberse logueado on usuario y pasar el token).

```
curl -X POST -H "Authorization: TOKEN" -H "Content-Type: application/json" -d "{\"name\": \"Device 7\", \"metric_id_1\": 10, \"metric_id_2\": 20, \"metric_id_3\": 30, \"ip\": \"19
2.168.1.100\", \"date_creation\": \"2024-03-05T10:00:00Z\"}" http://localhost:8080/devices/register
{"device_id":6,"name":"Device 7","metric_id_1":10,"metric_id_2":20,"metric_id_3":30,"ip":"192.168.1.100","date_creation":"2024-03-06T00:52:12.8561028-06:00"}
```
    
b) POST /devices/:id/metrics: Guarda métricas para un dispositivo específico.

```
curl -X POST -H "Content-Type: application/json" -d "{\"metric_1\": 15, \"metric_1_value\": 100, \"metric_2\": 25, \"metric_2_value\": 200, \"metric_3\": 35, \"metric_3_value\": 300}" http://localhost:8080/devices/1/metrics
{"device_id":1,"metric_1":15,"metric_1_value":100,"metric_2":25,"metric_2_value":200,"metric_3":35,"metric_3_value":300,"timestamp":"0001-01-01T00:00:00Z"}
```
    
c) GET /devices/:id/metrics: Obtiene las últimas métricas para un dispositivo específico.

```
curl http://localhost:8080/devices/1/metrics
```

d) GET /devices/:id/metrics/history: Obtiene las métricas históricas para un dispositivo específico.

```
curl http://localhost:8080/devices/1/metrics/history
```

e) POST /login: Para autenticacion de usuario, devuelve un token JWT.

```
curl -X POST -H "Content-Type: application/json" -d "{\"email\": \"user@example.com\", \"password\": \"password123\"}" http://localhost:8080/login
```

f) POST /signup: Registra un nuevo usuario.

```
curl -X POST -H "Content-Type: application/json" -d "{\"name\": \"angel q\", \"email\": \"test@yahoo.com\", \"password\": \"securepassword\", \"role\": \"admin\"}" http://localhost:8080/signup
```

g) GET /logout: Cierra sesión e invalida el token JWT.

```
curl http://localhost:8080/logout
```
