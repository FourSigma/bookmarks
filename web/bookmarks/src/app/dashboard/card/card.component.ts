import { Component, OnInit } from '@angular/core';
import { BookmarkService } from '../../core/service';
import { Bookmark } from '../../core/models';
import { Observable, Subscription} from 'rxjs';

@Component({
  selector: 'card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.css']
})
export class CardComponent implements OnInit{

cards = [
    { title: 'Card 1', cols: 2, rows: 1 },
    { title: 'Card 2', cols: 1, rows: 1 },
    { title: 'Card 3', cols: 1, rows: 2 },
    { title: 'Card 3', cols: 1, rows: 2 },
    { title: 'Card 3', cols: 1, rows: 2 },
    { title: 'Card 3', cols: 1, rows: 2 },
    { title: 'Card 3', cols: 1, rows: 2 },
    { title: 'Card 3', cols: 1, rows: 2 },
    { title: 'Card 3', cols: 1, rows: 2 },
];

  constructor(){}

  ngOnInit():void{

  }

}
