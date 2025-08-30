# DEPLOYMENT — Crypto Flow App

Resumo rápido dos ambientes (dev/staging/prod) e passos de deploy.

## Principais pontos
- Pipeline CI/CD: descrição breve (ex.: GitHub Actions -> Build -> Tests -> Deploy)
- Tags/versões: usar semver e criar tag no release

## Passos de deploy (exemplo)
1. Atualizar changelog e versão.
2. Buildar artefato localmente.
3. Subir tag e abrir release.
4. Pipeline executa deploy automático para staging.
5. Testes de smoke em staging.
6. Promover para produção (manual/auto conforme política).

## Rollback
- Reverter para tag anterior via CI/CD.
- Verificar integridade pós-rollback (smoke tests, monitoramento).

## Variáveis de ambiente por ambiente
- Referenciar ENVIRONMENT_VARIABLES.md

## Contatos
- Time de deploy: ops@empresa.local