import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FailToBanComponent } from './fail-to-ban.component';

describe('FailToBanComponent', () => {
  let component: FailToBanComponent;
  let fixture: ComponentFixture<FailToBanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FailToBanComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FailToBanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
