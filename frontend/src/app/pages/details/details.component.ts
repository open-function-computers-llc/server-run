import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { first, map, Observable, share } from 'rxjs';
import { ServerService } from 'src/app/server.service';
import { UptimeService } from 'src/app/uptime/uptime.service';
import { Website } from 'src/app/Website';

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss'],
})
export class DetailsComponent implements OnInit {
  site: Observable<Website>;
  analyticsView: string = '1';
  domain: string = '';
  linkURL: string = '';
  analyticsPath: string = '';
  showingDomains: boolean = true;
  showingPubkey: boolean = false;
  showingChart: boolean = false;
  showingExport: boolean = false;
  monitorLogs: boolean = false;
  chartType: string = 'total-requests';
  showingTerminateVerification: boolean = false;
  temporaryCopyAnimationShowing: boolean = false;
  isAddingDomain: boolean = false;
  addDomainNow: boolean = false;
  newDomain: string = '';

  constructor(
    private serverService: ServerService,
    private route: ActivatedRoute,
    private router: Router
  ) {
    this.linkURL = 'https://' + this.domain;
  }

  ngOnInit(): void {
    const domain: string = this.route.snapshot.paramMap.get('domain') || '';
    this.domain = domain;
    this.site = this.serverService.getSiteDetails(domain);
    this.setAnalyticsURL();
  }

  setAnalyticsURL() {
    const token = this.serverService.getToken();
    this.analyticsPath =
      '/api/analytics?domain=' +
      this.domain +
      '&period=' +
      this.analyticsView +
      '&token=' +
      token;
  }

  setDetailView(e: any) {
    const selectedValue = e.target.value;
    if (selectedValue === '') {
      this.showingDomains = false;
      this.showingPubkey = false;
      this.showingChart = false;
      this.showingTerminateVerification = false;
      this.showingExport = false;
      this.monitorLogs = false;
      return;
    }

    if (selectedValue === 'showPubKey') {
      this.showingPubkey = true;
      this.showingDomains = false;
      this.showingChart = false;
      this.showingExport = false;
      this.showingTerminateVerification = false;
      this.monitorLogs = false;
      return;
    }

    if (selectedValue === 'terminateAccount') {
      this.showingPubkey = false;
      this.showingDomains = false;
      this.showingChart = false;
      this.showingExport = false;
      this.showingTerminateVerification = true;
      this.monitorLogs = false;
      return;
    }

    if (selectedValue === 'showDomains') {
      this.showingDomains = true;
      this.showingPubkey = false;
      this.showingChart = false;
      this.showingExport = false;
      this.showingTerminateVerification = false;
      this.monitorLogs = false;
      return;
    }

    if (selectedValue === 'analyticChart') {
      this.showingDomains = false;
      this.showingPubkey = false;
      this.showingChart = true;
      this.showingExport = false;
      this.showingTerminateVerification = false;
      this.monitorLogs = false;
      return;
    }

    if (selectedValue === 'prepareForExport') {
      this.showingDomains = false;
      this.showingPubkey = false;
      this.showingChart = false;
      this.showingExport = true;
      this.showingTerminateVerification = false;
      this.monitorLogs = false;
      return;
    }

    if (selectedValue === 'monitorLogs') {
      this.showingDomains = false;
      this.showingPubkey = false;
      this.showingChart = false;
      this.showingExport = false;
      this.showingTerminateVerification = false;
      this.monitorLogs = true;
      return;
    }
  }

  sshPubkeyToClipboard(key: string) {
    navigator.clipboard.writeText(key);
    this.temporaryCopyAnimationShowing = true;
    setTimeout(() => {
      this.temporaryCopyAnimationShowing = false;
    }, 1000);
  }

  generateLogArgs(type: string) {
    if (type === 'access') {
      return '-n|20|/var/log/httpd/' + this.domain + '_access.log';
    }
    if (type === 'error') {
      return '-n|20|/var/log/httpd/' + this.domain + '_error.log';
    }
    return 'INVALID FILE TYPE';
  }

  unlockSite() {
    this.router.navigate(['process', 'unlock'], { relativeTo: this.route });
  }

  lockSite() {
    this.router.navigate(['process', 'lock'], { relativeTo: this.route });
  }

  cloneAccount() {
    this.router.navigate(['..', 'clone', this.domain], {
      relativeTo: this.route,
    });
  }

  exportAccount() {
    this.router.navigate(['..', 'export', this.domain], {
      relativeTo: this.route,
    });
  }

  deleteAlternateDomain(domain: string) {
    console.log('delete: ' + domain);
  }

  setAsPrimaryDomain(domain: string) {
    console.log('set to primary: ' + domain);
  }

  toggleIsAddingDomain() {
    this.isAddingDomain = !this.isAddingDomain;
  }

  addDomain() {
    console.log('adding ' + this.newDomain);
    this.addDomainNow = true;
  }

  onFinishedAddingDomain(e: boolean): void {
    this.isAddingDomain = false;
    this.newDomain = '';
    window.location.reload();
    // this.addDomainNow = false;
  }

  generateAddDomainENV(): string {
    return `ACCOUNT=${this.domain}|DOMAIN=${this.newDomain}`;
  }

  generateExportENV(): string {
    return `ACCOUNT_NAME=${this.domain}`; // TODO: update "domain" to "account" pretty much everywhere
  }

  created(ts: string): string {
    return new Date(ts).toDateString();
  }
}
