import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { LoadAverageComponent } from './load-average/load-average.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { HomeComponent } from './pages/home/home.component';
import { LoginComponent } from './pages/login/login.component';
import { SitesComponent } from './pages/accounts/accounts.component';
import { DetailsComponent } from './pages/details/details.component';
import { FormsModule } from '@angular/forms';
import { ErrorComponent } from './error/error.component';
import { NgxBootstrapIconsModule, allIcons } from 'ngx-bootstrap-icons';
import { AuthService } from './auth/auth.service';
import { AuthGuard } from './auth/auth.guard';
import { SafeUrlPipePipe } from './pipes/safe-url-pipe.pipe';
import { ProcessComponent } from './pages/process/process.component';
import { UptimeComponent } from './uptime/uptime.component';
import { UptimeService } from './uptime/uptime.service';
import { ScriptViewerComponent } from './script-viewer/script-viewer.component';
import { BoostrapComponent } from './boostrap/boostrap.component';
import { FailToBanComponent } from './pages/fail-to-ban/fail-to-ban.component';
import { AddAccountComponent } from './pages/add-account/add-account.component';
import { TerminateAccountComponent } from './terminate-account/terminate-account.component';
import { AllDomainsComponent } from './pages/all-domains/all-domains.component';
import { HighchartsChartModule } from 'highcharts-angular';
import { ChartComponent } from './chart/chart.component';
import { ImportableAccountsComponentComponent } from './pages/importable-accounts-component/importable-accounts-component.component';
import { SafeHtmlPipe } from './pipes/safe-html.pipe';
import { ImportAccountComponentComponent } from './pages/import-account-component/import-account-component.component';
import { ExportAccountComponent } from './pages/export-account/export-account.component';
import { ServiceListComponent } from './pages/service-list/service-list.component';
import { ServiceDetailsComponent } from './pages/service-details/service-details.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    LoadAverageComponent,
    SidebarComponent,
    HomeComponent,
    LoginComponent,
    SitesComponent,
    DetailsComponent,
    ErrorComponent,
    SafeUrlPipePipe,
    ProcessComponent,
    UptimeComponent,
    ScriptViewerComponent,
    BoostrapComponent,
    FailToBanComponent,
    AddAccountComponent,
    TerminateAccountComponent,
    AllDomainsComponent,
    ChartComponent,
    ImportableAccountsComponentComponent,
    SafeHtmlPipe,
    ImportAccountComponentComponent,
    ExportAccountComponent,
    ServiceListComponent,
    ServiceDetailsComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    NgxBootstrapIconsModule.pick(allIcons),
    HighchartsChartModule,
  ],
  providers: [
    AuthService,
    AuthGuard,
    UptimeService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
