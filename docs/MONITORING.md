# MONITORING e SLOs

## Métricas essenciais
- Disponibilidade (uptime)
- Latência de endpoints críticos
- Erros por minuto (5xx)
- Taxa de sucesso de integrações com wallets

## Ferramentas
- Prometheus (métricas), Grafana (dashboards), Alertmanager (alertas)
- Logs centralizados (ELK/Cloud)

## Alertas mínimos
- Serviço offline
- Erro 5xx > threshold
- Increase de latência > threshold

## SLOs e SLIs
- SLO: 99.9% disponibilidade para endpoints críticos
- Definir SLI correspondente e políticas de escalonamento