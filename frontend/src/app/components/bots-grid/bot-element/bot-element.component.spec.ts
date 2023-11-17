import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BotElementComponent } from './bot-element.component';

describe('BotElementComponent', () => {
  let component: BotElementComponent;
  let fixture: ComponentFixture<BotElementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [BotElementComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(BotElementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
