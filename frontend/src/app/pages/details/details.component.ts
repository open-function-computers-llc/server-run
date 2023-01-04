import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { first, map, Observable, share } from 'rxjs';
import { ServerService } from 'src/app/server.service';
import { UptimeService } from 'src/app/uptime/uptime.service';
import { Website } from 'src/app/Website';

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss']
})
export class DetailsComponent implements OnInit {
  site: Observable<Website>;
  analyticsView: string = "1";
  domain: string = "";
  analyticsPath: string = "";
  showingPubkey: boolean = false;
  showingDomains: boolean = false;
  showingChart: boolean = false;
  showingExport: boolean = false;
  chartType: string = "total-requests";
  showingTerminateVerification: boolean = false;
  temporaryCopyAnimationShowing: boolean = false;

  constructor(
    private serverService: ServerService,
    private route: ActivatedRoute,
    private router: Router,
  ) { }

  ngOnInit(): void {
    const domain:string = this.route.snapshot.paramMap.get("domain") || "";
    this.domain = domain;
    this.site = this.serverService.getSiteDetails(domain);
    this.setAnalyticsURL();
  }

  setAnalyticsURL() {
    const token = this.serverService.getToken();
    this.analyticsPath = "/api/analytics?domain=" + this.domain + "&period=" + this.analyticsView + "&token=" + token;
  }

  setDetailView(e:any) {
    const selectedValue = e.target.value;
    if (selectedValue === "") {
      this.showingDomains = false;
      this.showingPubkey = false;
      this.showingChart = false;
      this.showingTerminateVerification = false;
      this.showingExport = false;
      return;
    }

    if (selectedValue === 'showPubKey') {
      this.showingPubkey = true;
      this.showingDomains = false;
      this.showingChart = false;
      this.showingExport = false;
      this.showingTerminateVerification = false;
      return;
    }

    if (selectedValue === "terminateAccount") {
      this.showingPubkey = false;
      this.showingDomains = false;
      this.showingChart = false;
      this.showingExport = false;
      this.showingTerminateVerification = true;
      return;
    }

    if (selectedValue === 'showDomains') {
      this.showingDomains = true;
      this.showingPubkey = false;
      this.showingChart = false;
      this.showingExport = false;
      this.showingTerminateVerification = false;
      return;
    }

    if (selectedValue === 'analyticChart') {
      this.showingDomains = false;
      this.showingPubkey = false;
      this.showingChart = true;
      this.showingExport = false;
      this.showingTerminateVerification = false;
      return;
    }

    if (selectedValue === 'prepareForExport') {
      this.showingDomains = false;
      this.showingPubkey = false;
      this.showingChart = false;
      this.showingExport = true;
      this.showingTerminateVerification = false;
      return;
    }
  }

  sshPubkeyToClipboard(key:string) {
    navigator.clipboard.writeText(key);
    this.temporaryCopyAnimationShowing = true;
    setTimeout(() => {
      this.temporaryCopyAnimationShowing = false;
    }, 1000);
  }

  unlockSite() {
    this.router.navigate(['process', 'unlock'], { relativeTo: this.route });
  }

  lockSite() {
    this.router.navigate(['process', 'lock'], { relativeTo: this.route });
  }

  cloneAccount() {
    this.router.navigate(["..", "clone", this.domain], { relativeTo: this.route });
  }

  generateExportENV(): string {
    return `ACCOUNT_NAME=${this.domain}`; // TODO: update "domain" to "account" pretty much everywhere
  }
}
