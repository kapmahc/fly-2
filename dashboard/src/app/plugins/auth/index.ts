import {UsersSignInComponent} from './users/sign-in.component';
import {UsersSignUpComponent} from './users/sign-up.component';
import {UsersConfirmComponent} from './users/confirm.component';
import {UsersUnlockComponent} from './users/unlock.component';
import {UsersForgotPasswordComponent} from './users/forgot-password.component';

export default {
  routes: [
    { path: 'users/sign-in', component: UsersSignInComponent },
    { path: 'users/sign-up', component: UsersSignUpComponent },
    { path: 'users/confirm', component: UsersConfirmComponent },
    { path: 'users/unlock', component: UsersUnlockComponent },
    { path: 'users/forgot-password', component: UsersForgotPasswordComponent }
  ]
}
