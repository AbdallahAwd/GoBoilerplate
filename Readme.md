# Plate - Go Project Structure Generator

![Go Boilerplate](https://i.postimg.cc/ZR4TjKQG/Go-Boilerplate.jpg)

Plate is a command-line tool that helps you quickly generate ideal project structures for various types of Go applications. Whether you're building a CLI tool, an API, or a general-purpose project, Plate can set up the initial structure for you, saving time and ensuring consistency across your projects.

<!-- ## Installation

To install Plate, make sure you have Go installed on your system, then run: -->

<!-- ```bash
go install github.com/AbdallahAwd/plate@latest
``` -->

## Usage

Plate uses a simple command-line interface to generate project structures. The basic syntax is:

```
plate -template=<template_type> -name=<project_name>

```

## Available Templates

1. CLI Project

```
plate -template=cli -name=MyCliProject
```

2. API Project

```
plate -template=api -name=MyApiProject
```

3. General Project (includes structures for both Web and API)

```
plate -template=all -name=MyProject
```

## Project Structures

1. CLI

```
MyCliProject/
├── cmd/
│   └── root.go
├── internal/
│   ├── commands/
│   └── utils/
├── pkg/
│   └── shared/
├── main.go
├── go.mod
└── README.md
```

2. API Project Structure

```
MyApiProject/
├── cmd/
│   └── myapi/
│       └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── handlers/
│   ├── models/
│   ├── repositories/
│   ├── services/
│   └── router/
├── pkg/
│   ├── logger/
│   └── middleware/
├── scripts/
├── .env
├── Dockerfile
├── go.mod
└── README.md
```

3. General Project Structure

```
MyProject/
├── cmd/
├── internal/
├── pkg/
├── api/
├── web/
├── scripts/
├── configs/
├── tests/
├── docs/
├── go.mod
└── README.md
```

### Customization

After generating your project structure, you can customize the files and directories to fit your specific needs. The generated structure provides a starting point based on best practices and common Go project layouts.

### Contributing

Contributions to Plate are welcome! Please feel free to submit a Pull Request.

### License

This project is licensed under the MIT License - see the LICENSE file for details.
