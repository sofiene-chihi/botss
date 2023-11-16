import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BotItemComponent } from './bot-item.component';

describe('BotItemComponent', () => {
  let component: BotItemComponent;
  let fixture: ComponentFixture<BotItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [BotItemComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(BotItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
