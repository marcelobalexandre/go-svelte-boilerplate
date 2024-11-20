# Go-Svelte Boilerplate

This repository provides a boilerplate setup for building web applications with a Go backend and a SvelteKit frontend. It was designed specifically for a **hackathon event** at **Codeminer42** to help quickly spin up projects.

⚠️ **Disclaimer:** This boilerplate is **not production-ready**. It is intended for rapid prototyping, hackathons, and learning purposes.

---

## Features

- **Go Backend**:
  A lightweight HTTP API built using the [Blueprint](https://go-blueprint.dev) project.

- **SvelteKit Frontend**:
  A frontend powered by [SvelteKit](https://kit.svelte.dev).

- **SQLite3 Database**:
  The boilerplate uses **SQLite3** for simple and lightweight data storage.
  - The database can be initialized, and migrations run with the following command:
    ```bash
    dbin/db-up
    ```
  - Uses [goose](https://github.com/pressly/goose) to handle database migrations.

- **Basic JWT Authentication**:
  Includes a simple implementation of **JWT-based authentication** to handle secure communication between the backend and frontend.
  - The backend generates a JWT token upon user login.
  - The frontend stores the token and uses it to authenticate API requests.

- **Hackathon-Friendly**:
  Simplified setup and minimal boilerplate code to allow teams to get started quickly.

- **Unified Development Workflow**:
  The entire app (frontend and backend) can be started using Docker Compose with a single command:
  ```bash
  dbin/dev
  ```

---

## Project Structure

```
/
├── backend/               # Go API backend
│   ├── cmd/api/main.go    # Entry point for the Go application
│   ├── internal/database  # Database files
│   ├── internal/server    # Server files
│
├── frontend/              # SvelteKit web application
│   ├── src/
│   │   ├── lib/           # Reusable frontend components and utilities
│   │   ├── routes/        # SvelteKit route definitions
│   │   └── app.html       # HTML template for SSR
│   ├── static/            # Static assets (images, etc.)
│   └── svelte.config.js   # SvelteKit configuration
│
├── dbin/                  # Docker Compose development scripts
│   ├── db-up              # Command to set up the SQLite database
│   └── dev                # Command to start the entire app using Docker Compose
│
└── docker-compose.yml     # Docker setup for development
```

---

## Requirements

- Docker Compose

or

- For running locally
  - Go (>= 1.23)
  - Node.js (>= 23.x)

---

## Getting Started

### Set Up the Database

Run the following command to create the SQLite database and apply migrations:
```bash
dbin/db-up
```

### Start the App

Run the following command to start the entire app (frontend and backend) using Docker Compose:
```bash
dbin/dev
```

- The backend API will be available at http://localhost:8080.
- The frontend will be available at http://localhost:3000.

### Executing Bash Commands in the Backend Docker Container

Run the bash command prefixed with `dbin/backend`, for example:
```bash
dbin/backend make db-generate-migration
```
Check [`backend/README.md`](backend/README.md) to find out more about the available commands.

---

Happy hacking! 🎉
