import React, { useState } from 'react';

interface AddressEntry {
  id: string;
  name: string;
  address: string;
}

interface AddressBookProps {
  onAddressSelect?: (address: string) => void;
}

const AddressBook: React.FC<AddressBookProps> = ({ onAddressSelect }) => {
  const [addresses, setAddresses] = useState<AddressEntry[]>([
    // Sample data - in a real app, this would come from local storage or a database
    { id: '1', name: 'My Wallet', address: '0x1234567890abcdef1234567890abcdef12345678' },
    { id: '2', name: 'Friend 1', address: '0xabcdef1234567890abcdef1234567890abcdef12' },
  ]);
  
  const [newName, setNewName] = useState('');
  const [newAddress, setNewAddress] = useState('');
  const [error, setError] = useState('');

  const handleAddAddress = () => {
    // Validate inputs
    if (!newName.trim() || !newAddress.trim()) {
      setError('Please fill in all fields');
      return;
    }

    // Check if address already exists
    if (addresses.some(addr => addr.address === newAddress)) {
      setError('Address already exists');
      return;
    }

    // Add new address
    const newEntry: AddressEntry = {
      id: Date.now().toString(),
      name: newName.trim(),
      address: newAddress.trim(),
    };

    setAddresses([...addresses, newEntry]);
    setNewName('');
    setNewAddress('');
    setError('');
  };

  const handleRemoveAddress = (id: string) => {
    setAddresses(addresses.filter(addr => addr.id !== id));
  };

  const handleSelectAddress = (address: string) => {
    if (onAddressSelect) {
      onAddressSelect(address);
    }
  };

  const truncateAddress = (address: string) => {
    if (address.length <= 20) return address;
    return address.substring(0, 10) + '...' + address.substring(address.length - 8);
  };

  return (
    <div className="address-book bg-gray-800 rounded-xl p-6 shadow-lg backdrop-blur-sm bg-opacity-70 border border-gray-700">
      <h3 className="text-xl font-bold text-white mb-4 bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
        Address Book
      </h3>
      
      {/* Add New Address Form */}
      <div className="mb-6 p-4 bg-gray-750 rounded-lg">
        <h4 className="text-lg font-semibold text-white mb-3">Add New Address</h4>
        {error && (
          <div className="bg-red-900 text-red-200 p-2 rounded mb-3 text-sm">
            {error}
          </div>
        )}
        <div className="space-y-3">
          <div>
            <label className="block text-gray-300 text-sm font-medium mb-1">
              Name
            </label>
            <input
              type="text"
              value={newName}
              onChange={(e) => setNewName(e.target.value)}
              placeholder="Enter name"
              className="w-full px-3 py-2 bg-gray-700 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div>
            <label className="block text-gray-300 text-sm font-medium mb-1">
              Address
            </label>
            <input
              type="text"
              value={newAddress}
              onChange={(e) => setNewAddress(e.target.value)}
              placeholder="Enter wallet address"
              className="w-full px-3 py-2 bg-gray-700 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <button
            onClick={handleAddAddress}
            className="w-full bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 text-white font-bold py-2 px-4 rounded-lg transition-all duration-300 transform hover:scale-105"
          >
            Add Address
          </button>
        </div>
      </div>
      
      {/* Address List */}
      <div>
        <h4 className="text-lg font-semibold text-white mb-3">Saved Addresses</h4>
        {addresses.length === 0 ? (
          <p className="text-gray-400 text-center py-4">No addresses saved</p>
        ) : (
          <div className="space-y-2 max-h-60 overflow-y-auto">
            {addresses.map((entry) => (
              <div 
                key={entry.id} 
                className="flex justify-between items-center p-3 bg-gray-750 rounded-lg hover:bg-gray-700 transition-colors duration-200"
              >
                <div 
                  className="cursor-pointer flex-1"
                  onClick={() => handleSelectAddress(entry.address)}
                >
                  <div className="font-medium text-white">{entry.name}</div>
                  <div className="text-blue-300 font-mono text-sm truncate">
                    {truncateAddress(entry.address)}
                  </div>
                </div>
                <button
                  onClick={() => handleRemoveAddress(entry.id)}
                  className="text-red-400 hover:text-red-300 ml-2"
                >
                  Remove
                </button>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default AddressBook;