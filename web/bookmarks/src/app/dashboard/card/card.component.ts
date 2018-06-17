import { Component, OnInit } from '@angular/core';
import { BookmarkService } from '../../core/service';
import { Bookmark } from '../../core/models';
import { Observable, Subscription} from 'rxjs';
import { switchMap, catchError} from 'rxjs/operators';

@Component({
  selector: 'card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.css']
})
export class CardComponent implements OnInit{

  constructor(public bookmark: BookmarkService){}

  private list$: Observable<Bookmark[]>;
  ngOnInit():void{
    this.list$ = this.bookmark.list()
  }

}
