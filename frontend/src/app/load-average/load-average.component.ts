import { Component, OnInit } from '@angular/core';
import { Observable, share } from 'rxjs';
import { ServerService } from '../server.service';
import { SystemLoad } from '../SystemLoad';

@Component({
  selector: 'app-load-average',
  templateUrl: './load-average.component.html',
  styleUrls: ['./load-average.component.scss']
})
export class LoadAverageComponent implements OnInit {
  systemLoad : Observable<SystemLoad>;

  constructor(
    private serverService: ServerService,
  ) { }

  ngOnInit(): void {
    this.systemLoad = this.serverService.streamSystemLoad();
  }

}
