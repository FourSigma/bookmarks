import { createSelector } from '@ngrx/store';
import { createEntityAdapter, EntityAdapter, EntityState } from '@ngrx/entity';
import { Bookmark } from '../models';
import { BookmarkActionsUnion, BookmarkActionTypes } from './action';


export interface BookmarkState extends EntityState<Bookmark> {
    collection: Bookmark[];
    isLoaded: boolean;
    selectedBookmarkId: string | null;
}

export const bookmarkAdapter: EntityAdapter<Bookmark> = createEntityAdapter<Bookmark>({
    selectId: (book: Bookmark) => book.id,
    sortComparer: false,
});

export const initialState: BookmarkState = bookmarkAdapter.getInitialState({
    collection: [],
    isLoaded: false,
    selectedBookmarkId: null
});

export function bookmarkReducer(state: BookmarkState = initialState, action: BookmarkActionsUnion) {
    switch (action.type) {
        default:
            return state;
    }
}
