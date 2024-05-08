# Employee CRUD API

This is a simple CRUD (Create, Read, Update, Delete) API for managing employees. It's built using Go and Gin framework.

## Getting Started

To get started with this project, follow these steps:

1. Clone this repository:

```
git clone https://github.com/akshay-singla/employee-crud-api.git
```

2. Install dependencies:

```
go mod tidy

docker-compose up -d 
```

3. Set up the database:
 - Update the database configuration in `config/config.go` file.
 - Ensure PostgreSQL is installed and running.


4. Run the application:

```
go run main.go
```

The server should now be running at `http://localhost:8080`.

## API Endpoints

- `POST /employee`: Create a new employee.
- `GET /employee`: Get a list of all employees.
- `GET /employee/:id`: Get details of a specific employee by ID.
- `PUT /employee/:id`: Update details of a specific employee.
- `DELETE /employee/:id`: Delete a specific employee by ID.

## Dependencies

- Gin: Web framework for Go.
- pq: PostgreSQL driver for Go.


## Docker
You can build a Docker image for this service using the provided Dockerfile:

    docker build -t employee-crud-api .

To run the Docker container:

    docker run -p 8080:8080 employee-crud-api


## Testing with Postman
A Postman collection is provided for testing all the APIs. You can easily import this collection into Postman by following these steps:

1. Open Postman.
2. Click on the "Import" button located at the top left corner of the window.
3. Select the option to import from file.
4. Choose the postman_collection.json file provided.

Once imported, all the requests will be available in the Postman collection. You can use these requests to test each API endpoint individually.




## Configuration

The application reads configuration from environment variables. You can set the following environment variables:

- `DB_HOST`: PostgreSQL database host.
- `DB_PORT`: PostgreSQL database port.
- `DB_USER`: PostgreSQL database username.
- `DB_PASSWORD`: PostgreSQL database password.
- `DB_NAME`: PostgreSQL database name.

Alternatively, you can update the configuration in `config/config.go` file directly.

## Contributing

Contributions are welcome! Please feel free to open issues or pull requests.
