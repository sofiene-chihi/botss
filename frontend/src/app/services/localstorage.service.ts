import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class LocalstorageService {
  constructor() {}

  // Method to set an item in local storage
  setItem(key: string, value: any): void {
    localStorage.setItem(key, JSON.stringify(value));
  }

  // Method to get an item from local storage
  getItem(key: string): any {
    const storedItem = localStorage.getItem(key);
    return storedItem ? JSON.parse(storedItem) : null;
  }

  // Method to remove an item from local storage
  removeItem(key: string): void {
    localStorage.removeItem(key);
  }

  // Method to clear all items from local storage
  clear(): void {
    localStorage.clear();
  }
}
