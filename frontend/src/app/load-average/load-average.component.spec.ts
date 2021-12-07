import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoadAverageComponent } from './load-average.component';

describe('LoadAverageComponent', () => {
  let component: LoadAverageComponent;
  let fixture: ComponentFixture<LoadAverageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LoadAverageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LoadAverageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
