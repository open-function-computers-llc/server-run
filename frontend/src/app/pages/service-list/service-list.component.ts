import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { ServerService } from 'src/app/server.service';
import { SystemService } from 'src/app/SystemService';

@Component({
  selector: 'app-service-list',
  templateUrl: './service-list.component.html',
  styleUrls: ['./service-list.component.scss']
})
export class ServiceListComponent implements OnInit {
  services: Observable<SystemService[]>;

  constructor(private serverService: ServerService) { }

  ngOnInit(): void {
    this.services = this.serverService.getSystemServiceStatuses();
  }

}
