import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './auth/auth.guard';
import { ErrorComponent } from './error/error.component';
import { DetailsComponent } from './pages/details/details.component';
import { HomeComponent } from './pages/home/home.component';
import { LoginComponent } from './pages/login/login.component';
import { ProcessComponent } from './pages/process/process.component';
import { SitesComponent } from './pages/accounts/accounts.component';
import { ScriptViewerComponent } from './script-viewer/script-viewer.component';
import { BoostrapComponent } from './boostrap/boostrap.component';
import { FailToBanComponent } from './pages/fail-to-ban/fail-to-ban.component';
import { AddAccountComponent } from './pages/add-account/add-account.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: '', redirectTo: '/system/home', pathMatch: 'full' },
  { path: 'system', canActivate: [AuthGuard], component: BoostrapComponent, children:[
    { path: 'home', component: HomeComponent },
    { path: 'accounts', component: SitesComponent },
    { path: 'accounts/add', component: AddAccountComponent },
    { path: 'accounts/clone/:domain', component: AddAccountComponent },
    { path: 'accounts/:domain', component: DetailsComponent },
    { path: 'accounts/:domain/process/:action', component: ProcessComponent },
    { path: 'f2ban/status', component: FailToBanComponent },
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
