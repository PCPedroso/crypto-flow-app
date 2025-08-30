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