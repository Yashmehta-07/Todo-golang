# Task Manager API

A simple RESTful API built with Go . It supports basic CRUD operations:

- **Create**: Add a new task
- **Read**: List all tasks
- **Update**: Update an existing task
- **Delete**: Delete a task

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
