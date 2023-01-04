import { Component, ElementRef, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { ServerService } from '../server.service';

@Component({
  selector: 'app-script-viewer',
  templateUrl: './script-viewer.component.html',
  styleUrls: ['./script-viewer.component.scss']
})
export class ScriptViewerComponent implements OnInit {
  @Input('script-name') scriptName: string = "default"; // TODO: this isn't a valid default value... do something with that
  @Input('script-arg') scriptArg: string = "";
  @Input('script-env') scriptEnv: string = "";
  @Input('line-action') lineAction: string = "";
  @Input('line-action-label') lineActionLabel: string = "";
  @Input('line-method') lineMethod: string = "";
  @Output() isCompleted: EventEmitter<boolean> = new EventEmitter<boolean>();

  scriptOutput: Subscription;
  messages: string[] = [];
  isComplete: boolean = false;

  constructor(
    private serverService: ServerService,
    private domSanitizer: DomSanitizer,
    private elReference: ElementRef,
    private router: Router,
  ) {
  }

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
        this.bindButtons();
      }
    );
  }

  buttonClicked(e:any) {
    const lineText: string = e.target.nextElementSibling.textContent;
    // this code kinda sucks, compare string to call the appropriate method.
    // see the available methods at the end of this class
    if (this.lineMethod === "importAccount") {
      this.triggerImport(lineText);
      return;
    }

    // invalid default
    console.log("No bound click handler!", lineText);
  }

  bindButtons() {
    const buttons = this.elReference.nativeElement.querySelectorAll("button");
    console.log(buttons);
    if (!buttons) {
      return;
    }
    for (let index = 0; index < buttons.length; index++) {
      const b = buttons[index];
      b.disabled = false;
      b.addEventListener('click', this.buttonClicked.bind(this));
      console.log(b);
    }
  }




  // Action Methods... not sure this should live here but for now it's good enough
  triggerImport(filename: string) {
    console.log("import trigger!", filename);
    this.router.navigate(["/system/accounts/import/", filename]);
  }
}
