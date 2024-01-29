import { Component, EventEmitter, Input, Output } from '@angular/core';
import { MatSidenavModule } from '@angular/material/sidenav';
@Component({
  selector: 'side-bar-menu',
  standalone: true,
  imports: [MatSidenavModule],
  templateUrl: './side-bar-menu.component.html',
  styleUrl: './side-bar-menu.component.scss',
})
export class SideBarMenuComponent {
  @Input() shouldBeOpen: boolean = false;
  @Output() close$ = new EventEmitter();

  public close(): void {
    this.close$.emit();
  }
}
