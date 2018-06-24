import { Action } from '@ngrx/store';
import { Bookmark } from '../../../core/models';

export enum PreviewActionTypes {
    REQUEST = '[Preview] Request',
    URL_CHANGED = '[Preview] Url Changed',
    FAIL = '[Preview] Failure',
    SUCCESS = '[Preview] Success',
    CLEAR = '[Preview] Clear',
    SAVE = '[Preview] Save',
    SAVE_SUCCESS = '[Preview] Save Success',
    SAVE_FAILURE = '[Preview] Save Failure'
}

export class PreviewActionRequest implements Action {
    readonly type = PreviewActionTypes.REQUEST;
    constructor(public url: string) { }
}

export class PreviewActionFail implements Action {
    readonly type = PreviewActionTypes.FAIL;
    constructor(public msg: string) { }
}

export class PreviewActionSuccess implements Action {
    readonly type = PreviewActionTypes.SUCCESS;
    constructor(public item: Bookmark) { }
}

export class PreviewActionClear implements Action {
    readonly type = PreviewActionTypes.CLEAR;
    constructor() { }
}

export class PreviewActionSave implements Action {
    readonly type = PreviewActionTypes.SAVE;
    constructor(public item: Bookmark) { }
}

export class PreviewActionSaveSucess implements Action {
    readonly type = PreviewActionTypes.SAVE_SUCCESS;
    constructor(public item: Bookmark) { }
}

export class PreviewActionSaveFailure implements Action {
    readonly type = PreviewActionTypes.SAVE_FAILURE;
    constructor(public msg: string) { }
}

export type PreviewActionUnion =
    | PreviewActionRequest
    | PreviewActionSuccess
    | PreviewActionFail
    | PreviewActionClear
    | PreviewActionSave
    | PreviewActionSaveSucess
    | PreviewActionSaveFailure;
