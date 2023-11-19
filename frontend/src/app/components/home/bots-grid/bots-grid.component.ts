import { Component } from '@angular/core';
import { ConversationService } from '../../../services/conversation.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-bots-grid',
  templateUrl: './bots-grid.component.html',
  styleUrl: './bots-grid.component.css',
})
export class BotsGridComponent {


  items: any[] = [
    { id: 1, name: 'Ecommerce Chatbot', image: 'ecommerce-image.webp' },
    { id: 2, name: 'xxxxxxx Chatbot', image: 'coming-soon.png' },
    { id: 3, name: 'xxxxxxx Chatbot', image: 'coming-soon.png' },
    { id: 4, name: 'xxxxxxx Chatbot', image: 'coming-soon.png' },
    { id: 5, name: 'xxxxxxx Chatbot', image: 'coming-soon.png' },
    { id: 6, name: 'xxxxxxx Chatbot', image: 'coming-soon.png' },
  ];


  constructor(private conversationService: ConversationService,private router: Router) {}


  openNewConversation(){
    this.conversationService.createConversation().subscribe(
      (response) => {
        console.log(response.conversationId);
        this.router.navigate(['/conversation', response.conversationId]);      },
      (error) => {
        console.error('Error fetching data:', error);
      }
    );
  
  }
}
