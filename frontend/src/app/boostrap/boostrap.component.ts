import { Component, OnInit } from '@angular/core';
import { UptimeService } from '../uptime/uptime.service';

@Component({
  selector: 'app-boostrap',
  templateUrl: './boostrap.component.html',
  styleUrls: ['./boostrap.component.scss']
})
export class BoostrapComponent implements OnInit {

  constructor(
    private uptimeService: UptimeService,
  ) { }

  ngOnInit(): void {
    this.uptimeService.bootstrapUptimeService();
  }

}
