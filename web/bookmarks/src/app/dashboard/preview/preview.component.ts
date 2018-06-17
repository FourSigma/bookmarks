import { Component, OnInit } from '@angular/core';
import { BookmarkService } from '../../core/service';
import { Bookmark } from '../../core/models';
import { Observable, Subscription,of, empty, never, } from 'rxjs';
import {NgbModal, NgbModalRef} from '@ng-bootstrap/ng-bootstrap';
import { FormControl } from '@angular/forms';
import { debounceTime, distinctUntilChanged,map, switchMap, shareReplay, catchError} from 'rxjs/operators';

@Component({
  selector: 'preview',
  templateUrl: './preview.component.html',
  styleUrls: ['./preview.component.css']
})
export class PreviewComponent implements OnInit{

  private formURL = new FormControl();
  private modalRef: NgbModalRef;
  private bookmark$:Observable<Bookmark>;
  private sub: Subscription;

  constructor(public bookmark: BookmarkService, private modal: NgbModal){}

  ngOnInit():void{
    this.bookmark$= this.formURL.valueChanges.pipe(
      debounceTime(100),
      distinctUntilChanged(),
      switchMap((url:string) => {
        if (!url || url === ''){
          return of(new Bookmark())
        }
        return this.preview(url)
      }),
    ); 
  }

  ngOnDestroy(){
    empty()
  }
  private clear():void{
    this.formURL.setValue('');
  }
  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
   
      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead
   
      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);
      this.log(error);
   
      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private pError: string;


  preview(url:string): Observable<Bookmark>{
    return this.bookmark.preview(url).pipe(
      catchError(this.handleError<Bookmark>("PREVIEW::")),
    );
  }

  save(bm:Bookmark){
    this.bookmark.create(bm).subscribe(
      resp => console.log("Bookmark created: ", resp)
    )
  }
  openLg(content) {
   this.modalRef = this.modal.open(content, { size: 'lg' });
   this.modalRef.result.then(()=>this.formURL.setValue(''),()=>this.formURL.setValue(''))
  }

  log(msg:any){
    console.log(msg);
  }

}
