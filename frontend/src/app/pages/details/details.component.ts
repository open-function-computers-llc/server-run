import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Observable, share } from 'rxjs';
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
    // return this.sanitizer.bypassSecurityTrustResourceUrl("/api/analytics?domain=" + this.domain + "&period=" + this.analyticsView + "&token=" + token);
  }

  setAnalyticView(e:any) {
    console.log(e.target.value);
    this.analyticsView = e.target.value;
    this.setAnalyticsURL();
  }

  unlockSite() {
    console.log("unlock!")
    this.router.navigate(['process', 'unlock'], { relativeTo: this.route });
  }

  lockSite() {
    console.log("lock!")
    this.router.navigate(['process', 'lock'], { relativeTo: this.route });
  }

}
