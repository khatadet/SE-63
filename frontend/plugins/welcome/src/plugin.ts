import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';

import  create from './components/create_Patientrights';
import Login from './components/Login';
import  pdf from './components/pdf';
import  TablePatientrights from './components/TablePatientrights';
import  create_Patientrightstype from './components/TablePatientrights';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/Log', Login);

    router.registerRoute('/create', create);

    router.registerRoute('/pdf', pdf);
    router.registerRoute('/TablePatientrights', TablePatientrights);
    router.registerRoute('/acc', create_Patientrightstype);



  },
});
