import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { Observable, Subscription, take } from 'rxjs';
import { ScriptMessage } from 'src/app/ScriptMessage';
import { ServerService } from 'src/app/server.service';

@Component({
  selector: 'app-process',
  templateUrl: './process.component.html',
  styleUrls: ['./process.component.scss']
})
export class ProcessComponent implements OnInit, OnDestroy {
  scriptOutput: Subscription;
  domain: string;
  action: string;
  messages: string[] = [];
  isComplete: boolean = false;

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private serverService: ServerService,
  ) { }

  ngOnInit(): void {
    this.action = this.route.snapshot.paramMap.get("action") || "";
    this.domain = this.route.snapshot.paramMap.get("domain") || "";
    this.scriptOutput = this.serverService.streamScriptProcess(this.action, this.domain).subscribe(
      (o) => {
        this.messages.push(o.output);
      },
      (e) => {
        // TODO: web socket error handling
      },
      () => {
        this.isComplete = true;
      }
    );
  }

  ngOnDestroy(): void {
    this.scriptOutput.unsubscribe();
  }

  return() {
    this.router.navigate(["../.."], { relativeTo: this.route })
  }
}
