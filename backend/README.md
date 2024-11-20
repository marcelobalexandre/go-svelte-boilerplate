# Go API Backend

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Live reload the application:
```bash
make dev 
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```

Generate a migration:
```bash
make db-generate-migration <migration_name>
```

Run migrations (and create the database if it doesn't exist):
```bash
make db-up
```

Rollback the last migration:
```bash
make db-down
```
