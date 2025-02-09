# HOBE API REST BOILERPLATE

For small MS that works with one or two entities.

## Features

- Create, read, update, and delete resources
- JSON-based API
- Simple and easy to extend

## Requirements

- Go 1.23
- A running instance of a datastore documental database (Google Cloud Datastore)

## Installation

1. Clone the repository:
  ```sh
  git clone <URL>
  ```
2. Navigate to the project directory:
  ```sh
  cd hobe-api-boilerplate
  ```
3. Do not forget to rename before install!

4. Install dependencies:
  ```sh
  go mod tidy
  ```

## Configuration

Add datastore_sa.json file in service_accounts folder
Add loggin_sa.json file in service_accounts folder

Create a `.env` file in the root directory and add your database configuration.
Watch .env.example for requires vars and values.

## Running the Application

To start the server, run:
```sh
go run cmd/server/main.go
```

The server will start on `http://localhost:8082`.

## API Endpoints

- `GET /players` - List all players
- `GET /players/{id}` - Get a specific resource
- `POST /players` - Create a new resource
- `PUT /players/{id}` - Update a specific resource
- `DELETE /players/{id}` - Delete a specific resource

## Generate DOCS

Run
```sh
swag init -g cmd/server/main.go -o docs/swagger
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.