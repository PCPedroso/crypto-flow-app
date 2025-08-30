# OPERATIONS / OPERAÇÕES

Filepath: docs/OPERATIONS.md
Purpose: Limites, monitoramento, retry/backoff e checklist operacional.
Owner: @dev
Status: DONE
Created: 2025-08-30
Updated: 2025-08-30

Resumo
- GPT-5 mini (Preview) — free: quotas dinâmicas → trate via retry/backoff e rate-limit.
- Recomendado: middleware para throttling, semaphore para concorrência e logs de headers (`Retry-After`, `X-RateLimit-*`).

Snippets (referência)
- Retry exponential, semaphore e rate limiter (ver exemplos no README e internal/middleware quando implementado).

Monitoring
- Expor métricas (Prometheus) e configurar alertas para 429 e spikes de custo.