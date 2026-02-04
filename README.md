### Description
Backend for a notes app that stores notes in memory instead of using a database for simplicity, the purpose of this app is to get familiar with golang basics and syntax.

---

### How to run?
1. #### Clone the repository
```bash
git clone https://github.com/abdul-ghaffar01/notes-server.git
```

2. #### Navigate to the main.go file
```bash
cd notes-server/cmd/server
```

3. #### Run project with this command
```bash
go run main.go
```

---

### Features
- Get all notes on /notes route with GET request
- Create a new note on /create route with title and description passed in json body.<br/>
Example:
```json
{
    "title": "The very first note",
    "description": "Description of the first note"
}
```

## Project Structure

```text
.
├── cmd/
│   └── server/
│       └── main.go
|
├── internal/
│   ├── health/
│   │   ├── handler.go
│   │   └── handler_test.go
│   ├── note/
│   |   ├── handler.go
│   |   ├── handler_test.go
│   |   ├── service.go
│   |   ├── service_test.go
│   |   └── model.go
|   └── server/
│       ├── routes.go
│       └── server.go
|
├── tests/
|   
├── go.mod
└── README.md
```
