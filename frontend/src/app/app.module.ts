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
import { SitesComponent } from './pages/sites/sites.component';
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
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    NgxBootstrapIconsModule.pick(allIcons)
  ],
  providers: [
    AuthService,
    AuthGuard,
    UptimeService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
