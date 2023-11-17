import { Component } from '@angular/core';

@Component({
  selector: 'app-bots-grid',
  templateUrl: './bots-grid.component.html',
  styleUrl: './bots-grid.component.css',
})
export class BotsGridComponent {
  items: any[] = [
    { id: 1, name: 'bot 1' },
    { id: 2, name: 'bot 2' },
    { id: 3, name: 'bot 3' },
    { id: 4, name: 'bot 1' },
    { id: 5, name: 'bot 2' },
    { id: 6, name: 'bot 3' },
  ];
}
