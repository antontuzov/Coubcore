import { ethers } from 'ethers';
import { WalletInfo } from '../types/blockchain';

// Wallet service for managing cryptocurrency wallets
class WalletService {
  private provider: ethers.JsonRpcProvider | null = null;
  private wallet: ethers.Wallet | null = null;

  // Initialize the wallet service
  async initialize(rpcUrl: string) {
    try {
      this.provider = new ethers.JsonRpcProvider(rpcUrl);
      await this.provider.ready;
      console.log('Wallet service initialized');
    } catch (error) {
      console.error('Error initializing wallet service:', error);
      throw error;
    }
  }

  // Create a new wallet
  async createWallet(): Promise<WalletInfo> {
    try {
      // Generate a new wallet
      this.wallet = ethers.Wallet.createRandom();
      
      // Get the address and public key
      const address = await this.wallet.getAddress();
      const publicKey = this.wallet.publicKey;
      
      // Get initial balance (should be 0 for new wallet)
      const balance = await this.getBalance(address);
      
      return {
        address,
        balance,
        publicKey
      };
    } catch (error) {
      console.error('Error creating wallet:', error);
      throw error;
    }
  }

  // Load wallet from private key
  async loadWallet(privateKey: string): Promise<WalletInfo> {
    try {
      if (!this.provider) {
        throw new Error('Wallet service not initialized');
      }
      
      // Load wallet from private key
      this.wallet = new ethers.Wallet(privateKey, this.provider);
      
      // Get the address and public key
      const address = await this.wallet.getAddress();
      const publicKey = this.wallet.publicKey;
      
      // Get balance
      const balance = await this.getBalance(address);
      
      return {
        address,
        balance,
        publicKey
      };
    } catch (error) {
      console.error('Error loading wallet:', error);
      throw error;
    }
  }

  // Get wallet balance
  async getBalance(address: string): Promise<number> {
    try {
      if (!this.provider) {
        throw new Error('Wallet service not initialized');
      }
      
      const balance = await this.provider.getBalance(address);
      // Convert from wei to ether (or your cryptocurrency unit)
      return parseFloat(ethers.formatEther(balance));
    } catch (error) {
      console.error(`Error fetching balance for ${address}:`, error);
      throw error;
    }
  }

  // Send transaction
  async sendTransaction(to: string, amount: number): Promise<string> {
    try {
      if (!this.wallet) {
        throw new Error('No wallet loaded');
      }
      
      if (!this.provider) {
        throw new Error('Wallet service not initialized');
      }
      
      // Convert amount to wei
      const amountWei = ethers.parseEther(amount.toString());
      
      // Create transaction
      const tx = await this.wallet.sendTransaction({
        to,
        value: amountWei
      });
      
      // Wait for transaction to be mined
      await tx.wait();
      
      return tx.hash;
    } catch (error) {
      console.error('Error sending transaction:', error);
      throw error;
    }
  }

  // Sign message
  async signMessage(message: string): Promise<string> {
    try {
      if (!this.wallet) {
        throw new Error('No wallet loaded');
      }
      
      return await this.wallet.signMessage(message);
    } catch (error) {
      console.error('Error signing message:', error);
      throw error;
    }
  }

  // Verify signature
  async verifySignature(message: string, signature: string, address: string): Promise<boolean> {
    try {
      const recoveredAddress = ethers.verifyMessage(message, signature);
      return recoveredAddress.toLowerCase() === address.toLowerCase();
    } catch (error) {
      console.error('Error verifying signature:', error);
      return false;
    }
  }

  // Get wallet info
  getWalletInfo(): WalletInfo | null {
    if (!this.wallet) {
      return null;
    }
    
    return {
      address: this.wallet.address,
      balance: 0, // Balance should be fetched separately
      publicKey: this.wallet.publicKey
    };
  }

  // Check if wallet is loaded
  isWalletLoaded(): boolean {
    return this.wallet !== null;
  }
}

export default WalletService;