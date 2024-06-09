# JWT Authentication Service with Gin Gonic

## Project Introduction

This project implements a JWT (JSON Web Token) authentication service using the Gin Gonic framework in Golang. It demonstrates the creation of APIs for user sign-up, login, and retrieving user details using MongoDB as the database. The project follows a clean architecture approach to ensure maintainability and scalability.

## Table of Contents

1. [Project Introduction](#project-introduction)
2. [Modules Description](#modules-description)
3. [Advantages and Main Functionality](#advantages-and-main-functionality)
4. [Tools, Technologies, and Platforms Used](#tools-technologies-and-platforms-used)
5. [Project Structure](#project-structure)
6. [Installation](#installation)
7. [Usage](#usage)
8. [Routes](#routes)
9. [Future Improvements](#future-improvements)

## Modules Description

- **Authentication Module**: Handles user registration and login. Provides JWT tokens upon successful login.
- **User Module**: Manages user data and provides endpoints to retrieve user information.
- **Middleware**: Implements authentication middleware to secure endpoints that require a valid JWT token.

## Advantages and Main Functionality

- **JWT Authentication**: Secure and stateless authentication using JWT.
- **MongoDB Integration**: Persistent storage for user data.
- **Clean Architecture**: Modular and maintainable codebase following best practices.
- **Scalable**: Easily extendable to add more features and handle more users.

## Tools, Technologies, and Platforms Used

- **Golang**: The primary programming language used for this project.
- **Gin Gonic**: A fast, minimalistic web framework for Golang.
- **MongoDB**: NoSQL database for storing user data.
- **JWT Go**: Library for handling JSON Web Tokens in Go.

## Project Structure

```
jwt-auth-service/
│
├── config/
│   └── config.go      # MongoDB connection and configuration
│
├── controllers/
│   ├── auth_controller.go    # Authentication handlers
│   └── user_controller.go    # User handlers
│
├── middleware/
│   └── auth_middleware.go    # JWT authentication middleware
│
├── models/
│   └── user.go        # User model
│
├── routes/
│   ├── auth_routes.go    # Authentication routes
│   └── user_routes.go    # User routes
│
├── services/
│   ├── auth_services.go    # Authentication business logic
│   └── user_services.go    # User business logic
│
├── main.go          # Entry point of the application
├── go.mod           # Go modules dependencies
└── README.md        # Project documentation
```

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Sushant-1999/jwt-auth.git
   cd jwt-auth-service
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Start MongoDB**:
   Ensure MongoDB is installed and running on your local machine. You can start it using:
   ```bash
   brew services start mongodb-community
   ```

4. **Run the application**:
   ```bash
   go run main.go
   ```

## Usage

### Routes

- **Sign Up**
  - **URL**: `/signup`
  - **Method**: `POST`
  - **Description**: Registers a new user.
  - **Body**:
    ```json
    {
      "username": "example",
      "email": "user@example.com",
      "password": "password"
    }
    ```

- **Login**
  - **URL**: `/login`
  - **Method**: `POST`
  - **Description**: Authenticates a user and returns a JWT token.
  - **Body**:
    ```json
    {
      "email": "user@example.com",
      "password": "password"
    }
    ```

- **Get All Users**
  - **URL**: `/users`
  - **Method**: `GET`
  - **Description**: Retrieves a list of all users. (Protected Route)

- **Get User by ID**
  - **URL**: `/users/:id`
  - **Method**: `GET`
  - **Description**: Retrieves user details by ID. (Protected Route)

## Future Improvements

- **Enhanced Error Handling**: Implement more detailed error handling and logging.
- **Unit Tests**: Add unit tests for all components to ensure robustness.
- **Role-Based Access Control**: Implement RBAC to handle different user roles and permissions.
- **Refresh Tokens**: Implement refresh tokens to improve security and usability.
- **API Documentation**: Integrate Swagger for better API documentation and testing.

## Contact

For any questions or suggestions, feel free to contact me at [sushantagawane99@gmail.com].