import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { catchError, Observable, take } from "rxjs";
import { ServerService } from "../server.service";
import { BootstrapUptimeResponse } from "./BootstrapUptimeResponse";
import { UptimeResponse } from "./UptimeResponse";

@Injectable()
export class UptimeService {
    uptimeMonitoringIsAvailable:boolean = false;

    constructor(
        private http: HttpClient,
        private serverService: ServerService,
    ) {}

    bootstrapUptimeService() {
        this.http.get<BootstrapUptimeResponse>("/api/uptime-provider", {headers: this.serverService.getHeaders()}).pipe(
            take(1),
        ).subscribe({
            next: (v) => {
                this.uptimeMonitoringIsAvailable = v.uptimeAvailable;
            },
            error: (e) => { alert("Nope!") }
        });
    }

    getUptimeFor(uri:string) : Observable<UptimeResponse> {
        return this.http.get<UptimeResponse>("/api/uptime?uri="+uri, {headers: this.serverService.getHeaders()})
    }

    setUptimeURIForDomain(domain: string, uri: string) : void {
        let body = new FormData();
        body.append("domain", domain);
        body.append("uri", uri);
        body.append("action", "update-uptime-uri");

        this.http.post("/api/update", body, {headers: this.serverService.getHeaders()}).pipe(
            take(1),
        ).subscribe({
            next: (v) => {
                console.log(v)
            },
            error: (err) => {
                console.log("ERROR! -> ", err)
            }
        });
    }
}
