import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Subject } from "rxjs";
import { take } from "rxjs/operators";
import { AuthResponse } from "./AuthResponse";
import { User } from "./user.model";

@Injectable({ providedIn: 'root' })
export class AuthService {
    user = new Subject<User|null>();

    constructor(
        private http: HttpClient,
        private router: Router,
    ) {}

    login(username: string, password: string) {
        let body = new FormData();
        body.append("user", username);
        body.append("pass", password);

        this.http.post<AuthResponse>("/api/auth", body, {
            headers: {},
        }).pipe(
            take(1),
        ).subscribe({
            next: (v) => {
                const user = new User(v.authToken, new Date(v.expiresAt));
                this.user.next(user);
                // TODO: localStorage set here
                localStorage.setItem("ofco-auth", JSON.stringify(v));

                this.router.navigate(["/system", "home"]);
            },
            error: (e) => { alert("Nope!") }
        });
    }

}
