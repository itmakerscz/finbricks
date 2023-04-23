# Finbricks Chat App

Voice chatbot

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)
- [Project setup](#project-setup)

## Installation

Brief instructions on how to install and set up the project.

docker build -t finbricks-chat-app:dev .

docker run -v ${PWD}:/app -v /app/node_modules -p 8081:8080 --rm finbricks-chat-app:dev

## Usage

Detailed instructions on how to use the project, including examples and screenshots if applicable.

## API Documentation

If your project includes an API, provide a link to the API documentation here.

## Contributing

Instructions on how to contribute to the project, including guidelines for pull requests and coding standards.

## License

This project is licensed under the terms of the MIT license. See the [LICENSE](LICENSE) file for details.

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
