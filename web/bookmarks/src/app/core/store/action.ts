import { Action } from '@ngrx/store';
import { Bookmark } from '../models';

export enum BookmarkActionTypes {
    Add = '[Bookmark] Add Bookmark',
    AddSuccess = '[Bookmark] Add Bookmark Success',
    AddFail = '[Bookmark] Add Bookmark Fail',
    Remove = '[Bookmark] Remove Bookmark',
    RemoveSuccess = '[Bookmark] Remove Bookmark Success',
    RemoveFail = '[Bookmark] Remove Bookmark Fail',
    Load = '[Bookmark] Load',
    LoadSuccess = '[Bookmark] Load Success',
    LoadFail = '[Bookmark] Load Fail',
}

/**
 * Add Bookmark to Bookmark Actions
 */
export class BookmarkAdd implements Action {
    readonly type = BookmarkActionTypes.Add;

    constructor(public bookmark: Bookmark) { }
}

export class BookmarkAddSuccess implements Action {
    readonly type = BookmarkActionTypes.AddSuccess;

    constructor(public bookmark: Bookmark) { }
}

export class BookmarkAddFail implements Action {
    readonly type = BookmarkActionTypes.AddFail;

    constructor(public bookmark: Bookmark) { }
}

/**
 * Remove Bookmark from Bookmark Actions
 */
export class BookmarkRemove implements Action {
    readonly type = BookmarkActionTypes.Remove;

    constructor(public bookmark: Bookmark) { }
}

export class BookmarkRemoveSuccess implements Action {
    readonly type = BookmarkActionTypes.RemoveSuccess;

    constructor(public bookmark: Bookmark) { }
}

export class BookmarkRemoveFail implements Action {
    readonly type = BookmarkActionTypes.RemoveFail;

    constructor(public bookmark: Bookmark) { }
}

/**
 * Load Bookmark Actions
 */
export class BookmarkLoad implements Action {
    readonly type = BookmarkActionTypes.Load;
}

export class BookmarkLoadSuccess implements Action {
    readonly type = BookmarkActionTypes.LoadSuccess;

    constructor(public list: Bookmark[]) { }
}

export class BookmarkLoadFail implements Action {
    readonly type = BookmarkActionTypes.LoadFail;

    constructor(public msg: string) { }
}

export type BookmarkActionUnion =
    | BookmarkAdd
    | BookmarkAddSuccess
    | BookmarkAddFail
    | BookmarkRemove
    | BookmarkRemoveSuccess
    | BookmarkRemoveFail
    | BookmarkLoad
    | BookmarkLoadSuccess
    | BookmarkLoadFail;
