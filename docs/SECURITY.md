# SECURITY — Política e práticas

## Gestão de segredos
- Não commitar .env; usar Vault / Azure Key Vault / AWS Secrets Manager
- Rotação periódica de chaves

## Vulnerabilidades
- Rodar scanner de dependências (ex.: Dependabot, Snyk)
- Fluxo de reporte: security@empresa.local

## Permissões
- Least privilege para contas de serviço
- Revisões periódicas de IAM

## Checklist pré-deploy
- Dependências com CVEs resolvidas
- Secrets válidos e acessíveis pelos runners