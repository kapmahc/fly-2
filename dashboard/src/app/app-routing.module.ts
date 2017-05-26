import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import plugins from './plugins';

@NgModule({
  imports: [ RouterModule.forRoot(plugins.routes) ],
  exports: [ RouterModule ]
})

export class AppRoutingModule {}
