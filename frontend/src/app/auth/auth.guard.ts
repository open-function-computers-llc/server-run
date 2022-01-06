import { Injectable } from "@angular/core";
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from "@angular/router";
import { map, Observable, take } from "rxjs";
import { AuthService } from "./auth.service";

@Injectable({ providedIn: 'root' })
export class AuthGuard implements CanActivate {
    constructor(private authService: AuthService, private router: Router) {}

    canActivate(route: ActivatedRouteSnapshot, router: RouterStateSnapshot):
        boolean |
        UrlTree |
        Promise<boolean | UrlTree> |
        Observable<boolean | UrlTree>
    {
        return true;
        return this.authService.user.pipe(
            take(1),
            map((u) => {
                const isAuthenticated = !!u;
                if (isAuthenticated) {
                    return true;
                }
                return this.router.createUrlTree(['/login']);
            }
        ));
    }
}
