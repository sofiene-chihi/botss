export interface Message {
  role: MessageRole;
  content: string;
}

export enum MessageRole {
  USER = 'user',
  ASSISTANT = 'assistant',
}
