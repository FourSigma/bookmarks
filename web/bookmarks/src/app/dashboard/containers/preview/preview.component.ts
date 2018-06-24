import { Component, OnInit, Output, EventEmitter, OnDestroy } from '@angular/core';
import { Bookmark } from '../../../core/models';
import { Observable, Subscription, of, empty, never, } from 'rxjs';
import { NgbModal, NgbModalRef } from '@ng-bootstrap/ng-bootstrap';
import { FormControl } from '@angular/forms';
import { Store } from '@ngrx/store';
import { PreviewState, PreviewActionRequest, PreviewActionClear, PreviewActionSave } from '../../store/preview';
import { debounceTime, distinctUntilChanged, map, take, switchMap, shareReplay, catchError } from 'rxjs/operators';
import { getPreviewState, getPreviewBookmark, DashboardState, getDashboardState } from '../../store/dashboard-store.module';

@Component({
  selector: 'preview',
  templateUrl: './preview.component.html',
  styleUrls: ['./preview.component.css']
})
export class PreviewComponent implements OnInit, OnDestroy {

  private formURL = new FormControl();
  private modalRef: NgbModalRef;
  private state$: Observable<PreviewState>;
  private sub: Subscription;
  private bookmark$: Observable<Bookmark>;

  constructor(
    private modal: NgbModal,
    private store: Store<DashboardState>,
  ) { }

  ngOnInit(): void {
    this.state$ = this.store.select(getPreviewState);
    this.bookmark$ = this.store.select(getPreviewBookmark);

    this.sub = this.formURL.valueChanges.pipe(
      debounceTime(100),
      distinctUntilChanged(),
    ).subscribe(url => this.store.dispatch<PreviewActionRequest>(new PreviewActionRequest(url)));
  }

  ngOnDestroy() {
    this.sub.unsubscribe();
  }

  preview(url: string): void {
    this.store.dispatch<PreviewActionRequest>(new PreviewActionRequest(url));
  }

  save(bm: Bookmark): void {
    this.store.dispatch<PreviewActionSave>(new PreviewActionSave(bm));
  }

  clear(): void {
    this.formURL.setValue('');
    this.store.dispatch<PreviewActionClear>(new PreviewActionClear());
  }

  openLg(content) {
    this.clear();
    this.modalRef = this.modal.open(content, { size: 'lg' });
  }
}
