import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-add-account',
  templateUrl: './add-account.component.html',
  styleUrls: ['./add-account.component.scss']
})
export class AddAccountComponent implements OnInit {
  newAccountName:string = "";
  isAddingAccount: boolean = false;
  completedAccountCreation: boolean = false;
  cloningFrom:string = "";

  constructor(
    private location: Location,
    private route: ActivatedRoute,
  ) { }

  ngOnInit(): void {
    const cloningFrom:string = this.route.snapshot.paramMap.get("domain") || "";
    this.cloningFrom = cloningFrom;
  }

  goBack(): void {
    this.location.back();
  }

  addAccount(): void {
    this.isAddingAccount = true;
  }

  onScriptCompleted(e:boolean): void {
    this.completedAccountCreation = true;
  }

  scriptToRun(): string {
    if (this.cloningFrom) {
      return "cloneAccount";
    }
    return "addAccount";
  }

  generateScriptENV(): string {
    if (this.cloningFrom) {
      return `SOURCE_ACCOUNT=${this.cloningFrom}|DESTINATION_ACCOUNT=${this.newAccountName}`;
    }
    return `ACCOUNT_NAME=${this.newAccountName}`;
  }
}
