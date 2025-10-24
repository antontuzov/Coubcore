import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom';
import Wallet from '../Wallet';

// Mock the QRCode component since it uses canvas
jest.mock('../QRCode', () => {
  return function MockQRCode() {
    return <div data-testid="qr-code">QR Code</div>;
  };
});

describe('Wallet Component', () => {
  const mockOnWalletCreate = jest.fn();
  const mockOnTransactionSend = jest.fn();

  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('renders wallet creation interface when no address is provided', () => {
    render(
      <Wallet
        address=""
        balance={0}
        onWalletCreate={mockOnWalletCreate}
        onTransactionSend={mockOnTransactionSend}
      />
    );

    expect(screen.getByText('Create New Wallet')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: 'Generate Wallet' })).toBeInTheDocument();
  });

  it('renders wallet info when address is provided', () => {
    const address = '0x1234567890abcdef1234567890abcdef12345678';
    const balance = 100.5;

    render(
      <Wallet
        address={address}
        balance={balance}
        onWalletCreate={mockOnWalletCreate}
        onTransactionSend={mockOnTransactionSend}
      />
    );

    expect(screen.getByText('Wallet Address')).toBeInTheDocument();
    expect(screen.getByText(address)).toBeInTheDocument();
    expect(screen.getByText(`${balance} COUB`)).toBeInTheDocument();
    expect(screen.getByRole('button', { name: 'Send Transaction' })).toBeInTheDocument();
  });

  it('calls onWalletCreate when Generate Wallet button is clicked', () => {
    render(
      <Wallet
        address=""
        balance={0}
        onWalletCreate={mockOnWalletCreate}
        onTransactionSend={mockOnTransactionSend}
      />
    );

    const generateButton = screen.getByRole('button', { name: 'Generate Wallet' });
    fireEvent.click(generateButton);

    expect(mockOnWalletCreate).toHaveBeenCalledTimes(1);
  });

  it('opens transaction form when Send Transaction button is clicked', () => {
    const address = '0x1234567890abcdef1234567890abcdef12345678';

    render(
      <Wallet
        address={address}
        balance={100.5}
        onWalletCreate={mockOnWalletCreate}
        onTransactionSend={mockOnTransactionSend}
      />
    );

    const sendButton = screen.getByRole('button', { name: 'Send Transaction' });
    fireEvent.click(sendButton);

    expect(screen.getByPlaceholderText('Recipient Address')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Amount')).toBeInTheDocument();
  });

  it('submits transaction form with valid data', () => {
    const address = '0x1234567890abcdef1234567890abcdef12345678';
    const recipient = '0xabcdef1234567890abcdef1234567890abcdef12';
    const amount = '50';

    render(
      <Wallet
        address={address}
        balance={100.5}
        onWalletCreate={mockOnWalletCreate}
        onTransactionSend={mockOnTransactionSend}
      />
    );

    // Open transaction form
    const sendButton = screen.getByRole('button', { name: 'Send Transaction' });
    fireEvent.click(sendButton);

    // Fill form
    const recipientInput = screen.getByPlaceholderText('Recipient Address');
    const amountInput = screen.getByPlaceholderText('Amount');
    const submitButton = screen.getByRole('button', { name: 'Send' });

    fireEvent.change(recipientInput, { target: { value: recipient } });
    fireEvent.change(amountInput, { target: { value: amount } });
    fireEvent.click(submitButton);

    expect(mockOnTransactionSend).toHaveBeenCalledWith(recipient, parseFloat(amount));
  });

  it('shows error for invalid amount', () => {
    const address = '0x1234567890abcdef1234567890abcdef12345678';
    const recipient = '0xabcdef1234567890abcdef1234567890abcdef12';
    const amount = 'invalid';

    render(
      <Wallet
        address={address}
        balance={100.5}
        onWalletCreate={mockOnWalletCreate}
        onTransactionSend={mockOnTransactionSend}
      />
    );

    // Open transaction form
    const sendButton = screen.getByRole('button', { name: 'Send Transaction' });
    fireEvent.click(sendButton);

    // Fill form with invalid amount
    const recipientInput = screen.getByPlaceholderText('Recipient Address');
    const amountInput = screen.getByPlaceholderText('Amount');
    const submitButton = screen.getByRole('button', { name: 'Send' });

    fireEvent.change(recipientInput, { target: { value: recipient } });
    fireEvent.change(amountInput, { target: { value: amount } });
    fireEvent.click(submitButton);

    expect(screen.getByText('Please enter a valid amount')).toBeInTheDocument();
    expect(mockOnTransactionSend).not.toHaveBeenCalled();
  });
});