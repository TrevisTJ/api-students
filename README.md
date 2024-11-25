# API Students 🎓

Este projeto foi desenvolvido durante o curso *"Golang do Zero"* e representa minha evolução como desenvolvedor. É uma API simples, mas funcional, para gerenciar estudantes. A ideia é aprender na prática conceitos importantes como CRUD, validação de dados e integração com banco de dados.

## 🚀 Funcionalidades

- **Cadastro de estudantes:** Insira dados como nome, CPF, e-mail, idade e status ativo/inativo.  
- **Consulta:** Liste todos os estudantes ou filtre por status.  
- **Edição e exclusão:** Atualize informações ou remova estudantes.  

## 🛠️ Tecnologias Utilizadas

- **Linguagem:** Go  
- **Banco de Dados:** SQLite  
- **Bibliotecas:** `gorilla/mux` para roteamento, entre outras.
- **Postman:** Para testes de requisições HTTP.

## 🗂️ Estrutura do Projeto

```plaintext
api-students/
├── controllers/  # Regras de negócio
├── models/       # Estruturas de dados e integração com SQLite
├── routes/       # Rotas da API
├── main.go       # Entrada do projeto
````

## ⚡ Como Executar

1. Clone o repositório:
   ```bash
   git clone https://github.com/TrevisTJ/api-students.git
   ````
2. Entre na pasta do projeto:
  ```bash
    cd api-students
  ```
3. Instale as dependências:
  ```bash
    go mod tidy
  ```
4. Execute a aplicação
  ```bash
    go run main.go
  ```
Acesse a API no navegador ou em ferramentas como Postman, utilizando o endereço:
http://localhost:8080

Essa seção dá instruções claras para que qualquer pessoa consiga executar o seu projeto localmente. Se precisar adicionar ou ajustar algo, é só me avisar! 🚀


