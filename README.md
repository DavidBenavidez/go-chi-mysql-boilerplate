# Golang Chi Mysql

# Requirements
GO 1.15

# Database
create a `sample` database in your local mysql
configure static/config.yaml to access your database

# Build
`$ go build ./cmd/api/`

# Run
`cd ./cmd/api`
`go run .`

# endpoints
POST /course
POST /student
DELETE /course

GET /course/xxx/student
POST /course/xxx/student

<!-- GET /course/xxx/student/xxx did not finish :( -->
