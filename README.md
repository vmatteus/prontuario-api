Prontuário API
# Prontuário API

Prontuário API is a backend service designed to manage medical records, following the hexagonal architecture pattern. This API provides a scalable and maintainable solution for handling patient data, appointments, and other healthcare-related information.

## Table of Contents
- [Getting Started](#getting-started)
- [Architecture](#architecture)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

This section will help you get a local copy of the project up and running for development and testing purposes.

## Architecture

The project follows Hexagonal Architecture (also known as Ports and Adapters), which promotes a decoupled structure between the core business logic and the external systems, such as databases, APIs, and UI.

## Requirements

Before you start, make sure you have the following installed:
- Golang (v1.20 or later)
- Docker (for database setup)
- SQLite
- PostgreSQL

## Installation

Clone the repository:
```bash
git clone https://github.com/vmatteus/prontuario-api.git
cd prontuario-api
```

Install dependencies:
```bash
go mod tidy
```

Set up the database:

The project uses PostgreSQL for the database. You can set it up locally or use Docker to run a PostgreSQL container.

Using Docker:
```bash
docker-compose up -d
```

Run the application:
```bash
go run main.go
```

## Usage

Once the application is running, you can access the API at `http://localhost:8080` (by default). Use tools like Postman or curl to interact with the endpoints.

## Environment Variables

The following environment variables are required to configure the application:

| Variable     | Description                  | Default Value |
|--------------|------------------------------|---------------|
| DB_HOST      | PostgreSQL host              | localhost     |
| DB_PORT      | PostgreSQL port              | 5432          |
| DB_USER      | PostgreSQL user              | postgres      |
| DB_PASSWORD  | PostgreSQL password          | password      |
| DB_NAME      | PostgreSQL database name     | prontuario    |
| PORT         | The port on which the API will run | 8080     |

To set these variables, create a `./configs/config.yml` with your custom values, like:
```bash
database:
  drive: "sqlite"
  sqlite:
    dbpath: "./sqlite.db"
  postgres:
    host: "localhost"
    port: 5432
    user: "postgres"
    password: "postgres"
    database: "postgres"
```

## API Documentation

For detailed information about the available endpoints, refer to the API documentation.

## Contributing

Contributions are welcome! If you'd like to contribute to this project, follow these steps:
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.