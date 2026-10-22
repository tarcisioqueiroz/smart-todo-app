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
