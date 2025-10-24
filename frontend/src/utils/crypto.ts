// Utility functions for cryptographic operations

// Convert a string to a hex representation
export const stringToHex = (str: string): string => {
  let hex = '';
  for (let i = 0; i < str.length; i++) {
    const charCode = str.charCodeAt(i);
    const hexValue = charCode.toString(16);
    hex += hexValue.padStart(2, '0');
  }
  return hex;
};

// Convert hex to a string
export const hexToString = (hex: string): string => {
  let str = '';
  for (let i = 0; i < hex.length; i += 2) {
    const hexValue = hex.substr(i, 2);
    const charCode = parseInt(hexValue, 16);
    str += String.fromCharCode(charCode);
  }
  return str;
};

// Generate a simple hash (for demonstration purposes only)
export const simpleHash = (data: string): string => {
  let hash = 0;
  for (let i = 0; i < data.length; i++) {
    const char = data.charCodeAt(i);
    hash = ((hash << 5) - hash) + char;
    hash = hash & hash; // Convert to 32bit integer
  }
  return Math.abs(hash).toString(16);
};

// Validate an address format (basic validation)
export const validateAddress = (address: string): boolean => {
  // Basic validation - check if it's a hex string of appropriate length
  const hexRegex = /^[0-9a-fA-F]+$/;
  return hexRegex.test(address) && address.length === 40;
};

export default {
  stringToHex,
  hexToString,
  simpleHash,
  validateAddress,
};