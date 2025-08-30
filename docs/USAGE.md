# USAGE / COMO USAR

Filepath: docs/USAGE.md
Purpose: Exemplos práticos de uso e debug rápido.
Owner: @dev
Status: DONE
Created: 2025-08-30
Updated: 2025-08-30

Quick usage
- Start server: `go run cmd/main.go`
- Open `http://localhost:8080/`
- Use "Carteiras" popup para conectar MetaMask / Phantom / Backpack

Debug tips
- 404 em `.js`: confirmar presença em `public/scripts/` e uso do prefixo `/static/` se FileServer configurado assim.
- Ver logs no DevTools Console.