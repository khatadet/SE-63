import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';

import  create from './components/create_Patientrights';
import Login from './components/Login';
//import  create_Abilitypatientrights from './components/create_Abilitypatientrights';
//import  create_Insurance from './components/create_Insurance';
//import  create_Patientrightstype from './components/create_Patientrightstype';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/Log', Login);

    router.registerRoute('/create', create);

    //router.registerRoute('/create_Abilitypatientrights', create_Abilitypatientrights);
    //router.registerRoute('/create_Insurance', create_Insurance);
    //router.registerRoute('/create_Patientrightstype', create_Patientrightstype);



  },
});
