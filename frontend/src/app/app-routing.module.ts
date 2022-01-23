import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './auth/auth.guard';
import { ErrorComponent } from './error/error.component';
import { DetailsComponent } from './pages/details/details.component';
import { HomeComponent } from './pages/home/home.component';
import { LoginComponent } from './pages/login/login.component';
import { ProcessComponent } from './pages/process/process.component';
import { SitesComponent } from './pages/accounts/accounts.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: '', redirectTo: '/system/home', pathMatch: 'full' },
  { path: 'system', canActivate: [AuthGuard], children:[
    { path: 'home', component: HomeComponent },
    { path: 'accounts', component: SitesComponent },
    { path: 'accounts/:domain', component: DetailsComponent },
    { path: 'accounts/:domain/process/:action', component: ProcessComponent },
  ]},
  { path: '**', component: ErrorComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {
    useHash: true
  })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
