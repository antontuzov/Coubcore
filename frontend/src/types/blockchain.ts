// Type definitions for blockchain entities

export interface Block {
  index: number;
  timestamp: string;
  previousHash: string;
  hash: string;
  data: any;
  nonce: number;
  difficulty: number;
  validator: string;
}

export interface Transaction {
  id: string;
  inputs: TransactionInput[];
  outputs: TransactionOutput[];
  time: string;
}

export interface TransactionInput {
  txid: string;
  vout: number;
  signature: string;
  pubKey: string;
}

export interface TransactionOutput {
  value: number;
  pubKeyHash: string;
}

export interface WalletInfo {
  address: string;
  balance: number;
  publicKey: string;
}

export interface BlockchainInfo {
  length: number;
  latest: number;
}

export interface Peer {
  id: string;
  address: string;
  lastSeen: string;
}