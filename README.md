# Task Manager API

A simple RESTful API built with Go and the `go-chi` routing library for managing tasks. It supports basic CRUD operations:

- **Create**: Add a new task
- **Read**: List all tasks
- **Update**: Update an existing task
- **Delete**: Delete a task

## Features

- Add, list, update, and delete tasks.
- Lightweight Go server using the `chi` router for handling HTTP requests.
- JSON-based API responses.

## Endpoints

### `GET /tasks`

- **Description**: Retrieves the list of all tasks.
- **Response**:
  - Status 200 OK: Returns an array of tasks.
  - Status 404 Not Found: Returns a message if no tasks are found.

### `POST /tasks`

- **Description**: Adds a new task to the list.
- **Request Body**:
  ```json
  {
    "Desc": "Task description"
  }
