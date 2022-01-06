import { Component, OnDestroy, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { AuthService } from '../auth/auth.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit, OnDestroy {
  private userSub: Subscription;
  isLoggedIn: boolean = false;

  constructor(private authService: AuthService, private router: Router) { }

  logout() {
    this.authService.user.next(null);
    this.router.navigate(["/login"]);
  }

  ngOnInit(): void {
    this.userSub = this.authService.user.subscribe((u) => {
      this.isLoggedIn = !!u;
    });
  }

  ngOnDestroy(): void {
    this.userSub.unsubscribe();
  }

}
