import { async,ComponentFixture, TestBed } from '@angular/core/testing';
import { Component, NO_ERRORS_SCHEMA } from '@angular/core';
import { DetailsComponent } from './details.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('DetailsComponent', () => {
  let component: DetailsComponent;
  let fixture: ComponentFixture<DetailsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      schemas: [NO_ERRORS_SCHEMA],
      declarations: [  ],
      imports: [ BrowserAnimationsModule]
    })
    .compileComponents();
  }));
  
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DetailsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should load filtered results', () => {
    expect(hostComponent.autoCompleteComponent).toBeTruthy();
    expect(hostComponent.autoCompleteComponent.filteredTypes).toBeFalsy();
    const testData = require("./test-data/auto-complete-load-filtered.json");
    hostComponent.autoCompleteComponent.typesList = testData.typesList;
    hostComponent.autoCompleteComponent.fcn = "name";
    hostComponent.autoCompleteComponent.form = hostComponent.fb.group({
      name: ['', []],
    });
    const results = [];
    hostFixture.detectChanges();
    hostComponent.autoCompleteComponent.filteredTypes.subscribe(values => {
      results.push(values);
    });
    hostComponent.autoCompleteComponent.form.controls["name"].setValue(testData.updatedNameVal);
    hostFixture.detectChanges();
    expect(results).toEqual(testData.results);
  });
});
