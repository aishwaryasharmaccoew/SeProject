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
      params:params/*,
      headers: {"Access-Control-Allow-Origin": "*", 
      "Access-Control-Allow-Headers": "access-control-allow-origin, access-control-allow-headers",
      "Content-Type":"application/json; charset=utf-8"
    }    */
    });
 
  }

  getItemDetails(id: string): Observable<Item> {
    const gameInfoRequest = this.http.get(`${env.BASE_URL}/product/${id}`);
    const gameTrailersRequest = this.http.get(
      `${env.BASE_URL}/games/${id}/movies`
    );
    const gameScreenshotsRequest = this.http.get(
      `${env.BASE_URL}/games/${id}/screenshots`
    );

    return forkJoin({
      gameInfoRequest,
      gameScreenshotsRequest,
      gameTrailersRequest,
    }).pipe(
      map((resp: any) => {
        return {
          ...resp['gameInfoRequest'],
          screenshots: resp['gameScreenshotsRequest']?.results,
          trailers: resp['gameTrailersRequest']?.results,
        };
      })
    );
  }
}
