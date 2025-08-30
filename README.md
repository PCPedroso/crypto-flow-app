# Crypto Flow App / Aplicação Crypto Flow

Esta documentação descreve o estado atual da aplicação, como executar, a estrutura de pastas e as funcionalidades já implementadas. O conteúdo está disponível em Português (PT-BR) e em Inglês (EN). Leia a seção no idioma que preferir.

---

## PT-BR — Visão geral e documentação técnica

### Resumo do projeto
Crypto Flow App é uma aplicação web para gerenciar transações de criptomoedas com integração a carteiras no navegador (extensões). Backend em Go e frontend estático (HTML/JS) servido a partir da pasta `public/`.

### Funcionalidades já implementadas
- Backend
  - Servidor HTTP em Go (cmd/main.go) usando gorilla/mux.
  - Rota GET "/" que entrega `public/index.html`.
  - Servidor de arquivos estáticos configurado: requisições para `/static/...` são mapeadas para a pasta `./public` via `http.StripPrefix("/static/", http.FileServer(http.Dir("./public")))`.
- Frontend
  - Página principal em `public/index.html` com UI básica (header, dashboard).
  - Pop-up de seleção de carteiras com botões: MetaMask, Phantom, Backpack e Fechar.
  - Scripts modulares ES (ES Modules) em `public/scripts/`:
    - `metamask.js` — exporta `setupMetaMask(walletPopup)` para conectar via `window.ethereum`.
    - `phantom.js` — exporta `setupPhantom(walletPopup)` para conectar via `window.phantom.solana`.
    - `backpack.js` — exporta `setupBackpack(walletPopup)` para conectar via `window.backpack` (correção aplicada para detectar a extensão corretamente).
    - `wallet-connector.js` — ponto central que importa e re-exporta os setups.
  - `index.html` importa os módulos com `<script type="module"> import { ... } from '/static/scripts/wallet-connector.js'` e inicializa os listeners chamando `setupMetaMask`, `setupPhantom` e `setupBackpack`.

### Observações importantes e correções aplicadas
- Backpack: havia uma verificação incorreta (`window.backpack !== 'undefined'`) que comparava com string; corrigido para `typeof window.backpack !== 'undefined'` (ou `window.hasOwnProperty('backpack')`) e checagem correta de `window.backpack.solana`.
- Evitar importações duplicadas: apenas importe o módulo central (`wallet-connector.js`) no HTML e use esse arquivo para exportar as funções dos conectores.
- Módulos ES: as tags `<script>` que importam módulos devem usar `type="module"`.
- Rotas e caminhos:
  - Quando usar `http.StripPrefix("/static/", http.FileServer(http.Dir("./public")))`, os caminhos públicos no HTML devem apontar para `/static/...` (por exemplo `/static/scripts/wallet-connector.js`) ou ajustar o `StripPrefix`/FileServer conforme preferência.
  - Alternativa recomendada para servir index: `http.ServeFile(w, r, "public/index.html")` no `homeHandler` para evitar transformações de template que possam alterar caminhos relativos.

### Estrutura de pastas (atual)
```
crypto-flow-app/
├── cmd/
│   └── main.go
├── internal/
│   └── routes/
│       └── routes.go
├── public/
│   ├── index.html
│   ├── style.css (em /static/ se aplicável)
│   └── scripts/
│       ├── wallet-connector.js
│       ├── metamask.js
│       ├── phantom.js
│       └── backpack.js
├── go.mod
├── go.sum
└── README.md
```

### Como executar (Windows)
1. Abrir terminal (PowerShell / cmd) na raiz do repositório.
2. Instalar dependências Go:
   ```
   go mod tidy
   ```
3. Rodar:
   ```
   go run cmd/main.go
   ```
4. Abrir no navegador:
   ```
   http://localhost:8080/
   ```

### Como testar as carteiras (recomendações)
- Verifique se as extensões estão instaladas no navegador (MetaMask, Phantom, Backpack).
- Abra DevTools (F12) e a aba Console para ver logs (contas conectadas / erros).
- Ao clicar em "Carteiras" -> escolher uma carteira -> aceitar conexão na extensão.
- Se receber erro 404 nos `.js`:
  - Confirme que o HTML referencia `/static/scripts/...` se você usa `StripPrefix("/static/", FileServer("./public"))`.
  - Se preferir usar caminhos sem `/static/`, altere `StripPrefix` para servir a partir da raiz ou ajuste `src` das tags `<script>`.

### Boas práticas e recomendações de desenvolvimento
- Valide a existência dos botões antes de registrar listeners (evita erros quando elementos forem removidos).
- Centralize estado de conexão (ex.: objeto `walletManager`) para gerenciar múltiplas carteiras.
- Ao adicionar novos conectores, crie arquivo modular em `public/scripts/` e exporte uma função `setupX(walletPopup)`. Atualize `wallet-connector.js` para re-exportar.
- Para deploy, valide cabeçalhos MIME e cache para arquivos JS estáticos.

### Próximos passos sugeridos
- Persistir estado de carteira no backend (sessão segura) quando necessário.
- Implementar UI de logout/desconectar carteira.
- Centralizar logs e melhorar tratamento de erros (ex.: modais em vez de alert()).
- Adicionar testes unitários para lógica do backend.

---

## EN — Overview and technical documentation

### Project summary
Crypto Flow App is a web application to manage cryptocurrency transactions and browser-wallet integrations. Backend in Go; frontend is static files served from `public/`.

### Implemented features
- Backend
  - Go HTTP server (cmd/main.go) using gorilla/mux.
  - GET "/" route serving `public/index.html`.
  - Static file server configured: `/static/...` → `./public` via `http.StripPrefix("/static/", http.FileServer(http.Dir("./public")))`.
- Frontend
  - Main page `public/index.html` with a simple UI.
  - Wallet selection popup with MetaMask, Phantom, Backpack buttons.
  - ES Modules under `public/scripts/`:
    - `metamask.js` exports `setupMetaMask(walletPopup)` (uses `window.ethereum`).
    - `phantom.js` exports `setupPhantom(walletPopup)` (uses `window.phantom.solana`).
    - `backpack.js` exports `setupBackpack(walletPopup)` (fixed detection logic for `window.backpack`).
    - `wallet-connector.js` imports and re-exports the connectors.
  - `index.html` imports connectors via `<script type="module"> import { ... } from '/static/scripts/wallet-connector.js'` and initializes listeners.

### Important notes and fixes
- Backpack detection bug fixed: replaced incorrect string comparison with `typeof`.
- Avoid duplicate module imports; use `wallet-connector.js` as the single entry point.
- Use `type="module"` on `<script>` tags for ES module support.
- Ensure the HTML script paths align with how static files are served (`/static/` prefix or adjust server).

### File structure (current)
(see the PT-BR section for tree)

### How to run (Windows)
1. From repo root:
   ```
   go mod tidy
   go run cmd/main.go
   ```
2. Open:
   ```
   http://localhost:8080/
   ```

### Wallet testing tips
- Confirm browser extensions are installed.
- Use DevTools console for logs.
- If `.js` return 404, ensure HTML script paths match `FileServer` configuration (use `/static/scripts/...` with current server settings).

### Development recommendations
- Check DOM elements before registering event listeners.
- Centralize wallet state and error handling.
- Add new wallet connectors as modules and re-export via `wallet-connector.js`.

---