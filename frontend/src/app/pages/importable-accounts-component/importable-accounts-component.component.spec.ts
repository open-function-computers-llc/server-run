import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ImportableAccountsComponentComponent } from './importable-accounts-component.component';

describe('ImportableAccountsComponentComponent', () => {
  let component: ImportableAccountsComponentComponent;
  let fixture: ComponentFixture<ImportableAccountsComponentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ImportableAccountsComponentComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ImportableAccountsComponentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
