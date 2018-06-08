import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { RouterModule } from '@angular/router';
import { MatGridListModule, MatCardModule, MatMenuModule, MatIconModule, MatButtonModule } from '@angular/material';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    RouterModule
  ],
  declarations: [],
  exports: [
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    RouterModule,
    BrowserAnimationsModule,
    MatGridListModule,
    MatCardModule,
    MatMenuModule,
    MatIconModule,
    MatButtonModule
  ]
})
export class SharedModule { }