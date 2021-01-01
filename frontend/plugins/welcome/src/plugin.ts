import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';

import  create from './components/create';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);

    router.registerRoute('/create', create);


  },
});
