import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-importable-accounts-component',
  templateUrl: './importable-accounts-component.component.html',
  styleUrls: ['./importable-accounts-component.component.scss']
})
export class ImportableAccountsComponentComponent implements OnInit {
  listedAllAccounts: boolean = false;

  constructor() { }

  ngOnInit(): void {
  }

  onScriptCompleted(e:boolean): void {
    this.listedAllAccounts = true;
  }

}
