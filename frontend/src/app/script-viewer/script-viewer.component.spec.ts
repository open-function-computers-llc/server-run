import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ScriptViewerComponent } from './script-viewer.component';

describe('ScriptViewerComponent', () => {
  let component: ScriptViewerComponent;
  let fixture: ComponentFixture<ScriptViewerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ScriptViewerComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ScriptViewerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
