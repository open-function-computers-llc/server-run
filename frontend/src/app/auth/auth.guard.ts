import { Injectable } from "@angular/core";
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from "@angular/router";
import { first, map, Observable } from "rxjs";
import { AuthService } from "./auth.service";

@Injectable()
export class AuthGuard implements CanActivate {
    constructor(private authService: AuthService, private router: Router) {}

    canActivate(route: ActivatedRouteSnapshot, router: RouterStateSnapshot):
        boolean |
        UrlTree |
        Promise<boolean | UrlTree> |
        Observable<boolean | UrlTree>
    {
        return this.authService.user.pipe(
            first(),
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
