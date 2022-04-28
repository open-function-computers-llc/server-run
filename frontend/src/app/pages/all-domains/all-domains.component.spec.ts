import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllDomainsComponent } from './all-domains.component';

describe('AllDomainsComponent', () => {
  let component: AllDomainsComponent;
  let fixture: ComponentFixture<AllDomainsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AllDomainsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AllDomainsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
