import {UsersSignInComponent} from './users/sign-in.component';
import {UsersSignUpComponent} from './users/sign-up.component';
import {UsersConfirmComponent} from './users/confirm.component';
import {UsersUnlockComponent} from './users/unlock.component';
import {UsersForgotPasswordComponent} from './users/forgot-password.component';

export default {
  components: [
    UsersSignInComponent,
    UsersSignUpComponent,
    UsersConfirmComponent,
    UsersUnlockComponent,
    UsersForgotPasswordComponent
  ],
  routes: [
    {
      path: 'users',
      children: [
        { path: 'sign-in', component: UsersSignInComponent },
        { path: 'sign-up', component: UsersSignUpComponent },
        { path: 'confirm', component: UsersConfirmComponent },
        { path: 'unlock', component: UsersUnlockComponent },
        { path: 'forgot-password', component: UsersForgotPasswordComponent }
      ]
    },
  ]
}
