import React from 'react';

interface QRCodeProps {
  data: string;
  size?: number;
}

const QRCode: React.FC<QRCodeProps> = ({ data, size = 200 }) => {
  // In a real implementation, we would use a QR code library like qrcode.react
  // For now, we'll create a simple placeholder
  
  return (
    <div 
      className="qrcode-placeholder bg-gray-700 rounded-lg flex items-center justify-center border-2 border-dashed border-gray-600"
      style={{ width: size, height: size }}
    >
      <div className="text-center">
        <div className="bg-gray-600 border-2 border-gray-500 rounded w-16 h-16 mx-auto mb-2 flex items-center justify-center">
          <div className="grid grid-cols-3 gap-1">
            <div className="w-2 h-2 bg-white rounded-sm"></div>
            <div className="w-2 h-2 bg-white rounded-sm"></div>
            <div className="w-2 h-2 bg-white rounded-sm"></div>
            <div className="w-2 h-2 bg-white rounded-sm"></div>
            <div className="w-2 h-2 bg-black rounded-sm"></div>
            <div className="w-2 h-2 bg-white rounded-sm"></div>
            <div className="w-2 h-2 bg-white rounded-sm"></div>
            <div className="w-2 h-2 bg-white rounded-sm"></div>
            <div className="w-2 h-2 bg-white rounded-sm"></div>
          </div>
        </div>
        <p className="text-gray-400 text-xs">QR Code</p>
        <p className="text-gray-500 text-xs mt-1">Scan to pay</p>
      </div>
    </div>
  );
};

export default QRCode;