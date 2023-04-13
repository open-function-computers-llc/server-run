import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-fail2ban-unlock',
  templateUrl: './fail2ban-unlock.component.html',
  styleUrls: ['./fail2ban-unlock.component.scss'],
})
export class Fail2banUnlockComponent implements OnInit {
  ipAddress: string = '';
  isUnbanning: boolean = false;

  constructor() {}

  ngOnInit(): void {}

  unban() {
    this.isUnbanning = true;

    setTimeout(() => {
      this.isUnbanning = false;
      this.ipAddress = '';
    }, 2500);
  }

  generateUnbanENV() {
    return 'IPADDRESS=' + this.ipAddress;
  }
}
