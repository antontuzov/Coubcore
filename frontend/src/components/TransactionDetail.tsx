import React from 'react';

interface TransactionInput {
  txid: string;
  vout: number;
  scriptSig: {
    asm: string;
    hex: string;
  };
  sequence: number;
}

interface TransactionOutput {
  value: number;
  n: number;
  scriptPubKey: {
    asm: string;
    hex: string;
    reqSigs: number;
    type: string;
    addresses: string[];
  };
}

interface TransactionDetailProps {
  transaction: {
    txid: string;
    hash: string;
    version: number;
    size: number;
    vsize: number;
    weight: number;
    locktime: number;
    vin: TransactionInput[];
    vout: TransactionOutput[];
    hex: string;
    blockhash: string;
    confirmations: number;
    time: number;
    blocktime: number;
  };
  onBack?: () => void;
}

const TransactionDetail: React.FC<TransactionDetailProps> = ({ transaction, onBack }) => {
  const formatTime = (timestamp: number) => {
    return new Date(timestamp * 1000).toLocaleString();
  };

  return (
    <div className="transaction-detail bg-gray-800 rounded-xl p-6 shadow-lg backdrop-blur-sm bg-opacity-70 border border-gray-700">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-white bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
          Transaction Details
        </h2>
        {onBack && (
          <button
            onClick={onBack}
            className="text-blue-400 hover:text-blue-300 transition-colors"
          >
            ← Back to Transactions
          </button>
        )}
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <div className="bg-gray-900 bg-opacity-50 rounded-lg p-4 border border-gray-700">
          <h3 className="text-lg font-semibold text-blue-400 mb-2">Transaction Info</h3>
          <div className="space-y-2">
            <div>
              <span className="text-gray-400">TxID:</span>
              <span className="text-white ml-2 break-all">{transaction.txid}</span>
            </div>
            <div>
              <span className="text-gray-400">Confirmations:</span>
              <span className="text-white ml-2">{transaction.confirmations}</span>
            </div>
            <div>
              <span className="text-gray-400">Size:</span>
              <span className="text-white ml-2">{transaction.size} bytes</span>
            </div>
            <div>
              <span className="text-gray-400">Weight:</span>
              <span className="text-white ml-2">{transaction.weight}</span>
            </div>
            <div>
              <span className="text-gray-400">Time:</span>
              <span className="text-white ml-2">{formatTime(transaction.time)}</span>
            </div>
          </div>
        </div>

        <div className="bg-gray-900 bg-opacity-50 rounded-lg p-4 border border-gray-700">
          <h3 className="text-lg font-semibold text-blue-400 mb-2">Block Info</h3>
          <div className="space-y-2">
            <div>
              <span className="text-gray-400">Block Hash:</span>
              <span className="text-white ml-2 break-all">{transaction.blockhash}</span>
            </div>
            <div>
              <span className="text-gray-400">Block Time:</span>
              <span className="text-white ml-2">{formatTime(transaction.blocktime)}</span>
            </div>
          </div>
        </div>
      </div>

      <div className="mb-8">
        <h3 className="text-xl font-semibold text-white mb-4">Inputs</h3>
        <div className="space-y-4">
          {transaction.vin.map((input, index) => (
            <div key={index} className="bg-gray-900 bg-opacity-50 rounded-lg p-4 border border-gray-700">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <span className="text-gray-400">Previous TxID:</span>
                  <span className="text-white ml-2 break-all">{input.txid || 'Coinbase'}</span>
                </div>
                <div>
                  <span className="text-gray-400">Output Index:</span>
                  <span className="text-white ml-2">{input.vout !== undefined ? input.vout : 'N/A'}</span>
                </div>
                <div className="md:col-span-2">
                  <span className="text-gray-400">Signature Script:</span>
                  <div className="mt-1 text-white break-all bg-black bg-opacity-30 p-2 rounded">
                    {input.scriptSig?.hex || 'N/A'}
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>

      <div>
        <h3 className="text-xl font-semibold text-white mb-4">Outputs</h3>
        <div className="space-y-4">
          {transaction.vout.map((output, index) => (
            <div key={index} className="bg-gray-900 bg-opacity-50 rounded-lg p-4 border border-gray-700">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <span className="text-gray-400">Amount:</span>
                  <span className="text-white ml-2">{output.value} COUB</span>
                </div>
                <div>
                  <span className="text-gray-400">Index:</span>
                  <span className="text-white ml-2">{output.n}</span>
                </div>
                <div className="md:col-span-2">
                  <span className="text-gray-400">Addresses:</span>
                  <div className="mt-1 space-y-1">
                    {output.scriptPubKey.addresses?.map((address, addrIndex) => (
                      <div key={addrIndex} className="text-white break-all bg-black bg-opacity-30 p-2 rounded">
                        {address}
                      </div>
                    )) || <span className="text-gray-500">No addresses</span>}
                  </div>
                </div>
                <div className="md:col-span-2">
                  <span className="text-gray-400">Public Key Script:</span>
                  <div className="mt-1 text-white break-all bg-black bg-opacity-30 p-2 rounded">
                    {output.scriptPubKey.hex}
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default TransactionDetail;