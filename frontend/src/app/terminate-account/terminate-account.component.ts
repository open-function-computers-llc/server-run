import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-terminate-account',
  templateUrl: './terminate-account.component.html',
  styleUrls: ['./terminate-account.component.scss']
})
export class TerminateAccountComponent implements OnInit {
  @Input()
  account: string = "";

  terminationVerified: boolean = false;
  terminationCompleted: boolean = false;

  constructor(
    private router: Router,
  ) { }

  ngOnInit(): void {
  }

  verifyTermination() {
    this.terminationVerified = true;
  }

  onScriptCompleted(e:boolean): void {
    this.terminationCompleted = true;
  }

  toAllAccounts() : void {
    this.router.navigate(["/system/accounts"]);
  }

  generateScriptENV(): string {
    return `REQUESTED_ACCOUNT=${this.account}`;
  }
}
