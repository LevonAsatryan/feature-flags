import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  public get isLoggedIn(): boolean {
    return this._isLoggedIn;
  }

  /**
   * This is a mock, delete later
   */
  private readonly _isLoggedIn: boolean = true;
}
