## Description

This project is a **RESTful API for event management** built with **Go** and **Gin**.  
It provides endpoints for:
- **User registration and login** with hashed passwords.
- **JWT-based authentication** and protected routes.
- **CRUD operations for events** (create, list, retrieve, update, delete).
- **User registrations for events** (register / cancel registration).

Data is stored in a **SQLite** database, with tables for `users`, `events`, and `registrations`.

## How to execute

- **Prerequisites**
  - **Go** installed (compatible with the version declared in `go.mod`).
  - No external database server is required; the app uses a local **SQLite** file (`api.db`).

- **Install dependencies**

```bash
go mod tidy
```

- **Run the application**

```bash
go run main.go
```

The server will start on **`http://localhost:8080`**.

- **Basic endpoints**
  - `POST /signup` – create a new user.
  - `POST /login` – authenticate and receive a JWT token.
  - `GET /events` – list all events.
  - `GET /events/:id` – get a single event by ID.
  - **Authenticated (JWT required in `Authorization` header):**
    - `POST /events` – create an event.
    - `PUT /events/:id` – update an event.
    - `DELETE /events/:id` – delete an event.
    - `POST /events/:id/register` – register the authenticated user to an event.
    - `DELETE /events/:id/register` – cancel a registration.

## Libraries used

Main libraries (see `go.mod` for full list):

- **`github.com/gin-gonic/gin`**: HTTP web framework used to build the REST API.
- **`github.com/mattn/go-sqlite3`**: SQLite driver for Go, used for persistence.
- **`github.com/golang-jwt/jwt/v5`**: Library for generating and validating JWT tokens.
- **`golang.org/x/crypto/bcrypt`**: Password hashing and verification (used via `bcrypt`).

Additional transitive/indirect libraries are used internally by Gin, validation, JSON handling, and other tooling.

## Architecture files

High-level architecture is organized by **responsibility-based packages**:

- **Entry point**
  - `main.go`: initializes the database, creates the Gin engine, and registers routes.

- **Routing layer (`routes/`)**
  - `routes/routes.go`: central registration of all HTTP routes and middleware.
  - `routes/events.go`: handlers for event endpoints (`/events`, `/events/:id`, registration endpoints).
  - `routes/users.go`: handlers for authentication-related endpoints (`/signup`, `/login`).

- **Database layer (`database/`)**
  - `database/db.go`: initializes the global `DB` connection (SQLite) and creates the tables `users`, `events`, and `registrations`.

- **Domain models (`models/`)**
  - `models/event.go`: `Event` entity and its persistence methods (`Save`, `GetAllEvents`, `GetEventByID`, `Update`, `Delete`, `Register`, `CalcelRegistration`).
  - `models/user.go`: `User` entity and its methods (`Save`, `ValidateCredentials`).

- **Middleware (`middleware/`)**
  - `middleware/auth.go`: `Authenticate` middleware that validates JWTs from the `Authorization` header and injects `userId` into the request context.

- **Utilities (`utils/`)**
  - `utils/jwt.go`: JWT generation and verification (`GenerateToken`, `VerifyToken`).
  - `utils/hash.go`: password hashing and comparison (`HashPassword`, `CheckPasswordHash`).

## Entities snippets

- **Event**

```go
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}
```

- **User**

```go
type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
```

