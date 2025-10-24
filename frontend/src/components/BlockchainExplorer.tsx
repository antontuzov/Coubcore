import React, { useState, useEffect } from 'react';
import { Block, Transaction } from '../types/blockchain';
import * as blockchainAPI from '../services/blockchainAPI';
import BlockList from './BlockList';
import BlockDetail from './BlockDetail';
import TransactionDetail from './TransactionDetail';
import NetworkStats from './NetworkStats';

const BlockchainExplorer: React.FC = () => {
  const [blocks, setBlocks] = useState<Block[]>([]);
  const [selectedBlock, setSelectedBlock] = useState<Block | null>(null);
  const [selectedTransaction, setSelectedTransaction] = useState<Transaction | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  // Fetch initial blocks
  useEffect(() => {
    const fetchBlocks = async () => {
      try {
        setLoading(true);
        // In a real implementation, we would fetch actual blocks
        // For now, we'll use sample data
        const sampleBlocks: Block[] = [
          {
            index: 100,
            timestamp: new Date().toISOString(),
            previousHash: '0000000000000000000a5d2a4a898d2a4a898d2a4a898d2a4a898d2a4a898d2a',
            hash: '0000000000000000000b5d2a4a898d2a4a898d2a4a898d2a4a898d2a4a898d2b',
            data: 'Block data 100',
            nonce: 12345,
            difficulty: 5,
            validator: 'validator_1'
          },
          {
            index: 99,
            timestamp: new Date(Date.now() - 1000 * 60 * 10).toISOString(),
            previousHash: '0000000000000000000a4d2a4a898d2a4a898d2a4a898d2a4a898d2a4a898d2a',
            hash: '0000000000000000000a5d2a4a898d2a4a898d2a4a898d2a4a898d2a4a898d2a',
            data: 'Block data 99',
            nonce: 23456,
            difficulty: 5,
            validator: 'validator_2'
          },
          {
            index: 98,
            timestamp: new Date(Date.now() - 1000 * 60 * 20).toISOString(),
            previousHash: '0000000000000000000a3d2a4a898d2a4a898d2a4a898d2a4a898d2a4a898d2a',
            hash: '0000000000000000000a4d2a4a898d2a4a898d2a4a898d2a4a898d2a4a898d2a',
            data: 'Block data 98',
            nonce: 34567,
            difficulty: 5,
            validator: 'validator_1'
          }
        ];
        setBlocks(sampleBlocks);
        setLoading(false);
      } catch (err) {
        setError('Failed to fetch blocks');
        setLoading(false);
      }
    };

    fetchBlocks();
  }, []);

  const handleBlockSelect = async (block: Block) => {
    setSelectedBlock(block);
    setSelectedTransaction(null);
  };

  const handleTransactionSelect = (transaction: Transaction) => {
    setSelectedTransaction(transaction);
  };

  const handleBackToBlocks = () => {
    setSelectedBlock(null);
  };

  const handleBackToBlockDetail = () => {
    setSelectedTransaction(null);
  };

  const handleSearch = async (query: string) => {
    try {
      // Check if query is a block index
      if (/^\d+$/.test(query)) {
        const blockIndex = parseInt(query, 10);
        // In a real implementation, we would fetch the block
        // const block = await blockchainAPI.getBlock(blockIndex);
        // For now, we'll use sample data
        const sampleBlock: Block = {
          index: blockIndex,
          timestamp: new Date().toISOString(),
          previousHash: '0000000000000000000a5d2a4a898d2a4a898d2a4a898d2a4a898d2a4a898d2a',
          hash: '0000000000000000000b5d2a4a898d2a4a898d2a4a898d2a4a898d2a4a898d2b',
          data: `Block data ${blockIndex}`,
          nonce: 12345,
          difficulty: 5,
          validator: 'validator_1'
        };
        setSelectedBlock(sampleBlock);
        setSelectedTransaction(null);
        return;
      }

      // Check if query is a transaction ID
      if (query.startsWith('0x') && query.length === 66) {
        // In a real implementation, we would fetch the transaction
        // const transaction = await blockchainAPI.getTransaction(query);
        // For now, we'll use sample data
        const sampleTransaction: Transaction = {
          id: query,
          inputs: [{
            txid: 'prev_tx_123',
            vout: 0,
            signature: 'sample_signature',
            pubKey: 'sample_public_key'
          }],
          outputs: [{
            value: 10,
            pubKeyHash: 'recipient_address'
          }],
          time: new Date().toISOString()
        };
        setSelectedTransaction(sampleTransaction);
        setSelectedBlock(null);
        return;
      }

      // If we reach here, the query didn't match any known formats
      setError('No block or transaction found with that identifier');
    } catch (err) {
      setError('Search failed. Please try again.');
    }
  };

  return (
    <div className="blockchain-explorer">
      <h1 className="text-3xl font-bold text-white mb-8 text-center bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
        Blockchain Explorer
      </h1>

      <NetworkStats onSearch={handleSearch} />

      {error && (
        <div className="mt-4 p-4 bg-red-900 bg-opacity-50 rounded-lg border border-red-700 text-red-200">
          {error}
        </div>
      )}

      <div className="mt-8">
        {!selectedBlock && !selectedTransaction && (
          <BlockList blocks={blocks} loading={loading} onBlockSelect={handleBlockSelect} />
        )}

        {selectedBlock && !selectedTransaction && (
          <BlockDetail block={selectedBlock} onBack={handleBackToBlocks} onTransactionSelect={handleTransactionSelect} />
        )}

        {selectedTransaction && (
          <TransactionDetail transaction={selectedTransaction} onBack={handleBackToBlockDetail} />
        )}
      </div>
    </div>
  );
};

export default BlockchainExplorer;