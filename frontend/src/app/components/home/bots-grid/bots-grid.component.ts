import { Component } from '@angular/core';

@Component({
  selector: 'app-bots-grid',
  templateUrl: './bots-grid.component.html',
  styleUrl: './bots-grid.component.css',
})
export class BotsGridComponent {
  items: any[] = [
    { id: 1, name: 'Ecommerce Chatbot', image: 'ecommerce-image.webp' },
    { id: 2, name: 'xxxxxxx Chatbot', image: 'ecommep' },
    { id: 3, name: 'xxxxxxx Chatbot', image: 'ecommep' },
    { id: 4, name: 'xxxxxxx Chatbot', image: 'ecommewebp' },
    { id: 5, name: 'xxxxxxx Chatbot', image: 'ecommerce-p' },
    { id: 6, name: 'xxxxxxx Chatbot', image: 'ecom.webp' },
  ];
}
