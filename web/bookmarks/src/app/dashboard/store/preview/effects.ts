import { Actions, Effect, ofType } from '@ngrx/effects';
import { Injectable } from '@angular/core';
import { BookmarkService } from '../../../core/service';
import { PreviewActionTypes, PreviewActionRequest, PreviewActionSuccess, PreviewActionFail } from './actions';
import { Action } from '@ngrx/store';
import { switchMap, map, catchError, mergeMap } from 'rxjs/operators';
import { Observable, of, empty } from 'rxjs';



@Injectable()
export class PreviewEffects {
    constructor(public bookmark: BookmarkService, private actions$: Actions) { }

    @Effect()
    get: Observable<Action> = this.actions$.pipe(
        ofType<PreviewActionRequest>(PreviewActionTypes.REQUEST),
        switchMap((action: PreviewActionRequest) => {
            if (action.url === '') {
                return empty();
            }
            return this.bookmark.preview(action.url).pipe(
                map(data => new PreviewActionSuccess(data)),
                catchError(err => of(new PreviewActionFail(err))),
            );
        }
        ),
    );

    // @Effect()
    // save: Observable<Action> = this.actions$.pipe(

    // );

}


