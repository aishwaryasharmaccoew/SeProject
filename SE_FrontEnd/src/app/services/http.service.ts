import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment as env } from 'src/environments/environment';
import { APIResponse, Item } from '../models';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  constructor(private http: HttpClient) { }

  getItemList(
    ordering: string,
    search?: string
  ): Observable<APIResponse<Item>> {
    let params = new HttpParams().set('ordering',ordering);
    if(search) {
      params = new HttpParams().set('ordering',ordering).set('search', search);
    }
    return this.http.get<APIResponse<Item>>(`${env.BASE_URL}/items`,{
      params: params,
    });
  }
}
