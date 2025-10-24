# Hardware Wallet Integration Guide

## Overview

This document describes how to integrate hardware wallet support into the Coubcore frontend application, allowing users to securely manage their cryptocurrency assets.

## Supported Hardware Wallets

1. **Ledger**: Ledger Nano S, Ledger Nano X, Ledger Stax
2. **Trezor**: Trezor One, Trezor Model T
3. **KeepKey**: KeepKey hardware wallet

## Implementation

### Ledger Integration

Ledger wallets can be integrated using the Ledger JavaScript libraries:

```bash
npm install @ledgerhq/hw-transport-webusb @ledgerhq/hw-app-eth
```

Basic Ledger connection:

```ts
// utils/ledger.ts
import TransportWebUSB from '@ledgerhq/hw-transport-webusb';
import Eth from '@ledgerhq/hw-app-eth';

export class LedgerWallet {
  private transport: any = null;
  private ethApp: any = null;

  async connect() {
    try {
      this.transport = await TransportWebUSB.create();
      this.ethApp = new Eth(this.transport);
      return true;
    } catch (error) {
      console.error('Ledger connection failed:', error);
      return false;
    }
  }

  async getAddress() {
    if (!this.ethApp) {
      throw new Error('Ledger not connected');
    }

    try {
      const result = await this.ethApp.getAddress("44'/60'/0'/0/0");
      return result.address;
    } catch (error) {
      console.error('Failed to get address:', error);
      throw error;
    }
  }

  async signTransaction(tx: any) {
    if (!this.ethApp) {
      throw new Error('Ledger not connected');
    }

    try {
      const result = await this.ethApp.signTransaction("44'/60'/0'/0/0", tx);
      return result;
    } catch (error) {
      console.error('Failed to sign transaction:', error);
      throw error;
    }
  }

  async disconnect() {
    if (this.transport) {
      await this.transport.close();
      this.transport = null;
      this.ethApp = null;
    }
  }
}
```

### Trezor Integration

Trezor wallets can be integrated using the Trezor Connect library:

```bash
npm install trezor-connect
```

Basic Trezor connection:

```ts
// utils/trezor.ts
import TrezorConnect from 'trezor-connect';

TrezorConnect.manifest({
  email: 'developer@xyz.com',
  appUrl: 'https://yourapplication.com'
});

export class TrezorWallet {
  async connect() {
    try {
      const response = await TrezorConnect.getPublicKey({
        path: "m/44'/60'/0'/0/0",
        coin: 'eth'
      });

      if (response.success) {
        return true;
      } else {
        console.error('Trezor connection failed:', response.payload.error);
        return false;
      }
    } catch (error) {
      console.error('Trezor connection failed:', error);
      return false;
    }
  }

  async getAddress() {
    try {
      const response = await TrezorConnect.ethereumGetAddress({
        path: "m/44'/60'/0'/0/0",
        showOnTrezor: true
      });

      if (response.success) {
        return response.payload.address;
      } else {
        throw new Error(response.payload.error);
      }
    } catch (error) {
      console.error('Failed to get address:', error);
      throw error;
    }
  }

  async signTransaction(tx: any) {
    try {
      const response = await TrezorConnect.ethereumSignTransaction({
        path: "m/44'/60'/0'/0/0",
        transaction: tx
      });

      if (response.success) {
        return response.payload;
      } else {
        throw new Error(response.payload.error);
      }
    } catch (error) {
      console.error('Failed to sign transaction:', error);
      throw error;
    }
  }
}
```

### KeepKey Integration

KeepKey wallets can be integrated using the KeepKey SDK:

```bash
npm install @keepkey/keepkey.js
```

Basic KeepKey connection:

