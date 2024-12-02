# Go REST API with JWT Authentication

This project is a REST API built with Go, providing the following functionality:

- **User Registration**: Allows users to register with a unique username and password.
- **User Login**: Authenticates users and generates a JWT token for authorization.
- **JWT Authentication**: Secures endpoints by validating JWT tokens.
- **Book Management**: Enables users to create and view books. Each user can only access their own books.

## Technologies Used
- **Go (Golang)**: The programming language used for building the API.
- **PostgreSQL**: The database for storing users and books.
- **JWT (JSON Web Tokens)**: For stateless authentication and securing endpoints.
- **Standard Library (`net/http`)**: For routing and handling HTTP requests.
