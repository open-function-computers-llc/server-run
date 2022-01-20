import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { BehaviorSubject, Subject } from "rxjs";
import { take } from "rxjs/operators";
import { AuthResponse } from "./AuthResponse";
import { User } from "./user.model";

@Injectable()
export class AuthService {
    user = new BehaviorSubject<User|null>(null);

    constructor(
        private http: HttpClient,
        private router: Router,
    ) {}

    login(username: string, password: string) {
        let body = new FormData();
        body.append("user", username);
        body.append("pass", password);

        this.http.post<AuthResponse>("/api/auth", body).pipe(
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

    autoLogin() {
        const localInfo = JSON.parse(localStorage.getItem("ofco-auth") || "");
        const token = localInfo.authToken || "";
        const expiresAt = new Date(localInfo.expiresAt);

        const user = new User(token, expiresAt);
        if (user.token) {
            this.user.next(user);
        }
    }

    logout() {
        this.user.next(null);
        localStorage.removeItem("ofco-auth");
    }

}
