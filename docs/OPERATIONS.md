# OPERATIONS — Limites, Monitoramento e Boas Práticas / Limits & Operations

PT-BR
----
Limites (contexto: GPT-5 mini Preview — plano free)
- Planos free/preview normalmente aplicam quotas dinâmicas: rate limits, quotas diárias/mensais, limites por requisição (tokens) e concorrência.
- Não há números fixos garantidos — consulte o painel do provedor.

Recomendações de operação:
- Logger: registre headers de respostas de API para capturar `Retry-After`, `X-RateLimit-*`.
- Retry: implemente retry exponencial e leitura de `Retry-After` para 429.
- Rate limiting: throttle no servidor por usuário/chave (golang.org/x/time/rate ou semáforo).
- Concurrency: limite chamadas paralelas ao provedor (semaphore).
- Monitoramento: exporte métricas (Prometheus) e configure alertas sobre 429, custo e latência.

Snippets úteis (Go — veja exemplos no README principal)
- Retry/backoff
- Semaphore para concorrência
- Rate limiter com `rate.NewLimiter`

EN
--
Limits (context: GPT-5 mini Preview — free tier)
- Free/preview tiers usually enforce dynamic quotas: rate limits, daily/monthly quotas, per-request token limits and concurrency limits.
- No fixed numbers guaranteed — check provider dashboard.

Operational recommendations:
- Logger: capture API response headers to read `Retry-After`, `X-RateLimit-*`.
- Retry: implement exponential backoff and honor `Retry-After` on 429s.
- Rate limiting: throttle in-server per user/key (use golang.org/x/time/rate or semaphore).
- Concurrency: cap parallel outbound calls to provider (semaphore).
- Monitoring: expose metrics (Prometheus) and alert on 429s, cost spikes, latency.

If you want, I can add an `internal/middleware/` example implementing retry + rate-limit.