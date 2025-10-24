import React from 'react';

interface ConnectionStatusProps {
  isConnected: boolean;
  nodeUrl: string;
}

const ConnectionStatus: React.FC<ConnectionStatusProps> = ({ isConnected, nodeUrl }) => {
  return (
    <div className="connection-status bg-gray-800 rounded-xl p-4 shadow-lg backdrop-blur-sm bg-opacity-70 border border-gray-700">
      <div className="flex items-center">
        <div className={`w-3 h-3 rounded-full mr-3 ${isConnected ? 'bg-green-500' : 'bg-red-500'}`}></div>
        <div>
          <div className="text-white font-medium">
            {isConnected ? 'Connected' : 'Disconnected'}
          </div>
          <div className="text-gray-400 text-sm truncate max-w-[200px]">
            {nodeUrl}
          </div>
        </div>
      </div>
    </div>
  );
};

export default ConnectionStatus;