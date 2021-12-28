import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { ServerService } from 'src/app/server.service';
import { Website } from 'src/app/Website';

@Component({
  selector: 'app-sites',
  templateUrl: './sites.component.html',
  styleUrls: ['./sites.component.scss']
})
export class SitesComponent implements OnInit {
  sites$!: Observable<Website[]>;

  constructor(private serverService: ServerService) { }

  ngOnInit(): void {
    this.sites$ = this.serverService.getSites();
  }

}
