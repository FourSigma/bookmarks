import { Component, OnInit, Input } from '@angular/core';
import { BookmarkService } from '../../../core/service';
import { Bookmark } from '../../../core/models';
import { Observable, Subject } from 'rxjs';
import { switchMap, catchError } from 'rxjs/operators';

@Component({
  selector: 'card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.css']
})
export class CardComponent implements OnInit {

  constructor() { }

  private bookmark$: Subject<Bookmark[]> = new Subject<Bookmark[]>();

  @Input() list: Bookmark[] = [];

  ngOnInit(): void { }


}
