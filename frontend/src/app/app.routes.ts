import { Routes } from '@angular/router';
import { PublicLayoutComponent } from './layout/public-layout/public-layout.component';
import { HomePageComponent } from './modules/home/home-page/home-page.component';
import { GamePageComponent } from './modules/game/game-page/game-page.component';

export const routes: Routes = [
    {
        path: '',
        component: PublicLayoutComponent,
        children: [
          {
            path: '',
            component: HomePageComponent
          },
          {
            path: 'game/:id',
            component: GamePageComponent
          },
        ]
      },
];
