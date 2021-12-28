import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Observable, of, share } from 'rxjs';
import { ServerService } from 'src/app/server.service';
import { Website } from 'src/app/Website';

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss']
})
export class DetailsComponent implements OnInit {
  site$: Observable<Website|null> = of(null);

  constructor(
    private serverService: ServerService,
    private route: ActivatedRoute,
  ) { }

  ngOnInit(): void {
    const domain:string = this.route.snapshot.paramMap.get("domain") || "";
    this.site$ = this.serverService.getSiteDetails(domain).pipe(share());
  }

}
