import React, { useState, useEffect } from 'react';
import { WalletInfo } from '../types/blockchain';
import { getBalance } from '../services/blockchainAPI';
import { validateAddress } from '../utils/crypto';

interface WalletProps {
  address?: string;
  balance?: number;
  onWalletCreate?: () => void;
  onTransactionSend?: (to: string, amount: number) => void;
}

const Wallet: React.FC<WalletProps> = ({ address, balance, onWalletCreate, onTransactionSend }) => {
  const [walletInfo, setWalletInfo] = useState<WalletInfo | null>(null);
  const [isCreating, setIsCreating] = useState(false);
  const [isSending, setIsSending] = useState(false);
  const [recipient, setRecipient] = useState('');
  const [amount, setAmount] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  // Load wallet info when address changes
  useEffect(() => {
    if (address) {
      loadWalletInfo(address);
    }
  }, [address]);

  const loadWalletInfo = async (walletAddress: string) => {
    try {
      const balanceData = await getBalance(walletAddress);
      setWalletInfo({
        address: walletAddress,
        balance: balanceData.balance || 0,
        publicKey: '', // In a real implementation, we would get this from the wallet
      });
    } catch (err) {
      setError('Failed to load wallet information');
      console.error(err);
    }
  };

  const handleCreateWallet = () => {
    setIsCreating(true);
    // In a real implementation, we would generate a new wallet
    // For now, we'll just call the onWalletCreate callback
    if (onWalletCreate) {
      onWalletCreate();
    }
    setIsCreating(false);
  };

  const handleSendTransaction = () => {
    // Validate inputs
    if (!recipient || !amount) {
      setError('Please fill in all fields');
      return;
    }

    if (!validateAddress(recipient)) {
      setError('Invalid recipient address');
      return;
    }

    const amountNum = parseFloat(amount);
    if (isNaN(amountNum) || amountNum <= 0) {
      setError('Invalid amount');
      return;
    }

    if (walletInfo && amountNum > walletInfo.balance) {
      setError('Insufficient balance');
      return;
    }

    setIsSending(true);
    setError('');
    setSuccess('');

    // In a real implementation, we would sign and send the transaction
    // For now, we'll just call the onTransactionSend callback
    if (onTransactionSend) {
      onTransactionSend(recipient, amountNum);
      setSuccess('Transaction sent successfully!');
      setRecipient('');
      setAmount('');
    }

    setIsSending(false);
  };

  const formatBalance = (balance: number) => {
    return balance.toFixed(4);
  };

  return (
    <div className="wallet-container bg-gray-800 rounded-xl p-6 shadow-lg backdrop-blur-sm bg-opacity-70 border border-gray-700">
      <h2 className="text-2xl font-bold text-white mb-6 text-center bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
        Wallet
      </h2>

      {!address ? (
        <div className="text-center">
          <p className="text-gray-300 mb-6">No wallet connected</p>
          <button
            onClick={handleCreateWallet}
            disabled={isCreating}
            className="bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300 transform hover:scale-105 disabled:opacity-50"
          >
            {isCreating ? 'Creating...' : 'Create New Wallet'}
          </button>
        </div>
      ) : (
        <div>
          {/* Wallet Info */}
          <div className="mb-6">
            <div className="flex justify-between items-center mb-2">
              <span className="text-gray-400">Address:</span>
              <span className="text-blue-300 font-mono text-sm truncate max-w-[200px]">
                {address}
              </span>
            </div>
            <div className="flex justify-between items-center">
              <span className="text-gray-400">Balance:</span>
              <span className="text-2xl font-bold bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
                {formatBalance(balance || walletInfo?.balance || 0)} COUB
              </span>
            </div>
          </div>

          {/* Send Transaction Form */}
          <div className="border-t border-gray-700 pt-6">
            <h3 className="text-xl font-semibold text-white mb-4">Send Transaction</h3>
            {error && (
              <div className="bg-red-900 text-red-200 p-3 rounded-lg mb-4 text-sm">
                {error}
              </div>
            )}
            {success && (
              <div className="bg-green-900 text-green-200 p-3 rounded-lg mb-4 text-sm">
                {success}
              </div>
            )}
            <div className="space-y-4">
              <div>
                <label className="block text-gray-300 text-sm font-medium mb-2">
                  Recipient Address
                </label>
                <input
                  type="text"
                  value={recipient}
                  onChange={(e) => setRecipient(e.target.value)}
                  placeholder="Enter recipient address"
                  className="w-full px-4 py-2 bg-gray-700 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <div>
                <label className="block text-gray-300 text-sm font-medium mb-2">
                  Amount (COUB)
                </label>
                <input
                  type="number"
                  value={amount}
                  onChange={(e) => setAmount(e.target.value)}
                  placeholder="0.00"
                  step="0.0001"
                  className="w-full px-4 py-2 bg-gray-700 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <button
                onClick={handleSendTransaction}
                disabled={isSending}
                className="w-full bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 text-white font-bold py-3 px-6 rounded-lg transition-all duration-300 transform hover:scale-105 disabled:opacity-50"
              >
                {isSending ? 'Sending...' : 'Send Transaction'}
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default Wallet;