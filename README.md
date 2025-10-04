# Financial App Backend

API REST em Go para controle financeiro seguindo Standard Go Project Layout.

## Estrutura do Projeto

```
financial-app-backend/
├── cmd/
│   └── financial-app/         # Aplicação principal
│       └── main.go
├── internal/                  # Código privado da aplicação
│   ├── app/                   # Lógica da aplicação
│   └── pkg/                   # Pacotes internos
├── pkg/                       # Bibliotecas públicas reutilizáveis
├── api/                       # Definições de API (OpenAPI/Swagger specs)
├── web/                       # Assets web
│   ├── static/                # Arquivos estáticos (CSS, JS, imagens)
│   └── template/              # Templates
├── configs/                   # Arquivos de configuração
├── deployments/               # Configurações de deploy (Docker, K8s)
├── test/                      # Testes adicionais e dados de teste
├── docs/                      # Documentação do projeto
├── scripts/                   # Scripts de build, install, análise, etc.
├── build/                     # Packaging e CI
└── assets/                    # Outros assets (imagens, logos, etc.)
```

## Standard Go Project Layout

Esta estrutura segue o [Standard Go Project Layout](https://github.com/golang-standards/project-layout), que é amplamente adotado pela comunidade Go:

### Diretórios Principais

- **`/cmd`**: Aplicações principais do projeto. O nome do diretório deve corresponder ao nome do executável
- **`/internal`**: Código privado da aplicação. Este código não pode ser importado por outras aplicações
- **`/pkg`**: Código de biblioteca que pode ser usado por aplicações externas
- **`/api`**: Especificações OpenAPI/Swagger, arquivos de esquema JSON, definições de protocolo

### Diretórios de Aplicação Web

- **`/web`**: Componentes específicos de aplicações web (assets estáticos, templates, SPAs)

### Diretórios Comuns

- **`/configs`**: Templates de arquivos de configuração ou configurações padrão
- **`/deployments`**: Configurações e templates de deploy (docker-compose, kubernetes, etc.)
- **`/test`**: Testes adicionais e dados de teste
- **`/docs`**: Documentação de design e usuário
- **`/scripts`**: Scripts para build, install, análise, etc.
- **`/build`**: Packaging e Continuous Integration
- **`/assets`**: Outros assets do repositório (imagens, logos, etc.)

## Tecnologias

- Go 1.21+
- Gin (HTTP framework)
- GORM (ORM)
- PostgreSQL
- Docker

## Como executar

```bash
# Instalar dependências
go mod tidy

# Executar a aplicação
go run cmd/financial-app/main.go
```

## Vantagens desta Estrutura

- ✅ Padrão amplamente adotado pela comunidade Go
- ✅ Separação clara entre código público e privado
- ✅ Facilita a organização de projetos grandes
- ✅ Compatível com ferramentas de CI/CD
- ✅ Estrutura familiar para desenvolvedores Go