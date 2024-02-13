import { Component } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { ROUTER_LINNKS } from '../../data/global-constants';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [MatCardModule, RouterModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
})
export class HomeComponent {
  public get routerLinks() {
    return ROUTER_LINNKS;
  }
}
