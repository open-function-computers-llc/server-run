export interface CertInfo {
    valid: boolean;
    expires: Date;
    error: string;
    names: string[];
}
export interface DomainInfo {
    status: string[];
    registered: Date;
    expires: Date;
    registrant: string;
    registrantEmail: string;
    registrar: string;
    error: string;
}
export interface Outage {
    start: Date;
    end: Date;
    duration: number;
}
export interface Uptime {
    days1: number;
    days30: number;
    days60: number;
    days7: number;
    days90: number;
}
export interface UptimeResponse {
    certInfo: CertInfo;
    domainInfo: DomainInfo;
    outages: Outage[];
    up: boolean;
    uptime: Uptime;
    url: string;
}
