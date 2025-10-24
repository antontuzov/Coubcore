# Testing Guide for Coubcore Frontend

## Overview

This document describes how to test the Coubcore frontend components and services.

## Prerequisites

Make sure all dependencies are installed:

```bash
npm install
```

## Running Tests

To run all tests:

```bash
npm test
```

## Test Structure

Tests are organized in the following structure:

```
src/
├── components/
│   ├── __tests__/
│   │   ├── Wallet.test.tsx
│   │   ├── TransactionHistory.test.tsx
│   │   ├── AddressBook.test.tsx
│   │   ├── QRCode.test.tsx
│   │   ├── ConnectionStatus.test.tsx
│   │   ├── BlockList.test.tsx
│   │   ├── BlockDetail.test.tsx
│   │   ├── TransactionDetail.test.tsx
│   │   ├── NetworkStats.test.tsx
│   │   └── BlockchainExplorer.test.tsx
│   └── [component files]
├── services/
│   ├── __tests__/
│   │   ├── blockchainAPI.test.ts
│   │   ├── webSocketService.test.ts
│   │   ├── walletService.test.ts
│   │   └── storageService.test.ts
│   └── [service files]
└── utils/
    ├── __tests__/
    │   └── [utility test files]
    └── [utility files]
```

## Component Tests

### Wallet Component

Test cases:
1. Renders wallet creation interface when no address is provided
2. Renders wallet info when address is provided
3. Calls onWalletCreate when Generate Wallet button is clicked
4. Opens transaction form when Send Transaction button is clicked
5. Submits transaction form with valid data
6. Shows error for invalid amount

### TransactionHistory Component

Test cases:
1. Renders empty state when no transactions are provided
2. Displays transactions in reverse chronological order
3. Shows transaction details correctly

### AddressBook Component

Test cases:
1. Renders empty state when no addresses are provided
2. Displays addresses correctly
3. Calls onAddressSelect when an address is clicked

### QRCode Component

Test cases:
1. Renders QR code with correct data
2. Handles different sizes correctly

### ConnectionStatus Component

Test cases:
1. Shows connected status when isConnected is true
2. Shows disconnected status when isConnected is false

### BlockList Component

Test cases:
1. Renders loading state initially
2. Displays blocks correctly
3. Calls onBlockSelect when a block is clicked

### BlockDetail Component

Test cases:
1. Displays block information correctly
2. Shows transaction list
3. Calls onBack when back button is clicked

### TransactionDetail Component

Test cases:
1. Displays transaction information correctly
2. Shows inputs and outputs correctly
3. Calls onBack when back button is clicked

### NetworkStats Component

Test cases:
1. Renders loading state initially
2. Displays blockchain info correctly
3. Calls onSearch when search form is submitted

### BlockchainExplorer Component

Test cases:
1. Renders all explorer components
2. Handles navigation between views
3. Manages state correctly

## Service Tests

### blockchainAPI Service

Test cases:
1. getBlockchainInfo returns correct data
2. getBlock returns correct data for valid index
3. getTransaction returns correct data for valid txid
4. getBalance returns correct data for valid address
5. getPeers returns correct data
6. sendTransaction sends data correctly

### webSocketService Service

Test cases:
1. Connects to WebSocket server
2. Handles message events correctly
3. Reconnects on connection loss
4. Subscribes and unsubscribes to events
5. Sends messages to server

### walletService Service

Test cases:
1. Creates new wallet correctly
2. Loads wallet from private key
3. Gets balance for address
4. Sends transactions
5. Signs and verifies messages

### storageService Service

Test cases:
1. Sets and gets items correctly
2. Removes items correctly
3. Clears all items with prefix
4. Handles item expiration

## Integration Tests

Integration tests should verify that components work together correctly:

1. Wallet creation flow
2. Transaction sending flow
3. Blockchain explorer navigation
4. Real-time updates with WebSocket

## End-to-End Tests

E2E tests should simulate real user interactions:

1. User creates wallet
2. User sends transaction
3. User explores blockchain
4. User views transaction history

## Test Coverage

Aim for the following coverage targets:

- Component tests: 80%+
- Service tests: 90%+
- Integration tests: 70%+
- E2E tests: 60%+

## Continuous Integration

Tests are run automatically on every push to the repository through GitHub Actions.