import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Observable } from 'rxjs';
import { ServerService } from 'src/app/server.service';
import { SystemService } from 'src/app/SystemService';

@Component({
  selector: 'app-service-details',
  templateUrl: './service-details.component.html',
  styleUrls: ['./service-details.component.scss']
})
export class ServiceDetailsComponent implements OnInit {
  service: Observable<SystemService>;
  serviceName: String;

  constructor(
    private serverService: ServerService,
    private route: ActivatedRoute,
    private router: Router,
  ) { }

  ngOnInit(): void {
    const s:string = this.route.snapshot.paramMap.get("name") || "";
    this.serviceName = s;
    this.service = this.serverService.getServiceDetails(this.serviceName);
  }

  restartService() {
    this.serverService.restartService(this.serviceName);
    this.router.navigate([this.router.url])
  }
}
