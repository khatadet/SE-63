import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table_Patientrightstype';
import Button from '@material-ui/core/Button';
import Timer from '../Timer';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ระบบ สัญญาเช่า' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ยินดีต้อนรับ เข้าสู่ ${profile.givenName || 'to Backstage'}`}
       subtitle="ของ บริษัท Blue moon."
      >      <Timer /></Header>
           <Content>



       <ContentHeader title="สัญญาเช่า">
         
         <Link component={RouterLink} to="">
           <Button variant="contained" color="primary">
           กลับสู่หน้าหลัก
           </Button>
         </Link>
         
         <Link component={RouterLink} to="/create">
           <Button variant="contained" color="primary">
             เพิ่มผู้เช่า
           </Button>
         </Link>


       </ContentHeader>
       
       <ComponanceTable></ComponanceTable>
        
       
     </Content>
   </Page>
 );
};
 
export default WelcomePage;