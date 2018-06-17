import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs';
import * as models from '../models';
import { map, switchMap, tap } from 'rxjs/operators';

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
        return this.http.get<models.Bookmark[]>(`${BaseURL}/bookmarks`, HttpOpts);
    }

    preview(url:string): Observable<models.Bookmark> {
        return this.http.get<models.Bookmark>(`${BaseURL}/bookmarks?preview=${url}`, HttpOpts).pipe(
            map(resp => models.Bookmark.fromJSON(resp)),
        );
    }

    create(b: models.Bookmark): Observable<models.Bookmark> {
        return this.http.post<models.Bookmark>(`${BaseURL}/bookmarks`,b, HttpOpts);
    }
}

