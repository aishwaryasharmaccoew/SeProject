import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FeatureSliderComponent } from './feature-slider.component';

describe('FeatureSliderComponent', () => {
  let component: FeatureSliderComponent;
  let fixture: ComponentFixture<FeatureSliderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FeatureSliderComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FeatureSliderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
