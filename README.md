# smart-todo-app

Um app de gerenciamento de tarefas multiplataforma integrado com inteligência artificial, possibilitando a criação e gerenciamento das atividades por comandos em linguagem natural.

## Stack

- Go;
- React Native;
- PostgreSQL;

## Como executar

### banco de dados
- Na raiz do projeto, crie o arquivo .env copiando o modelo do arquivo .env.example;
- Após isso, rode o comando `docker compose up -d`.

### backend
- Entre no diretorio backend e rode o comando `go run ./cmd/api`

### Frontend
- entre na pasta mobile,execute `npm install` para instalar as dependências;
- No mesmo diretório rode o comando `npx expo start` para iniciar o servidor metro do expo, que deixara o app disponível para ser executado via emuladores ou aplicativo mobile Expo GO;

### Migrations

Instale o goose (uma vez por computador):

```powershell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Dentro de `backend/`, com o banco rodando:

```powershell
$env:GOOSE_DRIVER = "postgres"
$env:GOOSE_DBSTRING = "postgres://USUARIO:SENHA@localhost:5432/BANCO?sslmode=disable"
$env:GOOSE_MIGRATION_DIR = "migrations"
goose up
```
