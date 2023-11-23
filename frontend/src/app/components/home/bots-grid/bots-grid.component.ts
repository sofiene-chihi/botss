import { Component, OnInit } from '@angular/core';
import { ConversationService } from '../../../services/conversation.service';
import { Router } from '@angular/router';
import { LocalstorageService } from '../../../services/localstorage.service';
import { CreateConversationResponse } from '../../../interfaces/createConversationResponse.interface';

@Component({
  selector: 'app-bots-grid',
  templateUrl: './bots-grid.component.html',
  styleUrl: './bots-grid.component.css',
})
export class BotsGridComponent implements OnInit {
  isPhone: boolean = false;
  gridColumns: number = 3
  rowHeight1: number = 3
  rowHeight2: number = 2

  items: any[] = [
    {
      id: 1,
      name: 'Ecommerce Shop Chatbot',
      image: 'ecommerce-image.webp',
    },
    { id: 2, name: 'Customer Service Chatbot', image: 'customer-service.png' },
    {
      id: 3,
      name: 'Pizzeria Chatbot',
      image: 'restaurant-order.png',
    },
  ];

  constructor(
    private conversationService: ConversationService,
    private router: Router,
    private localStorageService: LocalstorageService,
  ) {}

  ngOnInit(): void {
    this.checkIfPhone();
    if (this.isPhone) {
      this.gridColumns = 1
      this.rowHeight1 = 1.7
      this.rowHeight2 = 1
    }
  }

  private checkIfPhone(): void {
    const windowWidth = window.innerWidth

    const phoneWidthThreshold = 768; 
    this.isPhone = windowWidth < phoneWidthThreshold;
  }

  openNewConversation() {
    this.conversationService.createConversation().subscribe(
      (response:CreateConversationResponse) => {
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
