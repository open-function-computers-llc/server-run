import { Component, Input, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { UptimeService } from './uptime.service';
import { UptimeResponse } from './UptimeResponse';

@Component({
  selector: 'app-uptime',
  templateUrl: './uptime.component.html',
  styleUrls: ['./uptime.component.scss']
})
export class UptimeComponent implements OnInit {
  @Input()
  uptimeURI: string = "";
  @Input()
  domain: string = "";

  originalUptimeURI: string = "";
  uptimeAvailable: boolean = false;
  uptimeInfo: Observable<UptimeResponse>;
  isUpdatingURI: boolean = false;

  constructor(
    private uptimeService: UptimeService,
  ) { }

  ngOnInit(): void {
    this.originalUptimeURI = this.uptimeURI;

    this.uptimeAvailable = this.uptimeService.uptimeMonitoringIsAvailable;
    if (this.uptimeAvailable) {
      this.uptimeInfo = this.uptimeService.getUptimeFor(this.uptimeURI);
    }
  }

  toggleForm(): void {
    if (this.uptimeURI !== this.originalUptimeURI) {
      this.uptimeURI = this.originalUptimeURI; // in case we hit cancel
    }
    this.uptimeInfo = this.uptimeService.getUptimeFor(this.uptimeURI);

    this.isUpdatingURI = !this.isUpdatingURI;
  }

  updateUptimeURI(): void {
    this.isUpdatingURI = false;
    this.uptimeInfo = this.uptimeService.getUptimeFor(this.uptimeURI);
    this.uptimeService.setUptimeURIForDomain(this.domain, this.uptimeURI);
  }
}
