export function setupMetaMask(walletPopup) {
    const connectMetaMaskButton = document.getElementById('connectMetaMask');

    connectMetaMaskButton.addEventListener('click', async function() {
        if (typeof window.ethereum !== 'undefined') {
            try {
                // Solicita acesso à conta MetaMask
                const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
                console.log('Conta conectada:', accounts[0]);
                alert('MetaMask conectada! Conta: ' + accounts[0]);
                walletPopup.style.display = 'none';
            } catch (error) {
                console.error('Erro ao conectar MetaMask:', error);
                alert('Erro ao conectar MetaMask: ' + error.message);
            }
        } else {
            alert('MetaMask não detectada. Por favor, instale a extensão MetaMask.');
        }
    });
}