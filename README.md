# The repository contains 2 projects. The `awesomeProject` directory refers to the agent, and the `apiMetrics` project to the REST API.
# Agent Project

This agent project is designed to collect information about CPU, memory, and disk space usage from a device and send this data to a remote server.

## Requirements

- Go (Golang)
- Project dependencies: `github.com/shirou/gopsutil`, `gopkg.in/yaml.v2`
- The `apiMetrics` project must be running, as the agent consumes the API to send metric information
- A PostgreSQL database must be installed (a Docker version can be used)

## Installation

1. Clone this repository:

git clone https://github.com/tu-usuario/agent-project.git


2. Install dependencies:

go mod tidy


## Usage

The project uses a `registerdevices.yaml` file located at the root of the project. This file contains information about the devices the agent should monitor, including device IDs and the metrics to collect.

1. Run the agent:

go run main.go


The agent will collect information about CPU, memory, and disk usage according to the configuration provided in the YAML file and send this data to the specified server (in this case, a PostgreSQL database in the `monitor_details` table).

# API Metrics

API Metrics is a GoLang project designed to manage device metrics through a RESTful API.

## Introduction

API Metrics is a flexible and scalable solution for managing device metrics using a RESTful API. It provides a simple and efficient interface to register devices, store and retrieve metrics, and authenticate users.

## Features

- Device registration: Users can register new devices on the platform.
- Metrics management: Devices can send and retrieve historical metrics.
- User authentication: JSON Web Tokens (JWT) are used to authenticate and authorize users.

## Prerequisites

Before starting, make sure you have the following installed:

- Go 1.15 or higher
- PostgreSQL (a database named `metrics` must be created; tables are created automatically except `metrics`). Data is stored in 4 tables:
  a) `devices`  
  b) `metrics` (not used in this exercise, as it would be more useful when the backend is integrated with a frontend to display the name instead of just the ID — a script is provided)  
  c) `monitor_details`  
  d) `users`
- Git

## Installation

To set up and run the project locally, follow these steps:

1. Clone the repository:

git clone https://github.com/tuusuario/apiMetrics.git


2. Install dependencies:

go mod tidy


3. Configure environment variables by creating a `.env` file and setting the necessary configuration for the PostgreSQL database used.

## Usage

To run the project, use the following command:

go run main.go


## Endpoints

The API exposes the following endpoints:

a) `POST /devices/register`: Registers a new device (requires authentication — you must be logged in and provide the token).

curl -X POST -H "Authorization: TOKEN" -H "Content-Type: application/json" -d "{"name": "Device 7", "metric_id_1": 10, "metric_id_2": 20, "metric_id_3": 30, "ip": "192.168.1.100", "date_creation": "2024-03-05T10:00:00Z"}" http://localhost:8080/devices/register
{"device_id":6,"name":"Device 7","metric_id_1":10,"metric_id_2":20,"metric_id_3":30,"ip":"192.168.1.100","date_creation":"2024-03-06T00:52:12.8561028-06:00"}


b) `POST /devices/:id/metrics`: Saves metrics for a specific device.

curl -X POST -H "Content-Type: application/json" -d "{"metric_1": 15, "metric_1_value": 100, "metric_2": 25, "metric_2_value": 200, "metric_3": 35, "metric_3_value": 300}" http://localhost:8080/devices/1/metrics
{"device_id":1,"metric_1":15,"metric_1_value":100,"metric_2":25,"metric_2_value":200,"metric_3":35,"metric_3_value":300,"timestamp":"0001-01-01T00:00:00Z"}


c) `GET /devices/:id/metrics`: Retrieves the latest metrics for a specific device.

curl http://localhost:8080/devices/1/metrics


d) `GET /devices/:id/metrics/history`: Retrieves the historical metrics for a specific device.

curl http://localhost:8080/devices/1/metrics/history


e) `POST /login`: Authenticates a user and returns a JWT token.

curl -X POST -H "Content-Type: application/json" -d "{"email": "user@example.com", "password": "password123"}" http://localhost:8080/login


f) `POST /signup`: Registers a new user.

curl -X POST -H "Content-Type: application/json" -d "{"name": "angel q", "email": "test@yahoo.com", "password": "securepassword", "role": "admin"}" http://localhost:8080/signup


g) `GET /logout`: Logs out and invalidates the JWT token.

curl http://localhost:8080/logout
