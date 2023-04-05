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
import { AllDomainsComponent } from './pages/all-domains/all-domains.component';
import { ImportableAccountsComponentComponent } from './pages/importable-accounts-component/importable-accounts-component.component';
import { ImportAccountComponentComponent } from './pages/import-account-component/import-account-component.component';
import { ExportAccountComponent } from './pages/export-account/export-account.component';
import { ServiceListComponent } from './pages/service-list/service-list.component';
import { ServiceDetailsComponent } from './pages/service-details/service-details.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: '', redirectTo: '/system/home', pathMatch: 'full' },
  { path: 'system', canActivate: [AuthGuard], component: BoostrapComponent, children:[
    { path: 'home', component: HomeComponent },
    { path: 'accounts', component: SitesComponent },
    { path: 'accounts/add', component: AddAccountComponent },
    { path: 'accounts/clone/:domain', component: AddAccountComponent },
    { path: 'accounts/export/:domain', component: ExportAccountComponent },
    { path: 'accounts/import/:filename', component: ImportAccountComponentComponent },
    { path: 'accounts/:domain', component: DetailsComponent },
    { path: 'accounts/:domain/process/:action', component: ProcessComponent },
    { path: 'f2ban/status', component: FailToBanComponent },
    { path: 'all-domains', component: AllDomainsComponent },
    { path: 'show-importable-accounts', component: ImportableAccountsComponentComponent },
    { path: 'service-list', component: ServiceListComponent },
    { path: 'service/:name', component: ServiceDetailsComponent },
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
