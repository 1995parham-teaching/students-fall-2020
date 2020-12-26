# students.go

## Introduction

An application for gathering the students list and returning them.
It is used as an example for web programming with Golang.

## API

The following API is used for creating new students.

```sh
curl -vvv -X POST 127.0.0.1:8080/student -H "Content-Type: application/json" -d '{"first_name": "Parham", "last_name": "Alvani"}'
```
