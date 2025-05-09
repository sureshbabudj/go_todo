# Go Todo Application

This is a simple **Todo List API** built with Go. The application allows users to manage their tasks through a RESTful API. It supports operations such as creating, retrieving, updating, and deleting tasks. The tasks are stored in a JSON file (`data/todos.json`).

## Features

- **Create a Task**: Add a new task to the todo list.
- **Retrieve All Tasks**: Get a list of all tasks.
- **Retrieve a Task by ID**: Fetch details of a specific task using its ID.
- **Update a Task**: Modify the details of an existing task.
- **Delete a Task**: Remove a task from the list.

## API Endpoints

### `/api/todos` (HandleTodos)

- **GET**: Retrieve all tasks.
- **POST**: Create a new task.

### `/api/todos/{id}` (HandleTodoById)

- **GET**: Retrieve a task by its ID.
- **PATCH**: Update a task by its ID.
- **DELETE**: Delete a task by its ID.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/go_todo.git
   cd go_todo
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Generate frontend assets (CSS and JS):

   ```bash
   npm install
   npm run build
   ```

4. Run the application:

   ```bash
   go run cmd/web/main.go
   ```

5. The server will start on `http://localhost:8080`.

## Configuration

- **Data Storage**: Tasks are stored in `data/todos.json`. If the file or directory does not exist, it will be created automatically.
- **Port**: The application runs on port `8080` by default.

## Frontend

The frontend is built using **Tailwind CSS** and **TypeScript**. The assets are located in the `styles/` and `scripts/` directories and are compiled into the `public/` directory.

### Commands to Build Frontend Assets

- **Build CSS**:

  ```bash
  npm run build:css
  ```

  This command compiles the Tailwind CSS file (`styles/main.css`) into `public/styles.css`.

- **Build JavaScript**:

  ```bash
  npm run build:js
  ```

  This command bundles the TypeScript files (`scripts/main.ts`) into `public/scripts.js`.

- **Watch for Changes**:
  ```bash
  npm run watch
  ```
  This command watches for changes in CSS and TypeScript files and automatically rebuilds the assets.

### Frontend Assets

- **CSS**: The compiled CSS file is located at `public/styles.css`.
- **JavaScript**: The bundled JavaScript file is located at `public/scripts.js`.

Ensure these files are included in your HTML templates:

```html
<link href="/assets/styles.css" rel="stylesheet" />
<script src="/assets/scripts.js"></script>
```

## Example Usage

### Create a Task

```bash
curl -X POST -H "Content-Type: application/json" -d '{"description": "Buy groceries"}' http://localhost:8080/api/todos
```

### Get All Tasks

```bash
curl http://localhost:8080/api/todos
```

### Get a Task by ID

```bash
curl http://localhost:8080/api/todos/1
```

### Update a Task

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"description": "Buy groceries and milk"}' http://localhost:8080/api/todos/1
```

### Delete a Task

```bash
curl -X DELETE http://localhost:8080/api/todos/1
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
