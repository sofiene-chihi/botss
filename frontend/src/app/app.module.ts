import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/home/header/header.component';
import { BotsGridComponent } from './components/home/bots-grid/bots-grid.component';
import { BotElementComponent } from './components/home/bots-grid/bot-element/bot-element.component';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { ConversationComponent } from './components/conversation/conversation.component';
import { HomeComponent } from './components/home/home.component';
import { FormsModule } from '@angular/forms';
import { MessageItemComponent } from './components/conversation/message-item/message-item.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    BotsGridComponent,
    BotElementComponent,
    ConversationComponent,
    HomeComponent,
    MessageItemComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MatGridListModule,
    MatButtonModule,
    MatCardModule,
    HttpClientModule,
    FormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
