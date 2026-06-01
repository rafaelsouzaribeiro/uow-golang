# UOW + Clean Architecture (Go)

Projeto em Go com foco em **transações com Unit of Work (UoW)** usando **MySQL**.

## Tecnologias

- **Go**
- **MySQL** (persistência e transações)
- **Unit of Work** (controle transacional)
- **Chi** (web server / roteamento HTTP)
- **golang-migrate** (migrations de banco)
- **sqlc** (geração de código Go a partir de SQL)
- **Makefile** (atalhos para migrations)

## Migrations

As migrations ficam em `sql/migrations`.

### Subir tabelas (aplicar migrations)

```bash
make migrate
```

### Apagar/voltar migrations

```bash
make migratedown
```

### Criar novo arquivo de migration

```bash
make createmigration
```

## SQLC (geração de código Go)

O projeto usa `sqlc` para gerar código Go a partir dos arquivos SQL.

### Gerar código Go

```bash
sqlc generate
```

### Validar configuração e queries

> Certifique-se de ter o arquivo de configuração (`sqlc.yaml`) na raiz do projeto.

## Variáveis de ambiente (.env)

A aplicação usa um arquivo `.env` em `cmd/.env` com as variáveis:

```env
DB_USER=####
DB_PASSWORD=####
DB_HOST=####
DB_PORT=####
DB_NAME=####
WEB_PORT=####
```

> Caminho do arquivo: `cmd/.env`


## Configuração de banco (Makefile)

O projeto usa estas variáveis no `Makefile`:

- `DB_HOST`
- `DB_PORT`
- `DB_USER`
- `DB_PASS`
- `DB_NAME`

Ajuste esses valores antes de rodar os comandos de migration.

## Executar programa

Vai até a pasta `cmd/` e rode `go run main.go`