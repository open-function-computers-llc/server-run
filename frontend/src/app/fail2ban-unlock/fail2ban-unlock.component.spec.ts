import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Fail2banUnlockComponent } from './fail2ban-unlock.component';

describe('Fail2banUnlockComponent', () => {
  let component: Fail2banUnlockComponent;
  let fixture: ComponentFixture<Fail2banUnlockComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ Fail2banUnlockComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(Fail2banUnlockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
