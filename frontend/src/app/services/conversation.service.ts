import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';
import { SendMessageResponse } from '../interfaces/sendMessageResponse.interface';
import { GetConversationResponse } from '../interfaces/getConversationResponse.interface';
import { CreateConversationResponse } from '../interfaces/createConversationResponse.interface';

const API_URL = environment.apiUrl;

@Injectable({
  providedIn: 'root',
})
export class ConversationService {
  constructor(private http: HttpClient) {}

  createConversation(): Observable<CreateConversationResponse> {
    return this.http.get<CreateConversationResponse>(`${API_URL}/new-conversation`);
  }

  getConversationById(convesrationId: string): Observable<GetConversationResponse> {
    return this.http.get<GetConversationResponse>(`${API_URL}/conversation/${convesrationId}`);
  }

  sendNewMessage(message: string, conversationId: string): Observable<SendMessageResponse> {
    console.log('sending message' + message);
    return this.http.post<SendMessageResponse>(`${API_URL}/send-message`, {
      message: message,
      conversationId: conversationId,
    });
  }
}
