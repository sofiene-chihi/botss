import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-bot-element',
  templateUrl: './bot-element.component.html',
  styleUrl: './bot-element.component.css',
})
export class BotElementComponent {
  @Input() botName: string = '';
  @Input() botImage: string = '';
}
