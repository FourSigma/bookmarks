import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BookmarkService} from './service';
import { HttpClientModule} from '@angular/common/http';
import {RouterModule} from '@angular/router';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    RouterModule,
    HttpClientModule,
  ],
  providers: [
    BookmarkService
  ],
  declarations: [
    // NavComponent,
    // FooterComponent,
    // SearchBarComponent,
    // Error404Component
  ],
  exports: [
    RouterModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
  ],
})
export class CoreModule { }
