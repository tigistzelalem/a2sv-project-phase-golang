# Task Management API Documentation

## Endpoints

### GET /tasks
- Description: Get a list of all tasks.
- Response:
  - Status: 200 OK
  - Body: JSON array of tasks

### GET /tasks/:id
- Description: Get the details of a specific task.
- Parameters:
  - `id`: Task ID
- Response:
  - Status: 200 OK
  - Body: JSON object of the task
  - Status: 404 Not Found if the task does not exist

### POST /tasks
- Description: Create a new task.
- Request:
  - Body: JSON object with `title`, `description`, `due_date`, and `status`
- Response:
  - Status: 201 Created
  - Body: JSON object of the created task

### PUT /tasks/:id
- Description: Update a specific task.
- Parameters:
  - `id`: Task ID
- Request:
  - Body: JSON object with the new details of the task
- Response:
  - Status: 200 OK
  - Body: JSON object of the updated task
  - Status: 404 Not Found if the task does not exist

### DELETE /tasks/:id
- Description: Delete a specific task.
- Parameters:
  - `id`: Task ID
- Response:
  - Status: 204 No Content
  - Status: 404 Not Found if the task does not exist