```ts
// utils/keepkey.ts
import { KeepKeySdk } from '@keepkey/keepkey.js';

export class KeepKeyWallet {
  private sdk: any = null;

  async connect() {
    try {
      this.sdk = await KeepKeySdk.create({
        apiKey: 'your-api-key',
        pairingInfo: {
          name: 'Coubcore Wallet',
          imageUrl: 'https://yourapp.com/icon.png',
          basePath: 'https://yourapp.com',
        }
      });
      return true;
    } catch (error) {
      console.error('KeepKey connection failed:', error);
      return false;
    }
  }

  async getAddress() {
    if (!this.sdk) {
      throw new Error('KeepKey not connected');
    }

    try {
      const response = await this.sdk.address.ethereumGetAddress({
        addressNList: [2147483692, 2147483708, 2147483648, 0, 0],
        showDisplay: true
      });
      return response.address;
    } catch (error) {
      console.error('Failed to get address:', error);
      throw error;
    }
  }

  async signTransaction(tx: any) {
    if (!this.sdk) {
      throw new Error('KeepKey not connected');
    }

    try {
      const response = await this.sdk.eth.ethereumSignTx(tx);
      return response;
    } catch (error) {
      console.error('Failed to sign transaction:', error);
      throw error;
    }
  }
}
```

## Hardware Wallet Context

A context to manage hardware wallet connections:

```tsx
// contexts/HardwareWalletContext.tsx
import React, { createContext, useContext, useState } from 'react';
import { LedgerWallet } from '../utils/ledger';
import { TrezorWallet } from '../utils/trezor';
import { KeepKeyWallet } from '../utils/keepkey';

type WalletType = 'ledger' | 'trezor' | 'keepkey' | null;

interface HardwareWalletContextType {
  walletType: WalletType;
  isConnected: boolean;
  connectWallet: (type: WalletType) => Promise<boolean>;
  disconnectWallet: () => void;
  getAddress: () => Promise<string>;
  signTransaction: (tx: any) => Promise<any>;
}

const HardwareWalletContext = createContext<HardwareWalletContextType | undefined>(undefined);

export const HardwareWalletProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [walletType, setWalletType] = useState<WalletType>(null);
  const [isConnected, setIsConnected] = useState(false);
  const [walletInstance, setWalletInstance] = useState<any>(null);

  const connectWallet = async (type: WalletType): Promise<boolean> => {
    if (!type) {
      return false;
    }

    try {
      let wallet: any;
      
      switch (type) {
        case 'ledger':
          wallet = new LedgerWallet();
          break;
        case 'trezor':
          wallet = new TrezorWallet();
          break;
        case 'keepkey':
          wallet = new KeepKeyWallet();
          break;
        default:
          return false;
      }

      const success = await wallet.connect();
      if (success) {
        setWalletType(type);
        setIsConnected(true);
        setWalletInstance(wallet);
        return true;
      }
      return false;
    } catch (error) {
      console.error('Failed to connect wallet:', error);
      return false;
    }
  };

  const disconnectWallet = async () => {
    if (walletInstance) {
      await walletInstance.disconnect();
    }
    setWalletType(null);
    setIsConnected(false);
    setWalletInstance(null);
  };

  const getAddress = async (): Promise<string> => {
    if (!walletInstance) {
      throw new Error('No wallet connected');
    }
    return await walletInstance.getAddress();
  };

  const signTransaction = async (tx: any): Promise<any> => {
    if (!walletInstance) {
      throw new Error('No wallet connected');
    }
    return await walletInstance.signTransaction(tx);
  };

  const contextValue: HardwareWalletContextType = {
    walletType,
    isConnected,
    connectWallet,
    disconnectWallet,
    getAddress,
    signTransaction
  };

  return (
    <HardwareWalletContext.Provider value={contextValue}>
      {children}
    </HardwareWalletContext.Provider>
  );
};

export const useHardwareWallet = (): HardwareWalletContextType => {
  const context = useContext(HardwareWalletContext);
  
  if (context === undefined) {
    throw new Error('useHardwareWallet must be used within a HardwareWalletProvider');
  }
  
  return context;
};
```

## Hardware Wallet Components

### Wallet Connection Component

A component to connect hardware wallets:

