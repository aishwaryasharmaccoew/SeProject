import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Subscription } from 'rxjs';
import { Item, Product } from 'src/app/models';
import { HttpService } from 'src/app/services/http.service';

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss']
})
export class DetailsComponent implements OnInit, OnDestroy {
  stringJson: any;
  stringObject: any;
  itemRating = 0;

  productID: string;
  item: Item;
  product: Product;
  routeSub: Subscription;
  apiSub: Subscription;

  constructor(
    private activatedRoute: ActivatedRoute,
    private httpService: HttpService
  ) { }

  ngOnInit(): void {
    this.routeSub = this.activatedRoute.params.subscribe((params: Params) => {
      this.productID = params['id'];
      this.getProductDetails(this.productID);
    });
  }

  getProductDetails(id: string): void {
    this.apiSub = this.httpService
      .getItemDetails(id)
      .subscribe((apiResp: Product) => {
        this.product= apiResp;

      });
  }

  getColor(value: number): string {
    if (value > 75) {
      return '#5ee432';
    } else if (value > 50) {
      return '#fffa50';
    } else if (value > 30) {
      return '#f7aa38';
    } else {
      return '#ef4655';
    }
  }


  myFunction(){
    
    var popup = document.getElementById("myPopup");
    if (popup!=null) {
      popup.classList.toggle("show");
    }
    
  }

  ngOnDestroy(): void {
    if (this.apiSub) {
      this.apiSub.unsubscribe();
    }

    if (this.routeSub) {
      this.routeSub.unsubscribe();
    }
  }

}