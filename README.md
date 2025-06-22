<div align="center">
<h1>Go RESTful API & CRUD</h1>

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
</div>

A simple CRUD REST API built with Go. A project recommended by Professor **Tenille Martins** during the Go Developer Training at DIO.

## 🚀 Development
First, I defined structures to build the CRUD model. The structures represent books and their respective authors. After that, I created a global variable that simulates a database. After this first step, I implemented a CRUD through handler functions, i.e., a function that creates a new book, a function that lists all books or a single book by its ID, a function that updates a book, and a function that removes a book. Finally, I implemented the main function, which creates an instance of a router and maps the HTTP methods with their respective endpoints while creating the CRUD.

### Features

- Complete CRUD that allows creating, reading, updating, and removing data.
- Built with standard library packages and Gorilla Mux, which provides an HTTP router.

## 🎯 Getting Started

### Prerequisites

- Go (version 1.2x or higher)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/andrecodea/go-rest-api-crud.git
   ```
2. Navigate to the project directory:
   ```sh
   cd go-rest-api-crud
   ```
3. Run the application:
   ```sh
   go run main.go
   ```

📚 Learnings
- Structs in Go
- Routing with Gorilla Mux
- RESTful APIs
- CRUDs

## 🤝 Contributions
Feel free to explore the code and suggest improvements!

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE.txt) file for details.

## 👨‍💻 Author
### André Codea
[![LinkedIn](https://img.shields.io/badge/LinkedIn-%230570a8?style=for-the-badge&logo=LinkedIn&logoColor=white)](https://www.linkedin.com/in/andrecodea/) [![GitHub](https://img.shields.io/badge/GitHub-%23121011?style=for-the-badge&logo=GitHub&logoColor=white)](https://github.com/andrecodea)
