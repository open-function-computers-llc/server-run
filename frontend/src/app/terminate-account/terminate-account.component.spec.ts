import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TerminateAccountComponent } from './terminate-account.component';

describe('TerminateAccountComponent', () => {
  let component: TerminateAccountComponent;
  let fixture: ComponentFixture<TerminateAccountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TerminateAccountComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TerminateAccountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
