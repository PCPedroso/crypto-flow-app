# ARCHITECTURE — Arquitetura

PT-BR
----
Resumo:
- Backend: servidor HTTP em Go (cmd/main.go) usando gorilla/mux.
- Frontend: arquivos estáticos servidos a partir da pasta `public/` (HTML + ES Modules).
- Estrutura de rotas estáticas: o servidor mapeia `/static/*` para `./public` com:
  ```
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
  ```
- index.html é servido pela rota GET `/`.

Observações:
- Quando usar `StripPrefix("/static/", ...)`, os paths públicos no HTML devem usar o prefixo `/static/` (ex.: `/static/scripts/wallet-connector.js`).
- Alternativa para evitar inconsistências: servir `index.html` com `http.ServeFile(w, r, "public/index.html")` e configurar FileServer adequadamente.

EN
--
Summary:
- Backend: Go HTTP server (cmd/main.go) using gorilla/mux.
- Frontend: static files in `public/` (HTML + ES Modules).
- Static routing: server maps `/static/*` to `./public`:
  ```
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
  ```
- `index.html` is served on GET `/`.

Notes:
- With `StripPrefix("/static/", ...)` ensure HTML script tags reference `/static/...`.
- Optionally use `http.ServeFile` for `index.html` to avoid template-relative path issues.