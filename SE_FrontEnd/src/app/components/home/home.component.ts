import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { APIResponse, Item } from '../../models';
import { HttpService } from 'src/app/services/http.service';

interface Food {
  value: string;
  viewValue: string;
}
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit, OnDestroy {
  public sort: string;
  public items: Array<Item>;
  private routeSub: Subscription;
  private itemSub: Subscription;
  foods: Food[] = [
    {value: 'steak-0', viewValue: 'Steak'},
    {value: 'pizza-1', viewValue: 'Pizza'},
    {value: 'tacos-2', viewValue: 'Tacos'},
  ];
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
    this.itemSub = this.httpService
      .getItemList(sort, search)
      .subscribe((itemList: APIResponse<Item>) => {
        this.items = itemList.results;
        console.log(itemList);
      });
  }

  openItemDetails(id: string): void {
    this.router.navigate(['details', id]);
  }

  ngOnDestroy(): void {
    if (this.itemSub) {
      this.itemSub.unsubscribe();
    }

    if (this.routeSub) {
      this.routeSub.unsubscribe();
    }
  }

}
