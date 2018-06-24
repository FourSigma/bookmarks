import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { EffectsModule } from '@ngrx/effects';
import { StoreModule } from '@ngrx/store';
import { previewReducer, PreviewState, PreviewEffects } from './preview';
import {
    ActionReducerMap, createFeatureSelector, createSelector
} from '@ngrx/store';


export interface DashboardState {
    preview: PreviewState;
}

export const dashboardReducer: ActionReducerMap<DashboardState> = {
    preview: previewReducer,
};

export const getDashboardState = createFeatureSelector<DashboardState>('dashboard');

export const getPreviewState = createSelector(
    getDashboardState,
    (state: DashboardState) => state.preview
);

export const getPreviewBookmark = createSelector(
    getPreviewState,
    (state: PreviewState) => state.bookmark,
);



@NgModule({
    imports: [
        CommonModule,
        StoreModule.forFeature('dashboard', dashboardReducer),
        EffectsModule.forFeature([PreviewEffects])
    ],
    providers: [PreviewEffects]
})
export class DashboardStoreModule { }


