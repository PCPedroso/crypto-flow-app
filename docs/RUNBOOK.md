# RUNBOOK — Operações e resposta a incidentes

## Contatos on-call
- Time: equipe-ops - ops@empresa.local
- Escalada: SRE lead -> Eng Prod -> CTO

## Health checks
- Endpoint: GET /health
- Comando local:
  powershell -Command "Invoke-RestMethod -Uri http://localhost:8080/health"

## Incidentes comuns
1. Serviço indisponível
   - Verificar containers: docker ps
   - Logs: docker logs <container>
   - Reiniciar: docker restart <container>
2. Latência alta ou filas lentas
   - Verificar consumidores, reprocessar jobs
3. Falha na integração com carteiras
   - Verificar keys e endpoints externos

## Playbooks críticos
- Recuperação do banco de dados (ver BACKUP_AND_RECOVERY.md)
- Rollback de deploy (ver DEPLOYMENT.md)