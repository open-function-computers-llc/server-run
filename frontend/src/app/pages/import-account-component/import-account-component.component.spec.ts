import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ImportAccountComponentComponent } from './import-account-component.component';

describe('ImportAccountComponentComponent', () => {
  let component: ImportAccountComponentComponent;
  let fixture: ComponentFixture<ImportAccountComponentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ImportAccountComponentComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ImportAccountComponentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
