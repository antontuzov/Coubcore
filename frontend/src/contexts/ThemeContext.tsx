import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';

// Define theme types
export type Theme = 'light' | 'dark' | 'system';
export type ColorScheme = 'blue' | 'green' | 'purple' | 'red';

// Define the theme context type
interface ThemeContextType {
  theme: Theme;
  colorScheme: ColorScheme;
  toggleTheme: () => void;
  setTheme: (theme: Theme) => void;
  setColorScheme: (colorScheme: ColorScheme) => void;
}

// Create the theme context
const ThemeContext = createContext<ThemeContextType | undefined>(undefined);

// Theme provider component
export const ThemeProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [theme, setTheme] = useState<Theme>(() => {
    const savedTheme = localStorage.getItem('theme') as Theme | null;
    return savedTheme || 'system';
  });
  
  const [colorScheme, setColorScheme] = useState<ColorScheme>(() => {
    const savedColorScheme = localStorage.getItem('colorScheme') as ColorScheme | null;
    return savedColorScheme || 'blue';
  });

  // Apply theme to document
  useEffect(() => {
    const root = document.documentElement;
    
    // Remove existing theme classes
    root.classList.remove('light', 'dark');
    
    // Apply theme based on user preference or system preference
    if (theme === 'system') {
      const systemTheme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
      root.classList.add(systemTheme);
    } else {
      root.classList.add(theme);
    }
    
    // Save theme preference
    localStorage.setItem('theme', theme);
  }, [theme]);

  // Apply color scheme to document
  useEffect(() => {
    const root = document.documentElement;
    
    // Remove existing color scheme classes
    root.classList.remove('blue', 'green', 'purple', 'red');
    
    // Apply color scheme
    root.classList.add(colorScheme);
    
    // Save color scheme preference
    localStorage.setItem('colorScheme', colorScheme);
  }, [colorScheme]);

  // Toggle between light and dark theme
  const toggleTheme = () => {
    setTheme(prevTheme => {
      if (prevTheme === 'light') return 'dark';
      if (prevTheme === 'dark') return 'light';
      return 'light';
    });
  };

  // Provide the context value
  const contextValue: ThemeContextType = {
    theme,
    colorScheme,
    toggleTheme,
    setTheme,
    setColorScheme
  };

  return (
    <ThemeContext.Provider value={contextValue}>
      {children}
    </ThemeContext.Provider>
  );
};

// Custom hook to use the theme context
export const useTheme = (): ThemeContextType => {
  const context = useContext(ThemeContext);
  
  if (context === undefined) {
    throw new Error('useTheme must be used within a ThemeProvider');
  }
  
  return context;
};