import { Component, OnInit, Input, OnDestroy } from '@angular/core';
import * as Highcharts from 'highcharts';
import { Observable, Subscription } from 'rxjs';
import { ServerService } from '../server.service';

@Component({
  selector: 'app-chart',
  templateUrl: './chart.component.html',
  styleUrls: ['./chart.component.scss']
})
export class ChartComponent implements OnInit, OnDestroy {
  Highcharts: typeof Highcharts = Highcharts;
  chartSubscription: Subscription;
  chartOptions: Highcharts.Options = {
    series: [{
      data: [1, 2, 3, 4, 5],
      type: 'bar'
    }]
  };
  @Input()
  account: string = "default";
  type: string = "visitors";

  constructor(private serverService: ServerService) { }

  ngOnInit(): void {
    this.chartSubscription = this.serverService.getAccountAnalyticData(this.account, this.type).subscribe((data:Highcharts.Options) => {
      this.chartOptions = data;
    });
  }

  ngOnDestroy(): void {
    this.chartSubscription.unsubscribe();
  }

}
