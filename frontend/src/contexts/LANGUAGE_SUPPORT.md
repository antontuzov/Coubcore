# Language Support Implementation Guide

## Overview

This document describes how to implement internationalization (i18n) support in the Coubcore frontend application.

## Language Context

The language context provides a way to manage the application's language and translations across all components.

### Supported Languages

1. **English (en)**: Default language
2. **Spanish (es)**: Spanish translations
3. **French (fr)**: French translations
4. **German (de)**: German translations
5. **Chinese (zh)**: Chinese translations

## Implementation

### Language Context Provider

The `LanguageProvider` component wraps the entire application and provides language management functionality.

### Translation System

Translations are stored in JSON files for each language:

```
src/
├── locales/
│   ├── en.json
│   ├── es.json
│   ├── fr.json
│   ├── de.json
│   └── zh.json
```

Example translation file (en.json):
```json
{
  "wallet": {
    "title": "Wallet",
    "create": "Create Wallet",
    "balance": "Balance",
    "send": "Send Transaction"
  },
  "explorer": {
    "title": "Blockchain Explorer",
    "search": "Search",
    "blocks": "Blocks",
    "transactions": "Transactions"
  }
}
```

### Using the Language Context

Components can access the language context using the `useLanguage` hook:

```tsx
import { useLanguage } from '../contexts/LanguageContext';

const MyComponent = () => {
  const { language, setLanguage, t } = useLanguage();
  
  return (
    <div>
      <p>{t('wallet.title')}</p>
      <select value={language} onChange={(e) => setLanguage(e.target.value as any)}>
        <option value="en">English</option>
        <option value="es">Español</option>
        <option value="fr">Français</option>
        <option value="de">Deutsch</option>
        <option value="zh">中文</option>
      </select>
    </div>
  );
};
```

## Language Switcher Component

A language switcher component would allow users to change the application language:

```tsx
import { useLanguage } from '../contexts/LanguageContext';

const LanguageSwitcher = () => {
  const { language, setLanguage } = useLanguage();
  
  const languages = [
    { code: 'en', name: 'English' },
    { code: 'es', name: 'Español' },
    { code: 'fr', name: 'Français' },
    { code: 'de', name: 'Deutsch' },
    { code: 'zh', name: '中文' },
  ];
  
  return (
    <select 
      value={language} 
      onChange={(e) => setLanguage(e.target.value as any)}
      className="bg-gray-700 text-white rounded px-2 py-1"
    >
      {languages.map((lang) => (
        <option key={lang.code} value={lang.code}>
          {lang.name}
        </option>
      ))}
    </select>
  );
};
```

## Translation Function

The translation function (`t`) looks up translations based on keys:

```ts
const t = (key: string): string => {
  // Split the key by dots to navigate nested objects
  const keys = key.split('.');
  
  // Start with the current language translations
  let translation: any = translations[language];
  
  // Navigate through the nested objects
  for (const k of keys) {
    if (!translation || !translation[k]) {
      // Return the key if translation is not found
      return key;
    }
    translation = translation[k];
  }
  
  return translation;
};
```

## Pluralization Support

Advanced translation systems support pluralization:

```ts
const t = (key: string, count?: number): string => {
  // Implementation would handle pluralization based on count
  // and language-specific pluralization rules
};
```

## Date and Number Formatting

Different languages have different formatting rules for dates and numbers:

```ts
const formatDate = (date: Date): string => {
  return new Intl.DateTimeFormat(language).format(date);
};

const formatNumber = (number: number): string => {
  return new Intl.NumberFormat(language).format(number);
};
```

## Right-to-Left (RTL) Support

Some languages like Arabic and Hebrew are written right-to-left:

```css
:root[lang="ar"] {
  direction: rtl;
}

:root[lang="he"] {
  direction: rtl;
}
```

## Persistence

Language preferences are stored in localStorage to persist between sessions:

```ts
// Save language preference
localStorage.setItem('language', language);

// Load language preference
const savedLanguage = localStorage.getItem('language') as Language | null;
```

## Dynamic Imports

To reduce bundle size, translations can be loaded dynamically:

```ts
const loadTranslations = async (language: Language) => {
  const translations = await import(`../locales/${language}.json`);
  return translations.default;
};
```

## Future Enhancements

1. **Automatic Language Detection**: Detect user's preferred language
2. **Translation Management**: Tools for managing translations
3. **Context-Aware Translations**: Different translations based on context
4. **Voice Support**: Text-to-speech for accessibility
5. **Cultural Adaptation**: Adapting content to cultural norms