# TESTING — Guia rápido

## Tipos de teste
- Unit: rápido, isolado
- Integration: dependências reais (DB, serviços)
- E2E: fluxo completo (UI ou API)

## Comandos (exemplos)
- Instalar dependências:
  powershell -Command "npm ci"
- Rodar testes unitários:
  npm test
- Cobertura:
  npm run coverage

## CI
- Pipeline deve executar: install -> lint -> test -> integration (em runner com infra)