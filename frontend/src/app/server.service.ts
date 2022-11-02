import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { SystemLoad } from './SystemLoad';
import { interval, Observable, of } from 'rxjs';
import { catchError, mergeMap } from 'rxjs/operators';
import { Website } from './Website';
import { webSocket } from "rxjs/webSocket";
import { ScriptMessage } from './ScriptMessage';
import { AuthService } from './auth/auth.service';
import { Options } from 'highcharts';

@Injectable({
  providedIn: 'root'
})
export class ServerService {

  constructor(
    private http: HttpClient,
    private authService: AuthService,
  ) { }

  getHeaders() {
    const localInfo = JSON.parse(localStorage.getItem("ofco-auth") || "");
    const token = localInfo.authToken || "";
    return new HttpHeaders().set("Authorization", token);
  }

  getToken() : string {
    const localInfo = JSON.parse(localStorage.getItem("ofco-auth") || "");
    return localInfo.authToken || "";
  }

  streamSystemLoad() : Observable<SystemLoad> {
    const base = window.location.href.split('/').slice(0, 3).join('/') + "/"
    const subject = webSocket<SystemLoad>(base.replace("http", "ws") + "stream/system-load?token="+this.getToken());
    return subject;
  }

  streamScriptProcess(script: string, arg?: string, env?: string) : Observable<ScriptMessage> {
    // build script path
    let path = window.location.href.split('/').slice(0, 3).join('/') + "/";
    path = path.replace("http", "ws");
    path = path+"stream/script?token=" + this.getToken();
    path = path + "&script=" + script;
    if (!!env) {
      path = path + "&env="+env;
    } else {
      path = path + "&arg="+arg;
    }
    return webSocket<ScriptMessage>(path);
  }

  getSites() : Observable<Website[]> {
    return this.http.get<Website[]>("/api/sites", {headers: this.getHeaders()}).
      pipe(
        catchError(this.handleError<Website[]>('get-sites', []))
      )
  }

  getSiteDetails(d:string) : Observable<Website> {
    return this.http.get<Website>("/api/details?domain="+d, {headers: this.getHeaders()}).
      pipe(
        catchError(this.handleError<Website>('get-details', {
          account: d,
          isLocked: false,
          alternateDomains: [],
          domain: "",
          uptimeURI: "",
          sshPubKey: "",
        }))
      )
  }

  getAccountAnalyticData(a:string, t: string) : Observable<Options> {
    return this.http.get<Options>("/api/analytics-json?account="+a+"&type="+t, {headers: this.getHeaders()}).
      pipe(
        catchError(this.handleError<Options>('get-analytic-json', {
          series: [
            {
              data: [1, 2, 3],
              type: "line"
            }
          ]
        }))
      )
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.log(operation);
      console.error(error); // log to console instead
      this.authService.logout();

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
