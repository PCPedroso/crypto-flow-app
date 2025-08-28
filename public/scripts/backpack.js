export function setupBackpack(walletPopup) {
    const connectBackpackButton = document.getElementById('connectBackpack');

    connectBackpackButton.addEventListener('click', async function() {
        if (window.hasOwnProperty('backpack') && typeof window.backpack.xnft !== 'undefined') {
            try {
                const resp = await window.backpack.xnft.connect();
                const publicKey = resp.publicKey.toString();
                console.log('Backpack conectado! Conta:', publicKey);
                alert('Backpack conectado! Conta: ' + publicKey);
                walletPopup.style.display = 'none';
            } catch (error) {
                console.error('Erro ao conectar Backpack:', error);
                alert('Erro ao conectar Backpack: ' + error.message);
            }
        } else {
            alert('Backpack não detectada. Por favor, instale a extensão Backpack.');
        }
    });
}