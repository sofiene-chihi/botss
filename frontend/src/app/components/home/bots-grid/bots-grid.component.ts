import { Component } from '@angular/core';
import { ConversationService } from '../../../services/conversation.service';
import { Router } from '@angular/router';
import { LocalstorageService } from '../../../services/localstorage.service';

@Component({
  selector: 'app-bots-grid',
  templateUrl: './bots-grid.component.html',
  styleUrl: './bots-grid.component.css',
})
export class BotsGridComponent {
  items: any[] = [
    {
      id: 1,
      name: 'Ecommerce Shop Assistant Chatbot',
      image: 'ecommerce-image.webp',
    },
    { id: 2, name: 'Customer Service Chatbot', image: 'customer-service.png' },
    {
      id: 3,
      name: 'Restaurant Orderes Chatbot',
      image: 'restaurant-order.png',
    },
  ];

  constructor(
    private conversationService: ConversationService,
    private router: Router,
    private localStorageService: LocalstorageService
  ) {}

  openNewConversation() {
    this.conversationService.createConversation().subscribe(
      (response) => {
        console.log(response.conversationId);
        this.localStorageService.setItem(
          'conversationId',
          response.conversationId
        );
        this.router.navigate(['/conversation', response.conversationId]);
      },
      (error) => {
        console.error('Error fetching data:', error);
      }
    );
  }
}
