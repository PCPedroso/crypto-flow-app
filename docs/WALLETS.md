# WALLETS — Carteiras suportadas / Wallet integrations

PT-BR
----
Conectores implementados (arquivos em `public/scripts/`):
- metamask.js
  - exporta `setupMetaMask(walletPopup)`
  - usa `window.ethereum.request({ method: 'eth_requestAccounts' })`
- phantom.js
  - exporta `setupPhantom(walletPopup)`
  - usa `window.phantom.solana.connect()`
- backpack.js
  - exporta `setupBackpack(walletPopup)`
  - usa `window.backpack` / `window.backpack.solana.connect()`
  - Observação: verificar detecção correta com `typeof window.backpack !== 'undefined'` ou `window.hasOwnProperty('backpack')`

Padrão de implementação:
- Cada módulo exporta uma função `setupX(walletPopup)` que registra event listeners nos botões existentes em `index.html`.
- `wallet-connector.js` importa e re-exporta os setups; `index.html` importa o módulo central.

Boas práticas:
- Antes de registar listeners, verifique `document.getElementById(...) !== null`.
- Substituir `alert()` por modais para melhor UX.
- Centralizar estado (ex.: `walletManager`) para gerir múltiplas conexões, desconexões e persistência.

EN
--
Implemented connectors (files in `public/scripts/`):
- metamask.js
  - exports `setupMetaMask(walletPopup)`
  - uses `window.ethereum.request({ method: 'eth_requestAccounts' })`
- phantom.js
  - exports `setupPhantom(walletPopup)`
  - uses `window.phantom.solana.connect()`
- backpack.js
  - exports `setupBackpack(walletPopup)`
  - uses `window.backpack` / `window.backpack.solana.connect()`
  - Note: detect `window.backpack` with `typeof window.backpack !== 'undefined'` or `window.hasOwnProperty('backpack')`

Implementation pattern:
- Each module exports `setupX(walletPopup)` that attaches event listeners to buttons in `index.html`.
- `wallet-connector.js` imports and re-exports; `index.html` imports the central module.

Best practices:
- Check DOM element presence before attaching listeners.
- Replace `alert()` calls with proper modal UI.
- Centralize wallet state in a manager object for connect/disconnect handling.

# ARCHITECTURE / ARQUITETURA

Filepath: docs/ARCHITECTURE.md
Purpose: Descrever arquitetura, rotas e mapa de arquivos estáticos.
Owner: @dev
Status: DONE
Created: 2025-08-30
Updated: 2025-08-30

Resumo
- Backend: Go (cmd/main.go) com gorilla/mux.
- Frontend: arquivos estáticos em `public/` (ES Modules em `public/scripts/`).
- Static file server: `r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))`
- Rota principal: GET `/` serve `public/index.html`.

Notes
- Scripts devem usar caminho `/static/scripts/...` com a configuração atual.
- Recomenda-se usar `http.ServeFile` no handler se preferir servir index sem parse template.