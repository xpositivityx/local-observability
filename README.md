# Using Traces in Tests

## Prereqs
- Docker
- Docker Compose
- Vscode/Cursor
- Devcontainers extension

## Running the example
- Open vscode or cursor.
- In command palette "Dev Containers: Rebuild and Reopen in Container"
- run 
make test
- Open browser and navigate to [Jaeger](localhost:16686/search)

Run the tests multiple times. It is set up to fail about half the time so you can see what errors look like.