```tsx
// components/HardwareWalletConnect.tsx
import React, { useState } from 'react';
import { useHardwareWallet } from '../contexts/HardwareWalletContext';

const HardwareWalletConnect: React.FC = () => {
  const { walletType, isConnected, connectWallet, disconnectWallet } = useHardwareWallet();
  const [connecting, setConnecting] = useState(false);

  const handleConnect = async (type: 'ledger' | 'trezor' | 'keepkey') => {
    setConnecting(true);
    try {
      await connectWallet(type);
    } finally {
      setConnecting(false);
    }
  };

  if (isConnected) {
    return (
      <div className="hardware-wallet-connected">
        <p>Connected to {walletType} wallet</p>
        <button onClick={disconnectWallet} className="btn btn-secondary">
          Disconnect
        </button>
      </div>
    );
  }

  return (
    <div className="hardware-wallet-connect">
      <h3>Connect Hardware Wallet</h3>
      <div className="wallet-options">
        <button 
          onClick={() => handleConnect('ledger')}
          disabled={connecting}
          className="btn btn-primary"
        >
          {connecting ? 'Connecting...' : 'Connect Ledger'}
        </button>
        <button 
          onClick={() => handleConnect('trezor')}
          disabled={connecting}
          className="btn btn-primary"
        >
          {connecting ? 'Connecting...' : 'Connect Trezor'}
        </button>
        <button 
          onClick={() => handleConnect('keepkey')}
          disabled={connecting}
          className="btn btn-primary"
        >
          {connecting ? 'Connecting...' : 'Connect KeepKey'}
        </button>
      </div>
    </div>
  );
};

export default HardwareWalletConnect;
```

### Transaction Signing Component

A component to sign transactions with hardware wallets:

```tsx
// components/HardwareSignTransaction.tsx
import React, { useState } from 'react';
import { useHardwareWallet } from '../contexts/HardwareWalletContext';

interface HardwareSignTransactionProps {
  transaction: any;
  onSigned: (signature: any) => void;
  onCancel: () => void;
}

const HardwareSignTransaction: React.FC<HardwareSignTransactionProps> = ({ 
  transaction, 
  onSigned, 
  onCancel 
}) => {
  const { signTransaction } = useHardwareWallet();
  const [signing, setSigning] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSign = async () => {
    setSigning(true);
    setError(null);
    
    try {
      const signature = await signTransaction(transaction);
      onSigned(signature);
    } catch (err) {
      setError('Failed to sign transaction. Please try again.');
      console.error('Signing error:', err);
    } finally {
      setSigning(false);
    }
  };

  return (
    <div className="hardware-sign-transaction">
      <h3>Sign Transaction</h3>
      <p>Please confirm the transaction on your hardware wallet device.</p>
      
      {error && (
        <div className="error-message">
          {error}
        </div>
      )}
      
      <div className="actions">
        <button 
          onClick={handleSign}
          disabled={signing}
          className="btn btn-primary"
        >
          {signing ? 'Signing...' : 'Sign with Hardware Wallet'}
        </button>
        <button 
          onClick={onCancel}
          disabled={signing}
          className="btn btn-secondary"
        >
          Cancel
        </button>
      </div>
    </div>
  );
};

export default HardwareSignTransaction;
```

## Security Considerations

1. **Secure Communication**: Ensure all communication with hardware wallets is secure
2. **Transaction Verification**: Always verify transactions on the hardware wallet display
3. **Firmware Updates**: Encourage users to keep their hardware wallets updated
4. **Backup Phrases**: Educate users about securely storing backup phrases
5. **Phishing Protection**: Implement measures to prevent phishing attacks

## User Experience

1. **Clear Instructions**: Provide clear step-by-step instructions for each wallet type
2. **Error Handling**: Handle connection errors gracefully with helpful messages
3. **Loading States**: Show appropriate loading states during connection and signing
4. **Device Detection**: Automatically detect connected hardware wallets when possible
5. **Fallback Options**: Provide fallback options if hardware wallet connection fails

## Testing

1. **Unit Tests**: Test each wallet integration separately
2. **Integration Tests**: Test the complete connection and signing flow
3. **Browser Compatibility**: Test across different browsers and operating systems
4. **Device Testing**: Test with actual hardware devices
5. **Edge Cases**: Test edge cases like disconnection during signing

## Future Enhancements

1. **Multi-Signature Support**: Support for multi-signature transactions
2. **WalletConnect Integration**: Integrate with WalletConnect for mobile wallets
3. **Custom Firmware**: Support for custom firmware implementations
4. **Advanced Features**: Support for advanced wallet features like staking
5. **Biometric Authentication**: Integration with biometric authentication when supported