export class User {
    constructor(
        private _token: string,
        private _expiresAt: Date,
    ) {}

    get token() {
        if (!this._expiresAt || new Date() > this._expiresAt) {
            return "";
        }

        return this._token;
    }
}
