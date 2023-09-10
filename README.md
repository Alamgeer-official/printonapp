# PrintOnApp - GoLang Web Application with Gin, JWT, and GORM

---

## Overview

PrintOnApp is a web application built using the Go programming language, the Gin web framework, JWT middleware for authentication and authorization, and GORM for PostgreSQL database integration. The application is designed to demonstrate my skills in building a secure and efficient web application using Go.

## Features

- User authentication and authorization using JWT (JSON Web Tokens).
- PostgreSQL database integration for data storage.
- Structured routing with Gin's group functionality.
- Three main routes: admin, user, and home.
- API endpoints for user login and signup.
- API endpoint for adding products (admin route).
- API endpoint for retrieving user information.
- Sample home page endpoints for testing purposes.

## Installation

Install the required dependencies:

shell
Copy code
go mod tidy
Set up your PostgreSQL database and configure the connection in the config/config.go file.

Run the application:

shell
Copy code
go run main.go
Usage
User Authentication and Product Management
User Login:

To log in, send a POST request to /login with your credentials.

##Example:

shell
Copy code
curl -X POST -d '{"username": "your_username", "password": "your_password"}' http://localhost:8080/login
User Registration:

To register as a new user, send a POST request to /signup with your registration details.

##Example:

shell
Copy code
curl -X POST -d '{"username": "new_user", "password": "new_password"}' http://localhost:8080/signup
Add a Product (Admin Route):

To add a new product as an admin, send a POST request to /admin/product.

##Example:

shell
Copy code
curl -X POST -d '{"name": "Product Name", "description": "Product Description", "price": 10.99}' http://localhost:8080/admin/product
User Information
Retrieve User Information:

To retrieve user information, send a GET request to /users.

##Example:

shell
Copy code
curl http://localhost:8080/users/
##API Endpoints
POST /login: User login.
POST /signup: User registration.
POST /admin/product: Add a product (admin route).
GET /users/: Retrieve user information.
GET /homepage/: Sample home page routes.
##Technologies Used
Go programming language.
Gin web framework.
JWT middleware for authentication and authorization.
GORM for PostgreSQL database integration.
vb
   
