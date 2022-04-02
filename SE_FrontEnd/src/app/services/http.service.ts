import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { forkJoin,Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { environment as env } from 'src/environments/environment';
import { APIResponse, Item,Product } from '../models';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  constructor(private http: HttpClient) { }

  getItemList(
    ordering: string,
    search?: string
  ): Observable<APIResponse<Product>> {
    let params = new HttpParams().set('ordering',ordering);

    if(search) {
      params = new HttpParams().set('ordering',ordering).set('search', search);
    }
    
    return this.http.get<APIResponse<Product>>(`${env.BASE_URL}/`,{
      params:params
    });
 
  }

  getItemDetails(id: string): Observable<Product> {
    const gameInfoRequest = this.http.get(`${env.BASE_URL}/product/${id}`);
    

    return forkJoin({
      gameInfoRequest

    }).pipe(
      map((resp: any) => {
        return {
          ...resp['gameInfoRequest']
        };
      })
    );
  }
}
