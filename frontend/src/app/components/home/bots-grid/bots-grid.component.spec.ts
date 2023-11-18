import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BotsGridComponent } from './bots-grid.component';

describe('BotsGridComponent', () => {
  let component: BotsGridComponent;
  let fixture: ComponentFixture<BotsGridComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [BotsGridComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(BotsGridComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
