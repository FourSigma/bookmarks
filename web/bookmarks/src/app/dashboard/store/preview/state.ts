import { createEntityAdapter, EntityAdapter, EntityState } from '@ngrx/entity';
import { Action } from '@ngrx/store';
import { PreviewActionTypes, PreviewActionUnion } from './actions';
import { Bookmark } from '../../../core/models';

export interface PreviewState {
    isLoading?: boolean;
    bookmark: Bookmark | undefined;
    error?: string | null;
}
export const previewInitialState: PreviewState = {
    isLoading: false,
    bookmark: undefined,
    error: null
};

export function previewReducer(state: PreviewState = previewInitialState, action: PreviewActionUnion): PreviewState {
    switch (action.type) {
        case PreviewActionTypes.FAIL:
            return {
                ...state,
                isLoading: false,
                error: action.msg,
                bookmark: undefined,
            };
        case PreviewActionTypes.SAVE_FAILURE:
            return {
                ...state,
                isLoading: false,
                error: action.msg,
            };

        case PreviewActionTypes.REQUEST:
            return {
                ...state,
                isLoading: true,
                error: null,
                bookmark: undefined,
            };
        case PreviewActionTypes.SUCCESS:
            const p = {
                ...state,
                isLoading: false,
                error: null,
                bookmark: action.item,
            };
            console.log(state);
            console.log('p', p);
            return p;
        case PreviewActionTypes.CLEAR:
            return {
                ...state,
                ...previewInitialState
            };
        default:
            return state;
    }
}
