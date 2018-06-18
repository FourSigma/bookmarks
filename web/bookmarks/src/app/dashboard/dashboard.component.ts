import { Component, OnInit } from '@angular/core';
import { BookmarkService } from '../core/service';
import { Bookmark } from '../core/models';
import { Observable, Subscription } from 'rxjs';

@Component({
  selector: 'dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  constructor(public bookmark: BookmarkService) { }

  public list: Bookmark[] = [];

  ngOnInit(): void {
    this.getBookmarks();
  }

  getBookmarks(): void {
    this.bookmark.list().subscribe(
      (resp: Bookmark[]) => this.list = resp.reverse(),
    );
  }

  refresh() {
    this.getBookmarks();
  }

  addBookmark(b: Bookmark) {
    console.log("addBookmark", b);
    this.list.push(b);
  }
}
