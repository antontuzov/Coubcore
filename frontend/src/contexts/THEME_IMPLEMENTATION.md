# Theme Implementation Guide

## Overview

This document describes how to implement theme support in the Coubcore frontend application.

## Theme Context

The theme context provides a way to manage the application's theme and color scheme across all components.

### Theme Types

1. **Light Theme**: A bright, light-colored theme
2. **Dark Theme**: A dark, low-light theme (default)
3. **System Theme**: Follows the user's system preference

### Color Schemes

1. **Blue**: Primary blue color scheme
2. **Green**: Green color scheme
3. **Purple**: Purple color scheme
4. **Red**: Red color scheme

## Implementation

### Theme Context Provider

The `ThemeProvider` component wraps the entire application and provides theme management functionality.

### Using the Theme Context

Components can access the theme context using the `useTheme` hook:

```tsx
import { useTheme } from '../contexts/ThemeContext';

const MyComponent = () => {
  const { theme, colorScheme, toggleTheme, setTheme, setColorScheme } = useTheme();
  
  return (
    <div>
      <p>Current theme: {theme}</p>
      <p>Current color scheme: {colorScheme}</p>
      <button onClick={toggleTheme}>Toggle Theme</button>
    </div>
  );
};
```

## CSS Implementation

### CSS Variables

The theme is implemented using CSS variables that are updated when the theme changes:

```css
:root {
  /* Light theme variables */
  --bg-primary: #ffffff;
  --bg-secondary: #f5f5f5;
  --text-primary: #333333;
  --text-secondary: #666666;
}

:root.dark {
  /* Dark theme variables */
  --bg-primary: #1a1a1a;
  --bg-secondary: #2a2a2a;
  --text-primary: #ffffff;
  --text-secondary: #cccccc;
}

:root.blue {
  /* Blue color scheme */
  --accent: #3b82f6;
}

:root.green {
  /* Green color scheme */
  --accent: #10b981;
}

:root.purple {
  /* Purple color scheme */
  --accent: #8b5cf6;
}

:root.red {
  /* Red color scheme */
  --accent: #ef4444;
}
```

## Theme Toggle Component

A theme toggle component would allow users to switch between themes:

```tsx
import { useTheme } from '../contexts/ThemeContext';
import { SunIcon, MoonIcon } from '@heroicons/react/24/outline';

const ThemeToggle = () => {
  const { theme, toggleTheme } = useTheme();
  
  return (
    <button
      onClick={toggleTheme}
      className="p-2 rounded-full bg-gray-700 text-white hover:bg-gray-600 transition-colors"
      aria-label="Toggle theme"
    >
      {theme === 'dark' ? (
        <SunIcon className="h-5 w-5" />
      ) : (
        <MoonIcon className="h-5 w-5" />
      )}
    </button>
  );
};
```

## Color Scheme Selector

A color scheme selector would allow users to choose their preferred color scheme:

```tsx
import { useTheme } from '../contexts/ThemeContext';

const ColorSchemeSelector = () => {
  const { colorScheme, setColorScheme } = useTheme();
  
  const schemes = [
    { id: 'blue', name: 'Blue' },
    { id: 'green', name: 'Green' },
    { id: 'purple', name: 'Purple' },
    { id: 'red', name: 'Red' },
  ];
  
  return (
    <div className="flex space-x-2">
      {schemes.map((scheme) => (
        <button
          key={scheme.id}
          onClick={() => setColorScheme(scheme.id as any)}
          className={`px-3 py-1 rounded ${
            colorScheme === scheme.id
              ? 'bg-accent text-white'
              : 'bg-gray-700 text-gray-300 hover:bg-gray-600'
          }`}
        >
          {scheme.name}
        </button>
      ))}
    </div>
  );
};
```

## Integration with Tailwind CSS

To integrate with Tailwind CSS, we can use the `darkMode` configuration:

```js
// tailwind.config.js
module.exports = {
  darkMode: 'class', // or 'media' for system preference
  // ... other configuration
};
```

## Persistence

Theme preferences are stored in localStorage to persist between sessions:

```ts
// Save theme preference
localStorage.setItem('theme', theme);

// Load theme preference
const savedTheme = localStorage.getItem('theme') as Theme | null;
```

## System Preference Detection

The application can detect the user's system preference for dark mode:

```ts
const systemTheme = window.matchMedia('(prefers-color-scheme: dark)').matches 
  ? 'dark' 
  : 'light';
```

## Future Enhancements

1. **Custom Themes**: Allow users to create and save custom themes
2. **Theme Editor**: Provide a visual theme editor
3. **Accessibility**: Ensure proper contrast ratios for accessibility
4. **Animations**: Add smooth transitions when switching themes
5. **Theme Sharing**: Allow users to share their custom themes