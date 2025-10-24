import axios from 'axios';

// API base URL - this should match your Go backend
const API_BASE_URL = 'http://localhost:8080/api/v1';

// Create an axios instance
const apiClient = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
});

// Get blockchain info
export const getBlockchainInfo = async () => {
  try {
    const response = await apiClient.get('/info');
    return response.data;
  } catch (error) {
    console.error('Error fetching blockchain info:', error);
    throw error;
  }
};

// Get a block by index
export const getBlock = async (index: number) => {
  try {
    const response = await apiClient.get(`/block?index=${index}`);
    return response.data;
  } catch (error) {
    console.error(`Error fetching block ${index}:`, error);
    throw error;
  }
};

// Get a transaction by ID
export const getTransaction = async (txid: string) => {
  try {
    const response = await apiClient.get(`/transaction?txid=${txid}`);
    return response.data;
  } catch (error) {
    console.error(`Error fetching transaction ${txid}:`, error);
    throw error;
  }
};

// Get balance for an address
export const getBalance = async (address: string) => {
  try {
    const response = await apiClient.get(`/balance?address=${address}`);
    return response.data;
  } catch (error) {
    console.error(`Error fetching balance for ${address}:`, error);
    throw error;
  }
};

// Get connected peers
export const getPeers = async () => {
  try {
    const response = await apiClient.get('/peers');
    return response.data;
  } catch (error) {
    console.error('Error fetching peers:', error);
    throw error;
  }
};

// Send a transaction
export const sendTransaction = async (transaction: any) => {
  try {
    const response = await apiClient.post('/send', transaction);
    return response.data;
  } catch (error) {
    console.error('Error sending transaction:', error);
    throw error;
  }
};

export default {
  getBlockchainInfo,
  getBlock,
  getTransaction,
  getBalance,
  getPeers,
  sendTransaction,
};