import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-export-account',
  templateUrl: './export-account.component.html',
  styleUrls: ['./export-account.component.scss']
})
export class ExportAccountComponent implements OnInit {
  accountName:string = "";
  exportCompleted:boolean = false;

  constructor(
    private route: ActivatedRoute,
  ) { }

  ngOnInit(): void {
    const accountName:string = this.route.snapshot.paramMap.get("domain") || "";
    this.accountName = accountName;
  }

  generateScriptENV(): string {
    return `ACCOUNT=${this.accountName}`;
  }

  onScriptCompleted(e:boolean): void {
    this.exportCompleted = true;
  }

}
