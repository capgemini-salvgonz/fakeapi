# API Rest Stub Project

This is a template for creating an API Rest stub project using Go.

## Index

Welcome to the API Rest Stub Project!

Here are some useful links to get started:

- [Installation](#installation): Instructions on how to install the project.
- [Usage](#usage): Instructions on how to use the project.
- [API Documentation](#apiDocumentation): Description of mock API.
- [Contributing](#contributing): Guidelines for contributing to the project.
- [License](#license): Information about the project's license.

Feel free to explore the rest of the documentation to learn more about the API Rest Stub Project.


## Installation

To install this project, follow these steps:

1. Clone the repository.
2. Run `go mod download` to download the project dependencies.

## Usage

To use this project, follow these steps:

1. Implement your API endpoints in the `main.go` file.
2. Run `go run main.go` to start the server.

## API Documentation

### Endpoint

> GET /products

### Request example

> curl -X GET http://localhost:3000/products

### Response example

```json
{
  "products": [
    {
      "id": 1,
      "name": "Laptop Gamer",
      "price": 1500.00
    },
    {
      "id": 2,
      "name": "Teclado Mecánico",
      "price": 100.00
    },
    {
      "id": 3,
      "name": "Mouse Inalámbrico",
      "price": 50.00
    }
  ]
}
```

## Contributing

Contributions are welcome! Please follow the guidelines in the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## License

This project is licensed under the [MIT License](LICENSE).
