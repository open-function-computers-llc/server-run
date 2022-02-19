import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Subscription } from 'rxjs';
import { ServerService } from '../server.service';

@Component({
  selector: 'app-script-viewer',
  templateUrl: './script-viewer.component.html',
  styleUrls: ['./script-viewer.component.scss']
})
export class ScriptViewerComponent implements OnInit {
  @Input('script-name') scriptName: string;
  @Input('script-arg') scriptArg: string = "";
  @Input('script-env') scriptEnv: string = "";
  @Output() isCompleted: EventEmitter<boolean> = new EventEmitter<boolean>();

  scriptOutput: Subscription;
  messages: string[] = [];
  isComplete: boolean = false;

  constructor(
    private serverService: ServerService,
  ) { }

  ngOnInit(): void {
    console.log(this.scriptName, this.scriptEnv, this.scriptArg);
    const arg = this.scriptArg.split(" ").join("-");
    this.scriptOutput = this.serverService.streamScriptProcess(this.scriptName, arg, this.scriptEnv).subscribe(
      (o) => {
        this.messages.push(o.output);
      },
      (e) => {
        // TODO: web socket error handling
      },
      () => {
        this.isComplete = true;
        this.isCompleted.emit(true);
      }
    );
  }

}
