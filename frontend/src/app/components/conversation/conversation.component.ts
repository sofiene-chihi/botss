import { Component, OnInit } from '@angular/core';
import { ConversationService } from '../../services/conversation.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-conversation',
  templateUrl: './conversation.component.html',
  styleUrl: './conversation.component.css',
})
export class ConversationComponent implements OnInit {
  conversationId: string = '';
  messageInput: string = '';
  conversationMessages: string[] = [];

  constructor(private conversationService: ConversationService) {}
  ngOnInit(): void {
    this.conversationId = '';
    this.conversationMessages = ['salut', 'salut cv'];
  }

  handleKeyPress(event: KeyboardEvent) {
    console.log('sending message');
    if (event.key === 'Enter' || event.keyCode === 13) {
      this.sendMessage();
    }
  }

  sendMessage() {
    console.log('Sending message', this.messageInput);
    this.messageInput = '';
    this.conversationService
      .sendNewMessage(this.messageInput, this.conversationId)
      .subscribe(
        (response) => {
          console.log(response);
        },
        (error) => {
          console.error('Error fetching data:', error);
        }
      );
  }
}
