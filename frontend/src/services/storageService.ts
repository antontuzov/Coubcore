// Storage service for handling localStorage operations
class StorageService {
  private prefix: string;

  constructor(prefix: string = 'coubcore_') {
    this.prefix = prefix;
  }

  // Set item in localStorage
  setItem(key: string, value: any): void {
    try {
      const fullKey = this.prefix + key;
      const stringValue = typeof value === 'string' ? value : JSON.stringify(value);
      localStorage.setItem(fullKey, stringValue);
    } catch (error) {
      console.error(`Error setting item ${key} in localStorage:`, error);
    }
  }

  // Get item from localStorage
  getItem<T>(key: string, defaultValue: T | null = null): T | null {
    try {
      const fullKey = this.prefix + key;
      const item = localStorage.getItem(fullKey);
      
      if (item === null) {
        return defaultValue;
      }
      
      try {
        return JSON.parse(item);
      } catch {
        return item as unknown as T;
      }
    } catch (error) {
      console.error(`Error getting item ${key} from localStorage:`, error);
      return defaultValue;
    }
  }

  // Remove item from localStorage
  removeItem(key: string): void {
    try {
      const fullKey = this.prefix + key;
      localStorage.removeItem(fullKey);
    } catch (error) {
      console.error(`Error removing item ${key} from localStorage:`, error);
    }
  }

  // Clear all items with the prefix
  clear(): void {
    try {
      const keysToRemove: string[] = [];
      
      for (let i = 0; i < localStorage.length; i++) {
        const key = localStorage.key(i);
        if (key && key.startsWith(this.prefix)) {
          keysToRemove.push(key);
        }
      }
      
      keysToRemove.forEach(key => {
        localStorage.removeItem(key);
      });
    } catch (error) {
      console.error('Error clearing localStorage:', error);
    }
  }

  // Get all items with the prefix
  getAllItems(): { [key: string]: any } {
    try {
      const items: { [key: string]: any } = {};
      
      for (let i = 0; i < localStorage.length; i++) {
        const key = localStorage.key(i);
        if (key && key.startsWith(this.prefix)) {
          const shortKey = key.substring(this.prefix.length);
          items[shortKey] = this.getItem(shortKey);
        }
      }
      
      return items;
    } catch (error) {
      console.error('Error getting all items from localStorage:', error);
      return {};
    }
  }

  // Check if item exists
  hasItem(key: string): boolean {
    try {
      const fullKey = this.prefix + key;
      return localStorage.getItem(fullKey) !== null;
    } catch (error) {
      console.error(`Error checking item ${key} in localStorage:`, error);
      return false;
    }
  }

  // Set item with expiration
  setItemWithExpiration(key: string, value: any, expirationMinutes: number): void {
    try {
      const expirationTime = Date.now() + expirationMinutes * 60 * 1000;
      const item = {
        value,
        expiration: expirationTime
      };
      this.setItem(key, item);
    } catch (error) {
      console.error(`Error setting item with expiration ${key} in localStorage:`, error);
    }
  }

  // Get item with expiration check
  getItemWithExpiration<T>(key: string, defaultValue: T | null = null): T | null {
    try {
      const item = this.getItem<{ value: T; expiration: number }>(key, null);
      
      if (item === null) {
        return defaultValue;
      }
      
      if (Date.now() > item.expiration) {
        this.removeItem(key);
        return defaultValue;
      }
      
      return item.value;
    } catch (error) {
      console.error(`Error getting item with expiration ${key} from localStorage:`, error);
      return defaultValue;
    }
  }
}

export default StorageService;