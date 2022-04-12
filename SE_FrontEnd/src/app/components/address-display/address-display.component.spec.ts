import { Component, NO_ERRORS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ContactInfoPage } from 'src/app/pages/contact-info/contact-info.page';
import { UnderScore } from 'src/app/pipes/under-score/under-score';
import { AddressDisplayComponent } from './address-display.component';

@Component({
  selector: 'app-contact-info',
  template: '',
  providers: [
    {provide: ContactInfoPage, useClass: ContactInfoPageStub }
  ]
})
class ContactInfoPageStub {
  init = () => {};
}


describe('AddressDisplayComponent', () => {
  let component: AddressDisplayComponent;
  let fixture: ComponentFixture<AddressDisplayComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      schemas: [NO_ERRORS_SCHEMA],
      declarations: [ AddressDisplayComponent, UnderScore, ContactInfoPageStub ],
      imports: [ BrowserAnimationsModule]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AddressDisplayComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
