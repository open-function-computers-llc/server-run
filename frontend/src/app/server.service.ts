import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { SystemLoad } from './SystemLoad';
import { interval, Observable, of } from 'rxjs';
import { catchError, map, mergeMap, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ServerService {

  constructor(
    private http: HttpClient,
  ) { }

  systemLoadURL = "/api/system-load";

  getSystemLoad() : Observable<SystemLoad> {
    return interval(5000).pipe(
             mergeMap(() => {
                return this.http.get<SystemLoad>(this.systemLoadURL).
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
