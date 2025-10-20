# Gin Quickstart

Uma aplicação Go utilizando Gin Framework, GORM e PostgreSQL.

## 🚀 Tecnologias

- **Go 1.24** - Linguagem de programação
- **Gin** - Framework web
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **Docker** - Containerização

## 🏃‍♂️ Executando o projeto

### Pré-requisitos

- Docker e Docker Compose
- Go 1.24+ (para desenvolvimento local)

### Com Docker (Recomendado)

1. **Clone o repositório**

   ```bash
   git clone <seu-repo>
   cd gin-quickstart
   ```

2. **Configure as variáveis de ambiente**

   ```bash
   cp .env.example .env
   ```

3. **Suba os containers**

   ```bash
   make docker-up
   # ou
   docker-compose up -d
   ```

4. **Acesse a aplicação**
   - API: http://localhost:8000
   - Health Check: http://localhost:8000/helthz

### Modo Desenvolvedor (com pgAdmin)

```bash
make docker-up-dev
# ou
docker-compose --profile dev up -d
```

- pgAdmin: http://localhost:5050
  - Email: admin@admin.com
  - Senha: admin123

### Desenvolvimento Local

1. **Suba apenas o PostgreSQL**

   ```bash
   docker-compose up -d postgres
   ```

2. **Execute a aplicação**
   ```bash
   make run
   # ou
   go run ./cmd/api
   ```

## 📊 Comandos Úteis

```bash
# Ver todos os comandos disponíveis
make help

# Logs dos containers
make docker-logs

# Logs apenas da API
make docker-logs-api

# Parar containers
make docker-down

# Rebuild containers
make docker-rebuild

# Executar testes
make test

# Build local
make build
```

## 🗄️ Banco de Dados

### Configurações Padrão

- **Host**: localhost (ou `postgres` no Docker)
- **Porta**: 5432
- **Usuário**: postgres
- **Senha**: postgres123
- **Database**: gin_quickstart

### Migrações com GORM

As migrações são executadas automaticamente usando GORM AutoMigrate:

```go
// Exemplo em main.go ou config
db.AutoMigrate(&entities.Category{})
```

## 🏗️ Estrutura do Projeto

```
.
├── cmd/
│   └── api/              # Ponto de entrada da aplicação
├── internal/
│   ├── config/           # Configurações (banco, etc.)
│   ├── entities/         # Modelos/Entidades
│   ├── repositories/     # Camada de dados
│   └── use-cases/        # Regras de negócio
├── scripts/              # Scripts SQL
├── tmp/                  # Arquivos temporários
├── docker-compose.yml    # Configuração Docker
├── Dockerfile            # Imagem da aplicação
└── Makefile             # Comandos automatizados
```

## 🔧 Configuração do GORM

Exemplo de uso do GORM com PostgreSQL:

```go
// internal/config/database.go
config := config.NewConfig()
err := config.ConnectDB()
if err != nil {
    log.Fatal("Failed to connect to database:", err)
}

// Auto-migrate entidades
err = config.DB.AutoMigrate(&entities.Category{})
if err != nil {
    log.Fatal("Failed to migrate database:", err)
}
```

## 🐳 Docker

### Multi-stage Build

O Dockerfile utiliza multi-stage build para otimizar o tamanho da imagem final:

- **Stage 1**: Build da aplicação com Go completo
- **Stage 2**: Runtime com Alpine Linux (imagem final ~20MB)

### Características de Segurança

- Usuário não-root
- Imagem Alpine (menor superfície de ataque)
- Health checks configurados
- Certificados CA incluídos

## 🔍 Health Checks

- **API**: `GET /helthz`
- **PostgreSQL**: `pg_isready`
- **Docker**: Health checks automáticos configurados

## 📝 Variáveis de Ambiente

Consulte `.env.example` para todas as variáveis disponíveis.

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT.
