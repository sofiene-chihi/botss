import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment';
import { Observable } from 'rxjs';

const API_URL = environment.apiUrl;

@Injectable({
  providedIn: 'root',
})
export class ConversationService {
  constructor(private http: HttpClient) {}

  createConversation(): Observable<any> {
    return this.http.get(`${API_URL}/new-conversation`);
  }

  getConversationById(convesrationId: string): Observable<any> {
    return this.http.get(`${API_URL}/conversation/${convesrationId}`);
  }

  sendNewMessage(message: string, conversationId: string) {
    console.log('sending message' + message);
    return this.http.post(`${API_URL}/send-message`, {
      message: message,
      conversationId: conversationId,
    });
  }
}
