import { Component, OnInit, Input, OnDestroy, OnChanges, SimpleChanges } from '@angular/core';
import * as Highcharts from 'highcharts';
import { Observable, Subscription } from 'rxjs';
import { ServerService } from '../server.service';

@Component({
  selector: 'app-chart',
  templateUrl: './chart.component.html',
  styleUrls: ['./chart.component.scss']
})
export class ChartComponent implements OnInit, OnDestroy, OnChanges {
  Highcharts: typeof Highcharts = Highcharts;
  chartSubscription: Subscription;
  defaultChart: Highcharts.Options = {
    yAxis: {
      title: {
        text: ""
      },
    },
    title: {
      text: ""
    },
    legend: {
      enabled: false
    },
    series: [{
      data: [],
      type: 'column'
    }]
  };
  chartOptions: Highcharts.Options = this.defaultChart;
  @Input()
  account: string = "default";
  @Input()
  type: string = "visitors";

  constructor(private serverService: ServerService) {
    Highcharts.setOptions({
      lang: {
        thousandsSep: ","
      }
    });
  }

  ngOnInit(): void {
    this.chartOptions = this.defaultChart;

    this.getChartData();
  }

  ngOnChanges(changes: SimpleChanges): void {
    this.chartOptions = this.defaultChart;

    if (this.chartSubscription) {
      this.chartSubscription.unsubscribe();
    }

    this.getChartData();
  }

  getChartData(): void {
    this.chartSubscription = this.serverService.getAccountAnalyticData(this.account, this.type).subscribe((data: Highcharts.Options) => {
      this.chartOptions = data;
    });
  }

  ngOnDestroy(): void {
    this.chartSubscription.unsubscribe();
  }

}
