import { NgModule } from '@angular/core';
import { SharedModule } from '../shared';
import { CommonModule } from '@angular/common';
import { DashboardComponent } from './dashboard.component';
import { PreviewComponent } from './containers/preview/preview.component';
import { CardComponent } from './components/card/card.component';
import { DashboardStoreModule } from './store/dashboard-store.module';


@NgModule({
  declarations: [
    DashboardComponent,
    PreviewComponent,
    CardComponent,
  ],
  imports: [
    CommonModule,
    SharedModule,
    DashboardStoreModule,
  ],
  providers: [],
  exports: [
    DashboardComponent,
  ]
})

export class DashboardModule { }
