# ERP Acadêmico Backend

API em Go organizada em módulos para gestão acadêmica, recursos humanos e financeiro.

## Estrutura

- `models/`: modelos GORM compartilhados da aplicação, como aluno, curso, departamento e pagamento.
- `modules/academico`: endpoints de alunos e cursos.
- `modules/rh`: endpoints de funcionários.
- `modules/financeiro`: endpoints de pagamentos.
- `routes/`: registro das rotas e Swagger.
- `config/`: conexão com banco de dados e migrações.

## Execução

```bash
go run .
```

Por padrão o projeto usa SQLite em `erp_academico.db`. Para alterar, configure `DATABASE_DSN`.

## Swagger

A documentação Swagger fica disponível em:

```text
http://localhost:8080/swagger/index.html
```

## Funcionalidades iniciais

- Lançamento de notas pelo módulo acadêmico em `POST /api/v1/academico/notas`.
- Consulta de notas lançadas em `GET /api/v1/academico/notas`.
- Criação e listagem de encargos financeiros em `/api/v1/financeiro/encargos`.
- Validação de dívidas e encargos pendentes por aluno em `GET /api/v1/financeiro/alunos/{alunoID}/dividas/validar`.
