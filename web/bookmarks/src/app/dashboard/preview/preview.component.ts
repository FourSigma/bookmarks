import { Component, OnInit } from '@angular/core';
import { BookmarkService } from '../../core/service';
import { Bookmark } from '../../core/models';
import { Observable, Subscription} from 'rxjs';

@Component({
  selector: 'preview',
  templateUrl: './preview.component.html',
  styleUrls: ['./preview.component.css']
})
export class PreviewComponent implements OnInit{

  constructor(public bookmark: BookmarkService){}
  ngOnInit():void{

  }

  private sub: Subscription;
  private pBookmark:Bookmark;
  private pError: string;
  preview(url:string): Observable<Bookmark>{
    return this.bookmark.preview('https://www.popsci.com/forever-man-immortality-science');
    // return this.bookmark.preview('https://www.popsci.com/forever-man-immortality-science').subscribe(
    //   (bls: Bookmark) => return bls,
    //   (err) => console.log(err),
    // );

  }

}
