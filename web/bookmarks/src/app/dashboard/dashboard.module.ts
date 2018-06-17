import { NgModule } from '@angular/core';
import { SharedModule } from '../shared';
import { CommonModule } from '@angular/common';
import { DashboardComponent } from './dashboard.component';
import { PreviewComponent} from './preview/preview.component';
import { CardComponent } from './card/card.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { NgbModule} from '@ng-bootstrap/ng-bootstrap';


@NgModule({
  declarations: [
    DashboardComponent,
    PreviewComponent,
    CardComponent,
  ],
  imports: [
    CommonModule,
    FlexLayoutModule,
    SharedModule,
    NgbModule.forRoot(),
  ],
  providers: [],
  exports: [ 
      DashboardComponent,
  ]
})

export class DashboardModule { }
