import { Component, ElementRef, HostListener, ViewChild } from '@angular/core';
import { MatToolbar, MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { SideBarMenuComponent } from './side-bar-menu/side-bar-menu.component';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [
    MatToolbarModule,
    MatButtonModule,
    MatIconModule,
    SideBarMenuComponent,
  ],
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss',
})
export class HeaderComponent {
  public shouldBeOpen: boolean = false;

  @ViewChild('header', { static: true })
  public header: ElementRef | null = null;

  private prevScrollpos = window.scrollY;

  private readonly toolbarHeight = 64;

  public toggleMenu(): void {
    this.shouldBeOpen = !this.shouldBeOpen;
  }

  @HostListener('window:scroll', ['$event'])
  public onWindowScroll(event: Event): void {
    if (!this.header) return;
    const currentScrollPos = window.scrollY;
    if (this.prevScrollpos > currentScrollPos) {
      this.header.nativeElement.style.top = '0';
    } else {
      this.header.nativeElement.style.top = `-${this.toolbarHeight}px`;
    }
    this.prevScrollpos = currentScrollPos;
  }
}
