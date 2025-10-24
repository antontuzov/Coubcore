# Export Functionality Implementation Guide

## Overview

This document describes how to implement export functionality in the Coubcore frontend application, allowing users to export their data in various formats.

## Supported Export Formats

1. **JSON**: Raw data export in JSON format
2. **CSV**: Comma-separated values for spreadsheet applications
3. **PDF**: Printable document format
4. **TXT**: Plain text format

## Implementation

### JSON Export

JSON export is the simplest form of data export:

```ts
// utils/export.ts
export const exportToJSON = (data: any, filename: string) => {
  const json = JSON.stringify(data, null, 2);
  const blob = new Blob([json], { type: 'application/json' });
  const url = URL.createObjectURL(blob);
  
  const a = document.createElement('a');
  a.href = url;
  a.download = `${filename}.json`;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(url);
};
```

### CSV Export

CSV export requires formatting data as comma-separated values:

```ts
export const exportToCSV = (data: any[], filename: string) => {
  if (data.length === 0) return;
  
  // Create header row
  const headers = Object.keys(data[0]).join(',');
  
  // Create data rows
  const rows = data.map(obj => 
    Object.values(obj).map(value => 
      `"${String(value).replace(/"/g, '""')}"`
    ).join(',')
  );
  
  // Combine header and rows
  const csv = [headers, ...rows].join('\n');
  
  const blob = new Blob([csv], { type: 'text/csv' });
  const url = URL.createObjectURL(blob);
  
  const a = document.createElement('a');
  a.href = url;
  a.download = `${filename}.csv`;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(url);
};
```

### PDF Export

PDF export requires a library like jsPDF:

```ts
// First install jsPDF: npm install jspdf
import jsPDF from 'jspdf';
import 'jspdf-autotable';

export const exportToPDF = (data: any[], filename: string, title: string) => {
  const doc = new jsPDF();
  
  // Add title
  doc.setFontSize(18);
  doc.text(title, 14, 22);
  
  // Add table
  (doc as any).autoTable({
    head: [Object.keys(data[0] || {})],
    body: data.map(obj => Object.values(obj)),
    startY: 30,
  });
  
  doc.save(`${filename}.pdf`);
};
```

### TXT Export

TXT export creates a simple text file:

```ts
export const exportToTXT = (data: any, filename: string) => {
  const text = typeof data === 'string' ? data : JSON.stringify(data, null, 2);
  const blob = new Blob([text], { type: 'text/plain' });
  const url = URL.createObjectURL(blob);
  
  const a = document.createElement('a');
  a.href = url;
  a.download = `${filename}.txt`;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(url);
};
```

## Export Components

### Export Button

A reusable export button component:

```tsx
import React from 'react';
import { exportToJSON, exportToCSV, exportToPDF, exportToTXT } from '../utils/export';

interface ExportButtonProps {
  data: any;
  filename: string;
  title?: string;
}

const ExportButton: React.FC<ExportButtonProps> = ({ data, filename, title }) => {
  const handleExport = (format: 'json' | 'csv' | 'pdf' | 'txt') => {
    switch (format) {
      case 'json':
        exportToJSON(data, filename);
        break;
      case 'csv':
        if (Array.isArray(data)) {
          exportToCSV(data, filename);
        }
        break;
      case 'pdf':
        if (Array.isArray(data)) {
          exportToPDF(data, filename, title || filename);
        }
        break;
      case 'txt':
        exportToTXT(data, filename);
        break;
    }
  };

  return (
    <div className="export-buttons">
      <button onClick={() => handleExport('json')} className="btn btn-secondary">
        Export JSON
      </button>
      <button onClick={() => handleExport('csv')} className="btn btn-secondary">
        Export CSV
      </button>
      <button onClick={() => handleExport('pdf')} className="btn btn-secondary">
        Export PDF
      </button>
      <button onClick={() => handleExport('txt')} className="btn btn-secondary">
        Export TXT
      </button>
    </div>
  );
};

export default ExportButton;
```

### Export Modal

A modal for selecting export options:

```tsx
import React, { useState } from 'react';

interface ExportModalProps {
  data: any;
  filename: string;
  title?: string;
  onClose: () => void;
  onExport: (format: string) => void;
}

const ExportModal: React.FC<ExportModalProps> = ({ data, filename, title, onClose, onExport }) => {
  const [selectedFormat, setSelectedFormat] = useState('json');
  
  const formats = [
    { id: 'json', name: 'JSON', description: 'Raw data format' },
    { id: 'csv', name: 'CSV', description: 'Comma-separated values' },
    { id: 'pdf', name: 'PDF', description: 'Portable document format' },
    { id: 'txt', name: 'TXT', description: 'Plain text format' },
  ];

  const handleExport = () => {
    onExport(selectedFormat);
    onClose();
  };

  return (
    <div className="modal">
      <div className="modal-content">
        <h2>Export Data</h2>
        <p>Choose a format to export your data:</p>
        
        <div className="format-options">
          {formats.map(format => (
            <div 
              key={format.id}
              className={`format-option ${selectedFormat === format.id ? 'selected' : ''}`}
              onClick={() => setSelectedFormat(format.id)}
            >
              <h3>{format.name}</h3>
              <p>{format.description}</p>
            </div>
          ))}
        </div>
        
        <div className="modal-actions">
          <button onClick={onClose} className="btn btn-secondary">
            Cancel
          </button>
          <button onClick={handleExport} className="btn btn-primary">
            Export
          </button>
        </div>
      </div>
    </div>
  );
};
```

## Integration with Components

### Wallet Export

Export wallet transaction history:

```tsx
const Wallet: React.FC<WalletProps> = ({ transactions }) => {
  const exportTransactions = () => {
    exportToCSV(transactions, 'wallet-transactions');
  };

  return (
    <div>
      {/* Wallet content */}
      <button onClick={exportTransactions}>
        Export Transaction History
      </button>
    </div>
  );
};
```

### Blockchain Explorer Export

Export blockchain data:

```tsx
const BlockchainExplorer: React.FC = () => {
  const [blocks, setBlocks] = useState<Block[]>([]);
  
  const exportBlocks = () => {
    exportToJSON(blocks, 'blockchain-data');
  };

  return (
    <div>
      {/* Explorer content */}
      <button onClick={exportBlocks}>
        Export Blockchain Data
      </button>
    </div>
  );
};
```

## Security Considerations

1. **Data Validation**: Validate exported data to prevent injection attacks
2. **File Size Limits**: Implement limits on export file sizes
3. **Rate Limiting**: Prevent excessive export requests
4. **Privacy**: Ensure sensitive data is not exported without consent

## Performance Optimization

1. **Chunked Export**: For large datasets, export in chunks
2. **Web Workers**: Use web workers for heavy export processing
3. **Progress Indicators**: Show progress for long-running exports
4. **Compression**: Compress large exports to reduce file size

## Future Enhancements

1. **Scheduled Exports**: Allow users to schedule regular exports
2. **Export Templates**: Customizable export templates
3. **Cloud Storage**: Export directly to cloud storage services
4. **API Integration**: Export to external services
5. **Incremental Exports**: Export only changed data since last export