import React, { useState, useEffect } from 'react';
import { Block } from '../types/blockchain';
import { getBlockchainInfo } from '../services/blockchainAPI';

interface BlockListProps {
  onBlockSelect?: (block: Block) => void;
}

const BlockList: React.FC<BlockListProps> = ({ onBlockSelect }) => {
  const [blocks, setBlocks] = useState<Block[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [chainLength, setChainLength] = useState(0);

  useEffect(() => {
    loadLatestBlocks();
  }, []);

  const loadLatestBlocks = async () => {
    try {
      setLoading(true);
      setError('');
      
      // Get blockchain info
      const info = await getBlockchainInfo();
      setChainLength(info.length);
      
      // For demo purposes, we'll create sample blocks
      // In a real implementation, we would fetch actual blocks from the API
      const sampleBlocks: Block[] = [];
      const latestIndex = info.latest;
      
      for (let i = latestIndex; i >= Math.max(0, latestIndex - 9); i--) {
        sampleBlocks.push({
          index: i,
          timestamp: new Date(Date.now() - (latestIndex - i) * 60000).toISOString(),
          previousHash: i > 0 ? `0x${Math.random().toString(16).substr(2, 64)}` : '',
          hash: `0x${Math.random().toString(16).substr(2, 64)}`,
          data: `Block data for block #${i}`,
          nonce: Math.floor(Math.random() * 1000000),
          difficulty: Math.floor(Math.random() * 10) + 1,
          validator: `Validator ${Math.floor(Math.random() * 100)}`,
        });
      }
      
      setBlocks(sampleBlocks);
    } catch (err) {
      setError('Failed to load blocks');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const formatTimestamp = (timestamp: string) => {
    return new Date(timestamp).toLocaleString();
  };

  const truncateHash = (hash: string) => {
    if (hash.length <= 20) return hash;
    return hash.substring(0, 10) + '...' + hash.substring(hash.length - 8);
  };

  return (
    <div className="block-list bg-gray-800 rounded-xl p-6 shadow-lg backdrop-blur-sm bg-opacity-70 border border-gray-700">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-white bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
          Latest Blocks
        </h2>
        <div className="text-gray-400 text-sm">
          Total: {chainLength}
        </div>
      </div>

      {error && (
        <div className="bg-red-900 text-red-200 p-3 rounded-lg mb-4">
          {error}
        </div>
      )}

      {loading ? (
        <div className="flex justify-center items-center h-32">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
        </div>
      ) : (
        <div className="space-y-3 max-h-96 overflow-y-auto">
          {blocks.map((block) => (
            <div 
              key={block.index}
              className="border border-gray-700 rounded-lg p-4 hover:bg-gray-750 transition-all duration-300 cursor-pointer transform hover:scale-[1.02]"
              onClick={() => onBlockSelect && onBlockSelect(block)}
            >
              <div className="flex justify-between items-start mb-2">
                <div>
                  <span className="text-blue-400 font-bold">#{block.index}</span>
                  <span className="text-gray-400 text-sm ml-2">
                    {formatTimestamp(block.timestamp)}
                  </span>
                </div>
                <span className="text-xs bg-gray-700 text-gray-300 px-2 py-1 rounded">
                  Diff: {block.difficulty}
                </span>
              </div>
              
              <div className="grid grid-cols-2 gap-2 text-sm mt-2">
                <div>
                  <span className="text-gray-400">Hash:</span>
                  <span className="text-blue-300 font-mono ml-2 truncate block">
                    {truncateHash(block.hash)}
                  </span>
                </div>
                <div>
                  <span className="text-gray-400">Prev:</span>
                  <span className="text-purple-300 font-mono ml-2 truncate block">
                    {block.index > 0 ? truncateHash(block.previousHash) : 'Genesis'}
                  </span>
                </div>
              </div>
              
              <div className="flex justify-between items-center mt-3 pt-2 border-t border-gray-700">
                <div className="text-sm">
                  <span className="text-gray-400">Nonce:</span>
                  <span className="text-white ml-2">{block.nonce}</span>
                </div>
                <div className="text-sm">
                  <span className="text-gray-400">Validator:</span>
                  <span className="text-green-400 ml-2 truncate max-w-[100px] inline-block">
                    {truncateHash(block.validator)}
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

export default BlockList;