export interface Website {
    isLocked: boolean;
    account: string;
    alternateDomains: string[];
    domain: string;
    uptimeURI: string;
    sshPubKey: string;
}
