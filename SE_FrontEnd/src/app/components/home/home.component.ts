import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { APIResponse, Item, Product } from '../../models';
import { HttpService } from 'src/app/services/http.service';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit, OnDestroy {
  stringJson: any;
  stringObject: any;
  public sort: string;
  public items: Array<Item>;
  public products: Array<Product>;
  private routeSub: Subscription;
  private productSub: Subscription;
 
  constructor(    
    private httpService: HttpService,
    private router: Router,
    private activatedRoute: ActivatedRoute) { }

  ngOnInit(): void {
    this.routeSub = this.activatedRoute.params.subscribe((params: Params) => {
      if (params['game-search']) {
        this.searchItems('metacrit', params['game-search']);
      } else {
        this.searchItems('metacrit');
      }
    });
  }

  searchItems(sort: string, search?: string): void {




    this.productSub = this.httpService
      .getItemList(sort, search)
      .subscribe((productList: APIResponse<Product>) => {
       
        this.stringJson = JSON.stringify(productList);
        console.log("String json object :", this.stringJson);
        console.log("Type :", typeof this.stringJson);
        this.products = JSON.parse(this.stringJson);
        console.log("JSON object -", this.stringObject);

      });
  }

  openItemDetails(id: string): void {
    this.router.navigate(['details', id]);
  }

  ngOnDestroy(): void {
    if (this.productSub) {
      this.productSub.unsubscribe();
    }

    if (this.routeSub) {
      this.routeSub.unsubscribe();
    }
  }

}
