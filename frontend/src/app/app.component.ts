import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { AuthService } from './auth/auth.service';
import { UptimeService } from './uptime/uptime.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit, OnDestroy {
  private userSub: Subscription;
  isLoggedIn: boolean = false;

  constructor(
    private authService: AuthService,
    private uptimeService: UptimeService,
  ) {}

  ngOnInit() {
    this.userSub = this.authService.user.subscribe((u) => {
      this.isLoggedIn = !!u;
      console.log(this.isLoggedIn, u);
      if (this.isLoggedIn) {
        this.uptimeService.bootstrapUptimeService();
      }
    });

    this.authService.autoLogin();
  }

  ngOnDestroy() {
    this.userSub.unsubscribe();
  }
}
