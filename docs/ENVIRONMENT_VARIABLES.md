# ENVIRONMENT VARIABLES

## Exemplo (.env.example)
APP_ENV=development
APP_PORT=8080
DATABASE_URL=postgres://user:pass@localhost:5432/dbname
JWT_SECRET=replace_this_with_secret
WALLET_PROVIDER_URL=https://provider.example

## Lista (obrigatório/opcional)
- DATABASE_URL (obrigatório)
- JWT_SECRET (obrigatório)
- WALLETS_API_KEY (opcional)
- LOG_LEVEL (opcional, default=info)

## Notas
- Nunca commitar variáveis sensíveis.
- Uso recomendado: montar via CI/CD ou secret manager.