import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { SystemLoad } from './SystemLoad';
import { interval, Observable, of } from 'rxjs';
import { catchError, map, mergeMap, tap } from 'rxjs/operators';
import { Website } from './Website';

@Injectable({
  providedIn: 'root'
})
export class ServerService {

  constructor(
    private http: HttpClient,
  ) { }

  getSystemLoad() : Observable<SystemLoad> {
    return interval(5000).pipe(
             mergeMap(() => {
                return this.http.get<SystemLoad>("/api/system-load").
                  pipe(
                    catchError(this.handleError<SystemLoad>('get-system-load', {
                      oneMinute: "error",
                      fiveMinutes: "error",
                      fifteenMinutes: "error",
                    }))
                  )
              })
    );
  }

  getSites() : Observable<Website[]> {
    return this.http.get<Website[]>("/api/sites").
      pipe(
        catchError(this.handleError<Website[]>('get-sites', []))
      )
  }

  getSiteDetails(d:string) : Observable<Website> {
    return this.http.get<Website>("/api/details?domain="+d).
      pipe(
        catchError(this.handleError<Website>('get-details', {
          domain: d,
          isLocked: false,
          alternateDomains: [],
          uptimeURI: "",
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

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
