import { Component, Input, Output, EventEmitter } from '@angular/core';
import { shrinkInOut, fadeInOut } from 'src/app/animations/animations';
import { ContactInfo } from 'usthaan-ts-types';

@Component({
  selector: 'app-address-display',
  templateUrl: './address-display.component.html',
  styleUrls: ['./address-display.component.scss'],
  animations: [
    shrinkInOut('shrinkInOut', 125, 100),
    fadeInOut('fadeInOut', 100, 500),
  ]
})
export class AddressDisplayComponent {

  @Input() contact: ContactInfo;
  @Input() compact: boolean;
  @Input() header: boolean;
  @Input() controls = true;
  @Output() edit = new EventEmitter();
  @Output() create = new EventEmitter();
  @Output() revert = new EventEmitter();
  @Output() picking = new EventEmitter();
  @Output() selectedContact = new EventEmitter();
  editing = false;
  pick = false;

  constructor() { }

  formatter() {
    const addrFieldsInSequence = ['line1', 'line2', 'landmark', 'area', 'state', 'city', 'zip', 'country'];
    return stringifyObj(this.contact.address);
    function stringifyObj(objectToStringify, fieldsInSequence = addrFieldsInSequence, separator = ', ') {
      return fieldsInSequence.reduce((result, field) => {
        if (result.length) {
          if (objectToStringify[field]) {
            if (field !== 'zip' && field !== 'country') {
              result = result + separator + objectToStringify[field];
            } else {
              result = result + ' - ' + objectToStringify[field];
            }
          }
        } else if (objectToStringify[field]) {
          result = objectToStringify[field];
        }
        return result.toString();
      }, '');
    }
  }

  editContact() {
    this.pick = false;
    this.editing = !this.editing;
    this.edit.emit(this.editing);
  }

  cancelContact() {
    this.editing = false;
    this.revert.emit();
  }

  createContact() {
    this.editing = true;
    this.edit.emit(this.editing);
    this.create.emit();
  }

  pickContact() {
    this.pick = !this.pick;
    this.editing = this.pick ? false : this.editing;
    this.picking.emit();
  }

  pickedContact(contact: ContactInfo) {
    this.pick = false;
    this.selectedContact.emit(contact);
  }

}
