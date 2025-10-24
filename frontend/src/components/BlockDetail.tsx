import React from 'react';
import { Block } from '../types/blockchain';

interface BlockDetailProps {
  block: Block;
  onBack?: () => void;
}

const BlockDetail: React.FC<BlockDetailProps> = ({ block, onBack }) => {
  const formatTimestamp = (timestamp: string) => {
    return new Date(timestamp).toLocaleString();
  };

  return (
    <div className="block-detail bg-gray-800 rounded-xl p-6 shadow-lg backdrop-blur-sm bg-opacity-70 border border-gray-700">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-white bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
          Block #{block.index}
        </h2>
        {onBack && (
          <button
            onClick={onBack}
            className="text-blue-400 hover:text-blue-300 transition-colors"
          >
            ← Back to Blocks
          </button>
        )}
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="space-y-4">
          <div>
            <h3 className="text-lg font-semibold text-gray-300 mb-2">Block Information</h3>
            <div className="space-y-3">
              <div className="flex justify-between">
                <span className="text-gray-400">Timestamp:</span>
                <span className="text-white">{formatTimestamp(block.timestamp)}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-400">Nonce:</span>
                <span className="text-white">{block.nonce}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-400">Difficulty:</span>
                <span className="text-white">{block.difficulty}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-400">Validator:</span>
                <span className="text-green-400 font-mono truncate max-w-[200px]">
                  {block.validator}
                </span>
              </div>
            </div>
          </div>

          <div>
            <h3 className="text-lg font-semibold text-gray-300 mb-2">Hash Information</h3>
            <div className="space-y-3">
              <div>
                <span className="text-gray-400 block mb-1">Current Hash:</span>
                <span className="text-blue-300 font-mono text-sm break-all">
                  {block.hash}
                </span>
              </div>
              <div>
                <span className="text-gray-400 block mb-1">
                  {block.index > 0 ? 'Previous Hash:' : 'Genesis Block'}
                </span>
                {block.index > 0 ? (
                  <span className="text-purple-300 font-mono text-sm break-all">
                    {block.previousHash}
                  </span>
                ) : (
                  <span className="text-gray-500">This is the first block in the chain</span>
                )}
              </div>
            </div>
          </div>
        </div>

        <div>
          <h3 className="text-lg font-semibold text-gray-300 mb-2">Block Data</h3>
          <div className="bg-gray-750 rounded-lg p-4 h-full">
            <pre className="text-gray-300 text-sm overflow-auto max-h-60">
              {JSON.stringify(block.data, null, 2)}
            </pre>
          </div>
        </div>
      </div>
    </div>
  );
};

export default BlockDetail;