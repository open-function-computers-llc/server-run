import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-import-account-component',
  templateUrl: './import-account-component.component.html',
  styleUrls: ['./import-account-component.component.scss']
})
export class ImportAccountComponentComponent implements OnInit {
  filename: string = "";
  isComplete: boolean = false;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
  ) { }

  ngOnInit(): void {
    const filename:string = this.route.snapshot.paramMap.get("filename") || "";
    this.filename = filename;
  }

  generateScriptENV(): string {
    return `BACKUP_TARBALL=${this.filename}`;
  }

  onScriptCompleted(e:boolean) {
    console.log("completed!", e);
    this.isComplete = true;
  }

  navigateToAccounts() {
    this.router.navigate(["/system/accounts"]);
  }
}
