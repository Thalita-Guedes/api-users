# API de Cadastro de Usuários (CRUD)

Uma API simples em **Go** para operações de Cadastros, Leitura, Atualização e Exclusão (**CRUD**) de usuários. Os dados são armazenados no **DynamoDB**. 

## Tecnologias Utilizadas

Este projeto utiliza as seguintes tecnologias:

* **GO**
* **AWS SDK for GO v2**
* **DynamoDB Local** (para desenvolvimento)
* **Docker e Docker Compose**
* **Postman** (ou similar para testes de endpoints)

---

# Como Rodar o Projeto

A aplicação está configurada com **Docker Compose** para facilitar o ambiente de desenvolvimento. Ao subir os serviços, serão criados dois contêineres: um para a **API** em Go e outro para o **DynamoDB Local**, que ja inicia com a tabela de usuários criada.

### Pré-requsitos

Certifique-se de ter o **Docker** e o **Docker-Compose** instalados em seu sistema.

### Passos

1.  **Clone o repositório:**
```bash
git clone [http://github.com/seu-usuário/api-users.git]
cd api-users
```

2. **Suba os contêineres:**
Este comando irá construir a imagem da API e iniciar os serviços.
```bash
docker compose up --build
```

### Acesso

* A **API** estara disponivel na porta `8080`.
* O contêiner do **DynamoDB Local** estará na porta `8000`.

A API está pronta para receber requsições nos seus endpoints.

---

## Estrutura do Projeto

A estrutura de diretórios do projeto é a seguinte:

api-users/
|-handlers/
|-models
|-repository
|-utils
|-DockerFile
|-Docker-Compose.yml
|-main.go
|README.md



