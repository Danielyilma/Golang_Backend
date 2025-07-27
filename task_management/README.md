# Task API Documentation

This API provides a simple task management service using the Gin web framework. It supports full CRUD operations for task resources.

---

## Base URL

```
http://localhost:8090
```

---

## Endpoints

### 1. Get All Tasks

**GET** `/tasks`

#### Description

Fetches all tasks from the in-memory database.

#### Response

```json
{
  "tasks": [
    {
      "id": "string",
      "title": "string",
      "description": "string",
      "status": "IN PROGRESS | COMPLETED",
      "created_at": "RFC3339 timestamp",
      "updated_at": "RFC3339 timestamp"
    }
  ]
}
```

---

### 2. Get a Single Task

**GET** `/tasks/:id`

#### Description

Fetches a task by its ID.

#### Parameters

* `id` (path): Task ID

#### Response (Success)

```json
{
  "task": {
    "id": "string",
    "title": "string",
    "description": "string",
    "status": "string",
    "created_at": "RFC3339 timestamp",
    "updated_at": "RFC3339 timestamp"
  }
}
```

#### Response (Error)

```json
{
  "error": "task with <id> not found"
}
```

---

### 3. Create a Task

**POST** `/tasks`

#### Description

Creates a new task.

#### Request Body

```json
{
  "title": "string",
  "description": "string"
}
```

#### Response

```json
{
  "success": {
    "id": "uuid",
    "title": "string",
    "description": "string",
    "status": "IN PROGRESS",
    "created_at": "RFC3339 timestamp",
    "updated_at": "RFC3339 timestamp"
  }
}
```

---

### 4. Update a Task

**PUT** `/tasks/:id`

#### Description

Updates an existing task by its ID.

#### Parameters

* `id` (path): Task ID

#### Request Body

```json
{
  "title": "string",
  "description": "string",
  "status": "string"
}
```

#### Response (Success)

```json
{
  "sucess": {
    "id": "string",
    "title": "string",
    "description": "string",
    "status": "string",
    "created_at": "RFC3339 timestamp",
    "updated_at": "RFC3339 timestamp"
  }
}
```

#### Response (Error)

```json
{
  "error": "task with <id> not found"
}
```

---

### 5. Delete a Task

**DELETE** `/tasks/:id`

#### Description

Deletes a task by its ID.

#### Parameters

* `id` (path): Task ID

#### Response (Success)

```json
{
  "sucess": "task with <id> deleted"
}
```

#### Response (Error)

```json
{
  "error": "task with <id> not found"
}
```

---

## Task Object Schema

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "status": "string",
  "created_at": "RFC3339 timestamp",
  "updated_at": "RFC3339 timestamp"
}
```

---

## Environment Variables

| Variable  | Description         | Default |
| --------- | ------------------- | ------- |
| APP\_PORT | Port to run the app | 8090    |

---

## Notes

* This project uses an in-memory map as a temporary database.
* On restart, all tasks will be reset.
* Status field is manually provided during updates.
