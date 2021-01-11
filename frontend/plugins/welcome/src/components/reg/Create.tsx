import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme, ContentHeader, } from '@backstage/core';

import {

  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
 
  Link,
  Button,
} from '@material-ui/core';
import Timer from '../Timer';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command


import { Alert } from '@material-ui/lab';




import { Link as RouterLink } from 'react-router-dom';

// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
    display: 'flex',
    justifyContent: 'center',
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    margin: theme.spacing(3),
    width: 320,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: '25ch',
  },


  margin: {
    margin: theme.spacing(3),
  },
  input: {
    display: 'none',
  },

}));

interface Patientrights_Type {
  /**
     * 
     * @type {number}
     * @memberof ControllersPatientrights
     */
    insurance?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrights
     */
    medicalrecordstaff?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrights
     */
    patientrecord?: number;
    /**
     * Abilitypatientrights int
     * @type {number}
     * @memberof ControllersPatientrights
     */
    patientrightstype?: number;
}

const NewPatientright: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'สิทธ์' };
  const http = new DefaultApi();

  const [Patientrights, setPatientrights] =             React.useState<Partial<Patientrights_Type>>({});

 
 const [status, setStatus] = React.useState(false);
 const [alert, setAlert] = React.useState(true);


  

  // Lifecycle Hooks
  useEffect(() => {
   
  }, []);

  // set data to object Patientright
  const handleChange = (

    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof NewPatientright;
    const { value } = event.target;
    setPatientrights({ ...Patientrights, [name]: value });
  };


 


  const CreatePatientright = async () => {
   
    const res: any = await http.createPatientrights({ 
      patientrights:Patientrights
    
    
    });
    console.log(Patientrights);
    setStatus(true);
  
    
    if (res.id != ''){
      setAlert(true);
    } else {
      setAlert(false);
    }

    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
 
  
  };
  


 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ลงทะเบียน ${profile.givenName || 'to Backstage'}`}
       subtitle="Some quick intro and links."
     ><Timer /></Header>
     <Content>
       <ContentHeader title="ข้อมูล">
         
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 บันทึกสำเร็จ
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>

       <div>
              <br/><p>json {JSON.stringify(Patientrights)} </p>
              
              
          </div>

            
      











       <div className={classes.root}>
         <form noValidate autoComplete="off">
           
         
           <div className={classes.margin}>
             <Button
               onClick={() => {
                CreatePatientright();
               }}
               variant="contained"
               color="primary"
             >
               Submit
             </Button>
              
         <Link component={RouterLink} to="/">
           <Button variant="contained" color="primary">
           กลับสู่หน้าหลัก
           </Button>
         </Link>

           </div>
         </form>
       </div>
  
     </Content>
   </Page>
 );
};
export default NewPatientright;