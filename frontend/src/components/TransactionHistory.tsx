import React from 'react';
import { Transaction } from '../types/blockchain';

interface TransactionHistoryProps {
  transactions: Transaction[];
}

const TransactionHistory: React.FC<TransactionHistoryProps> = ({ transactions }) => {
  const formatTimestamp = (timestamp: string) => {
    return new Date(timestamp).toLocaleString();
  };

  const truncateString = (str: string, length: number = 10) => {
    if (str.length <= length) return str;
    return str.substring(0, length) + '...';
  };

  return (
    <div className="transaction-history bg-gray-800 rounded-xl p-6 shadow-lg backdrop-blur-sm bg-opacity-70 border border-gray-700">
      <h3 className="text-xl font-bold text-white mb-4 bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
        Transaction History
      </h3>
      
      {transactions.length === 0 ? (
        <p className="text-gray-400 text-center py-4">No transactions yet</p>
      ) : (
        <div className="space-y-3 max-h-96 overflow-y-auto">
          {transactions.map((tx) => (
            <div 
              key={tx.id} 
              className="border border-gray-700 rounded-lg p-4 hover:bg-gray-750 transition-colors duration-200"
            >
              <div className="flex justify-between items-start mb-2">
                <span className="text-blue-300 font-mono text-sm">
                  {truncateString(tx.id, 16)}
                </span>
                <span className="text-gray-400 text-sm">
                  {formatTimestamp(tx.time)}
                </span>
              </div>
              
              <div className="grid grid-cols-2 gap-2 text-sm">
                <div>
                  <span className="text-gray-400">Inputs:</span>
                  <span className="text-white ml-2">{tx.inputs.length}</span>
                </div>
                <div>
                  <span className="text-gray-400">Outputs:</span>
                  <span className="text-white ml-2">{tx.outputs.length}</span>
                </div>
              </div>
              
              <div className="mt-2 pt-2 border-t border-gray-700">
                <div className="flex justify-between">
                  <span className="text-gray-400">Total:</span>
                  <span className="text-green-400 font-medium">
                    {tx.outputs.reduce((sum, output) => sum + output.value, 0)} COUB
                  </span>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default TransactionHistory;