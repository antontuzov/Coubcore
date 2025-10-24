import React, { useState } from 'react';
import { ThemeProvider } from './contexts/ThemeContext';
import Wallet from './components/Wallet';
import TransactionHistory from './components/TransactionHistory';
import AddressBook from './components/AddressBook';
import QRCode from './components/QRCode';
import ConnectionStatus from './components/ConnectionStatus';
import BlockchainExplorer from './components/BlockchainExplorer';
import { Transaction } from './types/blockchain';
import './App.css';

function App() {
  const [walletAddress, setWalletAddress] = useState<string>('');
  const [balance, setBalance] = useState<number>(0);
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [isConnected, setIsConnected] = useState<boolean>(true);
  const [activeTab, setActiveTab] = useState<'wallet' | 'explorer'>('wallet');

  const handleWalletCreate = () => {
    // In a real implementation, we would generate a new wallet
    // For now, we'll use a sample address
    setWalletAddress('0x1234567890abcdef1234567890abcdef12345678');
    setBalance(100.5);
  };

  const handleTransactionSend = (to: string, amount: number) => {
    // In a real implementation, we would send the transaction
    // For now, we'll just update the balance and add a transaction
    setBalance(prev => prev - amount);
    
    // Add a new transaction to the history
    const newTransaction: Transaction = {
      id: `tx_${Date.now()}`,
      inputs: [{ txid: 'input_tx', vout: 0, signature: 'sig', pubKey: 'pubkey' }],
      outputs: [{ value: amount, pubKeyHash: to }],
      time: new Date().toISOString(),
    };
    
    setTransactions(prev => [newTransaction, ...prev]);
  };

  const handleAddressSelect = (address: string) => {
    // In a real implementation, we would use the selected address
    console.log('Selected address:', address);
  };

  return (
    <ThemeProvider>
      <div className="min-h-screen bg-gradient-to-br from-gray-900 to-gray-800 text-white font-sans">
        {/* Header */}
        <header className="bg-gray-800 border-b border-gray-700">
          <div className="container mx-auto px-4 py-4 flex justify-between items-center">
            <h1 className="text-2xl font-bold bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
              Coubcore Blockchain
            </h1>
            <div className="flex items-center space-x-4">
              <div className="flex bg-gray-700 rounded-lg p-1">
                <button
                  className={`px-4 py-2 rounded-md transition-colors ${
                    activeTab === 'wallet'
                      ? 'bg-gradient-to-r from-blue-500 to-purple-600 text-white'
                      : 'text-gray-300 hover:text-white'
                  }`}
                  onClick={() => setActiveTab('wallet')}
                >
                  Wallet
                </button>
                <button
                  className={`px-4 py-2 rounded-md transition-colors ${
                    activeTab === 'explorer'
                      ? 'bg-gradient-to-r from-blue-500 to-purple-600 text-white'
                      : 'text-gray-300 hover:text-white'
                  }`}
                  onClick={() => setActiveTab('explorer')}
                >
                  Explorer
                </button>
              </div>
              <ConnectionStatus 
                isConnected={isConnected} 
                nodeUrl="http://localhost:8080" 
              />
            </div>
          </div>
        </header>

        {/* Main Content */}
        <main className="container mx-auto px-4 py-8">
          {activeTab === 'wallet' ? (
            <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
              {/* Left Column - Wallet and QR Code */}
              <div className="lg:col-span-1 space-y-8">
                <Wallet 
                  address={walletAddress}
                  balance={balance}
                  onWalletCreate={handleWalletCreate}
                  onTransactionSend={handleTransactionSend}
                />
                
                {walletAddress && (
                  <div className="bg-gray-800 rounded-xl p-6 shadow-lg backdrop-blur-sm bg-opacity-70 border border-gray-700">
                    <h3 className="text-xl font-bold text-white mb-4 bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
                      Wallet QR Code
                    </h3>
                    <div className="flex justify-center">
                      <QRCode data={walletAddress} size={180} />
                    </div>
                    <p className="text-gray-400 text-center mt-3 text-sm">
                      Scan to receive payments
                    </p>
                  </div>
                )}
              </div>

              {/* Middle Column - Transaction History */}
              <div className="lg:col-span-1">
                <TransactionHistory transactions={transactions} />
              </div>

              {/* Right Column - Address Book */}
              <div className="lg:col-span-1">
                <AddressBook onAddressSelect={handleAddressSelect} />
              </div>
            </div>
          ) : (
            <div className="explorer-container">
              <BlockchainExplorer />
            </div>
          )}
        </main>

        {/* Footer */}
        <footer className="bg-gray-800 border-t border-gray-700 py-6 mt-12">
          <div className="container mx-auto px-4 text-center text-gray-400">
            <p>Coubcore Blockchain &copy; {new Date().getFullYear()} - Created by Anton Tuzov</p>
          </div>
        </footer>
      </div>
    </ThemeProvider>
  );
}

export default App;