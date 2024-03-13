## TodoList - API (Go-Lang)

### Description
This is a simple API for a TodoList. It was developed using Go-Lang and the Golang Fiber framework.

### How To Run
1. Clone the repository
Code:
```bash
$ git clone https://github.com/Hanifanta/todolist-api.git
```
2. Create Database and Database Migrations
```bash
$ make createdb
$ make migrateup
```
3. Run the application
```bash
$ make runapp
```
4. Import api.spec.yaml to Postman or Insomnia to test the API
5. To stop the application, press `Ctrl + C` in the terminal

### Features
Activities:
- Create a new activity
- Get all activities
- Get an activity by ID
- Update an activity
- Delete an activity

Todos:
- Create a new todo
- Get all todos
- Get a todo by ID
- Update a todo
- Delete a todo

### Technologies
- Go-Lang
- Golang Fiber
- Gorm
- MySQL
- Docker
- Makefile