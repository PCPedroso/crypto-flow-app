# FEATURE: Detect and fix Backpack connection

Filepath: docs/features/EXAMPLE_CONNECT_BACKPACK.md
Purpose: Documentação de correção e testes aplicados para o conector Backpack.
Owner: @dev
Status: DONE
Created: 2025-08-30
Updated: 2025-08-30

Resumo
- Corrigida detecção `typeof window.backpack !== 'undefined'`
- Conexão via `window.backpack.solana.connect()` verificada localmente

Acceptance
- [x] WALLETS.md atualizado
- [x] Código commitado e testado manualmente