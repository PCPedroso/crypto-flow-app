# BACKUP AND RECOVERY

## Estratégia
- Backups diários do banco de dados
- Snapshot de volumes semanais

## RTO / RPO
- RTO alvo: 1 hora
- RPO alvo: 4 horas

## Procedimento de restauração (resumo)
1. Identificar backup válido.
2. Restaurar em ambiente isolado.
3. Rodar sanity checks e testes de integridade.
4. Promover para produção após validação.

## Testes
- Testar restore trimestralmente