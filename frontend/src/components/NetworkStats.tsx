import React, { useState, useEffect } from 'react';

interface NetworkStatsProps {
  onSearch?: (query: string) => void;
}

interface BlockchainInfo {
  chain: string;
  blocks: number;
  headers: number;
  bestblockhash: string;
  difficulty: number;
  mediantime: number;
  verificationprogress: number;
  initialblockdownload: boolean;
  chainwork: string;
  size_on_disk: number;
  pruned: boolean;
}

const NetworkStats: React.FC<NetworkStatsProps> = ({ onSearch }) => {
  const [blockchainInfo, setBlockchainInfo] = useState<BlockchainInfo | null>(null);
  const [loading, setLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');

  // In a real implementation, this would fetch from the API
  useEffect(() => {
    // Simulate API call
    setTimeout(() => {
      setBlockchainInfo({
        chain: 'main',
        blocks: 12483,
        headers: 12483,
        bestblockhash: '0000000000000000000a5d2a4a898d2a4a898d2a4a898d2a4a898d2a4a898d2a',
        difficulty: 123456789.123456,
        mediantime: 1678886400,
        verificationprogress: 0.9999,
        initialblockdownload: false,
        chainwork: '0000000000000000000000000000000000000000000000000000000000000000',
        size_on_disk: 4294967296,
        pruned: false
      });
      setLoading(false);
    }, 500);
  }, []);

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    if (onSearch && searchQuery.trim()) {
      onSearch(searchQuery.trim());
    }
  };

  if (loading) {
    return (
      <div className="network-stats bg-gray-800 rounded-xl p-6 shadow-lg backdrop-blur-sm bg-opacity-70 border border-gray-700">
        <div className="animate-pulse">
          <div className="h-6 bg-gray-700 rounded w-1/4 mb-4"></div>
          <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
            {[...Array(4)].map((_, i) => (
              <div key={i} className="h-20 bg-gray-700 rounded"></div>
            ))}
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="network-stats bg-gray-800 rounded-xl p-6 shadow-lg backdrop-blur-sm bg-opacity-70 border border-gray-700">
      <div className="flex flex-col md:flex-row md:items-center md:justify-between mb-6">
        <h2 className="text-2xl font-bold text-white bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent mb-4 md:mb-0">
          Network Statistics
        </h2>
        
        <form onSubmit={handleSearch} className="flex">
          <input
            type="text"
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            placeholder="Search block, transaction or address..."
            className="px-4 py-2 rounded-l-lg bg-gray-700 text-white focus:outline-none focus:ring-2 focus:ring-blue-500 w-full md:w-64"
          />
          <button
            type="submit"
            className="bg-gradient-to-r from-blue-500 to-purple-600 text-white px-4 py-2 rounded-r-lg hover:from-blue-600 hover:to-purple-700 transition-all"
          >
            Search
          </button>
        </form>
      </div>

      <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div className="bg-gray-900 bg-opacity-50 rounded-lg p-4 border border-gray-700">
          <div className="text-gray-400 text-sm">Blocks</div>
          <div className="text-2xl font-bold text-white">{blockchainInfo?.blocks.toLocaleString()}</div>
        </div>
        
        <div className="bg-gray-900 bg-opacity-50 rounded-lg p-4 border border-gray-700">
          <div className="text-gray-400 text-sm">Difficulty</div>
          <div className="text-2xl font-bold text-white">{blockchainInfo?.difficulty.toFixed(2)}</div>
        </div>
        
        <div className="bg-gray-900 bg-opacity-50 rounded-lg p-4 border border-gray-700">
          <div className="text-gray-400 text-sm">Chain</div>
          <div className="text-2xl font-bold text-white capitalize">{blockchainInfo?.chain}</div>
        </div>
        
        <div className="bg-gray-900 bg-opacity-50 rounded-lg p-4 border border-gray-700">
          <div className="text-gray-400 text-sm">Progress</div>
          <div className="text-2xl font-bold text-white">
            {(blockchainInfo ? blockchainInfo.verificationprogress * 100 : 0).toFixed(2)}%
          </div>
        </div>
      </div>

      <div className="mt-6 bg-gray-900 bg-opacity-50 rounded-lg p-4 border border-gray-700">
        <div className="text-gray-400 text-sm mb-2">Best Block Hash</div>
        <div className="text-white break-all text-sm font-mono">
          {blockchainInfo?.bestblockhash}
        </div>
      </div>
    </div>
  );
};

export default NetworkStats;