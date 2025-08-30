# USAGE — Como usar / Examples

PT-BR
----
Fluxo básico:
1. Inicie o servidor (ver INSTALL.md).
2. Abra `http://localhost:8080/`.
3. Clique em "Carteiras" para abrir o pop-up.
4. Escolha MetaMask / Phantom / Backpack.
5. Autorize a conexão na extensão do navegador.
6. Verifique o console do DevTools para logs e o endereço público conectado.

Dicas de debug:
- Se `.js` retornarem 404, confirme:
  - Arquivos presentes em `public/scripts/`
  - Caminho usado em `index.html` corresponde ao mapeamento estático (`/static/scripts/` se FileServer usa StripPrefix).
- Use DevTools → Console para ver erros e mensagens de `console.log`.

EN
--
Basic flow:
1. Start server (see INSTALL.md).
2. Open `http://localhost:8080/`.
3. Click "Wallets" to open popup.
4. Pick MetaMask / Phantom / Backpack.
5. Approve connection in the browser extension.
6. Check DevTools console for connected public address logs.

Debug tips:
- If `.js` return 404, verify:
  - Files exist under `public/scripts/`
  - `index.html` script paths match static mapping (`/static/scripts/` if FileServer uses StripPrefix).
- Use DevTools Console for errors and `console.log` output.