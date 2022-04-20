import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DetailsComponent } from './components/details/details.component';
import { HomeComponent } from './components/home/home.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';

const routes: Routes = [
{
    path: '',
   component: LoginComponent,
  },
  {
    path: 'login',
   component: LoginComponent,
  },
  {
    path: 'home',
   component: HomeComponent,
  },
  {
    path: 'register',
    component: RegisterComponent,
  },
  
  {
    path: 'search/:item-search',
    component: HomeComponent
  },
  {
    path: 'details/:id',
    component: DetailsComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
