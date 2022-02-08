import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { AuthService } from './auth/auth.service';
import { User } from './auth/user.model';

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
  ) {}

  ngOnInit() {
    this.userSub = this.authService.user.subscribe((u:User|null) => {
      this.isLoggedIn = !!u;
    });

    this.authService.autoLogin();
  }

  ngOnDestroy() {
    this.userSub.unsubscribe();
  }
}
