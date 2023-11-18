import { Component, OnInit } from '@angular/core';
import { ConversationService } from '../../services/conversation.service';

@Component({
  selector: 'app-conversation',
  templateUrl: './conversation.component.html',
  styleUrl: './conversation.component.css',
})
export class ConversationComponent implements OnInit {
  constructor(private conversationService: ConversationService) {}

  ngOnInit(): void {
    this.conversationService.createConversation().subscribe(
      (response) => {
        console.log(response); // Handle the fetched data here
      },
      (error) => {
        console.error('Error fetching data:', error);
      }
    );
  }

  handleKeyPress(event: KeyboardEvent) {
    // Handle key press logic here
    // You can access event properties like event.key, event.keyCode, etc.
    // For example:
    if (event.key === 'Enter') {
      this.sendMessage();
    }
  }

  sendMessage() {
    console.log('Sending message');
    // Implement logic to send the message
    // This function will be called when the "Send" button is clicked
  }
}
