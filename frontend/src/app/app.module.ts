import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { BotsGridComponent } from './components/bots-grid/bots-grid.component';
import { BotElementComponent } from './components/bots-grid/bot-element/bot-element.component';
import { MatGridListModule } from '@angular/material/grid-list';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    BotsGridComponent,
    BotElementComponent,
  ],
  imports: [BrowserModule, AppRoutingModule, MatGridListModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
