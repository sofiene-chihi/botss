import { Component, OnInit } from '@angular/core';
import { ConversationService } from '../../services/conversation.service';
import { LocalstorageService } from '../../services/localstorage.service';
import { GetConversationResponse } from '../../interfaces/getConversationResponse.interface';
import { Message } from '../../interfaces/message.interface';
import { SendMessageResponse } from '../../interfaces/sendMessageResponse.interface';

@Component({
  selector: 'app-conversation',
  templateUrl: './conversation.component.html',
  styleUrl: './conversation.component.css',
})
export class ConversationComponent implements OnInit {
  conversationId: string = '';
  messageInput: string = '';
  conversationMessages: string[] = [];

  constructor(
    private conversationService: ConversationService,
    private localStorageService: LocalstorageService
  ) {}
  ngOnInit(): void {
    this.conversationId = this.localStorageService.getItem('conversationId');

    this.conversationService.getConversationById(this.conversationId).subscribe(
      (response: GetConversationResponse) => {
        console.log(response);
        if (response.conversationContent.length > 0) {
          console.log(response.conversationContent);
          response.conversationContent.map((message: Message) => {
            this.conversationMessages.push(message.content);
          });
        }
      },
      (error) => {
        console.error('Error fetching data:', error);
      }
    );
  }

  handleKeyPress(event: KeyboardEvent) {
    if (event.key === 'Enter' || event.keyCode === 13) {
      this.sendMessage();
    }
  }

  sendMessage() {
    console.log('Sending message', this.messageInput);
    if (this.messageInput.trim().length == 0) return;
    this.conversationMessages.push(this.messageInput);
    let messageToSend = this.messageInput;
    this.messageInput = '';
    this.conversationService
      .sendNewMessage(messageToSend, this.conversationId)
      .subscribe(
        (response: SendMessageResponse) => {
          console.log(response);
          this.conversationMessages.push(response.message);
        },
        (error) => {
          console.error('Error fetching data:', error);
        }
      );
  }
}
