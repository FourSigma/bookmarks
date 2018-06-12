import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
})

const URLGetBookmarks = 'http://localhost:4040/bookmarks';
const HttpOpts = {
    headers: new HttpHeaders({
        'Content-Type': 'application/json',
    })
};
export class BookmarkService {
    constructor(private http: HttpClient) { }

    getBookmarks(): Observable<any> {
        return this.http.get<any>(URLGetBookmarks, HttpOpts);
    }


}
