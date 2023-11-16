import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { BotsGridComponent } from './components/bots-grid/bots-grid.component';
import { BotItemComponent } from './components/bots-grid/bot-item/bot-item.component';
import { MatGridListModule } from '@angular/material/grid-list';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    BotsGridComponent,
    BotItemComponent,
  ],
  imports: [BrowserModule, AppRoutingModule, MatGridListModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
