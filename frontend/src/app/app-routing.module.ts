import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DetailsComponent } from './pages/details/details.component';
import { HomeComponent } from './pages/home/home.component';
import { LoginComponent } from './pages/login/login.component';
import { SitesComponent } from './pages/sites/sites.component';

const routes: Routes = [
  { path: 'home', component: HomeComponent },
  { path: 'sites', component: SitesComponent },
  { path: 'login', component: LoginComponent },
  { path: 'details/:domain', component:DetailsComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {
    useHash: true
  })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
