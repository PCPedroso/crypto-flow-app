export function setupPhantom(walletPopup) {
    const connectPhantomButton = document.getElementById('connectPhantom');

    connectPhantomButton.addEventListener('click', async function() {
        if (typeof window.phantom !== 'undefined' && typeof window.phantom.solana !== 'undefined') {
            try {
                const resp = await window.phantom.solana.connect();
                const publicKey = resp.publicKey.toString();
                console.log('Phantom conectado! Conta:', publicKey);
                alert('Phantom conectado! Conta: ' + publicKey);
                walletPopup.style.display = 'none';
            } catch (error) {
                console.error('Erro ao conectar Phantom:', error);
                alert('Erro ao conectar Phantom: ' + error.message);
            }
        } else {
            alert('Phantom não detectada. Por favor, instale a extensão Phantom.');
        }
    });
}