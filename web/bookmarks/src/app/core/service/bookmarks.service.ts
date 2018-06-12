import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs';
import * as models from '../models';
import { map, switchMap } from 'rxjs/operators';

const BaseURL = 'http://localhost:8080/v0';
const HttpOpts = {
    headers: new HttpHeaders({
        'Content-Type': 'application/json',
    })
};
@Injectable({
    providedIn: 'root'
})
export class BookmarkService {
    constructor(private http: HttpClient) { }

    list(): Observable<models.Bookmark[]> {
        return this.http.get<models.Bookmark[]>(`${BaseURL}/bookmarks`, HttpOpts).pipe(
            map(resp => resp.map((item) =>models.Bookmark.fromJSON(resp))),
        );
    }

    preview(url:string): Observable<models.Bookmark> {
        return this.http.get<models.Bookmark[]>(`${BaseURL}/bookmarks?preview=${url}`, HttpOpts).pipe(
            map(resp => models.Bookmark.fromJSON(resp)),
        );
    }
}

