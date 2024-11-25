API Students 🎓
Este projeto foi desenvolvido durante o curso "Golang do Zero" e representa minha evolução como desenvolvedor. É uma API simples, mas funcional, para gerenciar estudantes. A ideia é aprender na prática conceitos importantes como CRUD, validação de dados e integração com banco de dados.

🚀 Funcionalidades
Cadastro de estudantes: Insira dados como nome, CPF, e-mail, idade e status ativo/inativo.
Consulta: Liste todos os estudantes ou filtre por status.
Edição e exclusão: Atualize informações ou remova estudantes.
🛠️ Tecnologias Utilizadas
Linguagem: Go
Banco de Dados: SQLite
Bibliotecas: gorilla/mux para roteamento, entre outras.
🗂️ Estrutura do Projeto
plaintext
Copiar código
api-students/
├── controllers/  # Regras de negócio
├── models/       # Estruturas de dados e integração com SQLite
├── routes/       # Rotas da API
├── main.go       # Entrada do projeto
⚡ Como Executar
Clone o repositório:
bash
Copiar código
git clone https://github.com/TrevisTJ/api-students.git
Entre na pasta do projeto:
bash
Copiar código
cd api-students
Instale as dependências:
bash
Copiar código
go mod tidy
Execute a aplicação:
bash
Copiar código
go run main.go
Acesse a API em http://localhost:8080.

📈 Próximos Passos
Adicionar autenticação simples.
Melhorar validações de entrada (CPF e e-mail).
Criar testes unitários para garantir qualidade.
Este projeto representa meu aprendizado e crescimento como desenvolvedor. Feedbacks são muito bem-vindos! 😊

