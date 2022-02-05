import { Component, Input, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { ServerService } from '../server.service';

@Component({
  selector: 'app-script-viewer',
  templateUrl: './script-viewer.component.html',
  styleUrls: ['./script-viewer.component.scss']
})
export class ScriptViewerComponent implements OnInit {
  @Input('script-name') scriptName: string;
  scriptOutput: Subscription;
  messages: string[] = [];
  isComplete: boolean = false;

  constructor(
    private serverService: ServerService,
  ) { }

  ngOnInit(): void {
    console.log(this.scriptName);
    this.scriptOutput = this.serverService.streamScriptProcess(this.scriptName, "").subscribe(
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

}
