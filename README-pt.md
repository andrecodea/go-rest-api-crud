<div align="center">
<h1>API RESTful e CRUD em Go</h1>

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
</div>

Uma API REST CRUD simples construída com Go. Um projeto recomendado pela professora **Tenille Martins** durante a Formação Go Developer na DIO.

## 🚀 Desenvolvimento
Primeiramente, defini estruturas para construção do modelo do CRUD. As estruturas representam livros e seus respectivos autores, após isso, eu criei uma variável global que simula um banco de dados. Após esse primeiro passo, implementei um CRUD por meio de funções manipuladores (handlers), ou seja, uma função que cria um livro novo, uma função que lista todos os livros ou apenas um livro por meio de seu ID, uma função que atualiza um livro e uma função que remove um livro. Por último, implementei a função principal, que cria uma instância de um roteador e mapeiam os métodos HTTP com seus respectivos endpoints enquanto criava o CRUD.

### Funcionalidades

- CRUD completo que permite criar, ler, atualizar e remover dados.
- Construído com pacotes da biblioteca padrão e o Gorilla Mux, que fornece um roteador HTTP.

## 🎯 Começando

### Pré-requisitos

- Go (versão 1.2x ou superior)

### Instalação

1. Clone o repositório:
   ```sh
   git clone https://github.com/andrecodea/go-rest-api-crud.git
   ```
2. Navegue para o diretório do projeto:
   ```sh
   cd go-rest-api-crud
   ```
3. Execute a aplicação:
   ```sh
   go run main.go
   ```

📚 Aprendizados
- Estruturas em Go
- Roteamento com Gorilla Mux
- APIs RESTful
- CRUDs 

## 🤝 Contribuições
Sinta-se à vontade para explorar os códigos e sugerir melhorias!

## 📄 Licença
Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE.txt) para detalhes.

## 👨‍💻 Autor
### André Codea 
[![LinkedIn](https://img.shields.io/badge/LinkedIn-%230570a8?style=for-the-badge&logo=LinkedIn&logoColor=white)](https://www.linkedin.com/in/andrecodea/) [![GitHub](https://img.shields.io/badge/GitHub-%23121011?style=for-the-badge&logo=GitHub&logoColor=white)](https://github.com/andrecodea)