# INSTALL / INSTALAÇÃO

Filepath: docs/INSTALL.md
Purpose: Instruções para instalar e executar a aplicação localmente (Windows e Linux).
Owner: @dev
Status: DONE
Created: 2025-08-30
Updated: 2025-08-30

PT-BR
----

Requisitos:
- Go 1.20+ instalado
- Navegador moderno com extensões de carteiras (opcional para testes)

Passos para executar (Windows):
1. Abra PowerShell ou CMD na raiz do repositório.
2. Instale dependências:
   ```
   go mod tidy
   ```
3. Rode o servidor:
   ```
   go run cmd/main.go
   ```
4. Acesse:
   ```
   http://localhost:8080/
   ```

EN
--

Requirements:
- Go 1.20+ installed
- Modern browser with wallet extensions (optional for testing)

Run steps (Windows):
1. Open PowerShell or CMD at repo root.
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Run:
   ```
   go run cmd/main.go
   ```
4. Open:
   ```
   http://localhost:8080/
   ```

Short checklist
- [ ] Verificar variáveis de ambiente (se houver)
- [ ] Confirmar portas livres (8080)