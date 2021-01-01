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

//import { ControllersPatientrights } from '../../api/models/ControllersPatientrights';
import { EntPatientrightstype } from '../../api/models/EntPatientrightstype';
import { EntInsurance } from '../../api/models/EntInsurance';
import { EntMedicalrecordstaff } from '../../api/models/EntMedicalrecordstaff';
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
  abilitypatientrights?: number;
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
  patientrecord?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersPatientrights
   */
  patientrightstype?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersPatientrights
   */
  permissionDate?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersPatientrights
   */
  user?: number;
}

const NewPatientright: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'สิทธ์' };
  const http = new DefaultApi();

  const [Patientrights, setPatientrights] =            React.useState<Partial<Patientrights_Type>>({});

  const [Patientrightstype, setPatientrightstype] =     React.useState<EntPatientrightstype[]>([]);
 
 const [Insurance, setInsurance] =  React.useState<EntInsurance[]>([]);
 const [Medicalrecordstaff, setMedicalrecordstaff] =            React.useState<EntMedicalrecordstaff[]>([]);
 const [status, setStatus] = React.useState(false);
 const [alert, setAlert] = React.useState(true);


  const getMedicalrecordstaffs = async () => {
    const res = await http.listMedicalrecordstaff({ limit: 10, offset: 0 });
    setMedicalrecordstaff(res);
  };

  const getPatientrightstype = async () => {
    const res = await http.listPatientrightstype({ limit: 10, offset: 0 });
    setPatientrightstype(res);
  };

  const getInsurance = async () => {
    const res = await http.listInsurance({ limit: 10, offset: 0 });
    setInsurance(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getMedicalrecordstaffs();
    getPatientrightstype();
    //getRoom();
    getInsurance();
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
       <ContentHeader title="เช่าห้องพักอาศัย">
         
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 สร้างสัญญาสำเร็จ
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>

        {/*
       <div className={classes.root}>

         
         <form noValidate autoComplete="off">
         
         <FormControl variant="filled" className={classes.formControl}>
          <TextField
               name="rentDate"
               label="PatientrightDate"
               
               type="date"
               size="medium"
               
               value={Patientrights.rentDate}
               onChange={handleChange}

               InputLabelProps={{
                shrink: true,
              }}

             />

          </FormControl>
            


         </form>
       </div>
            */}

            
       <div className={classes.root}>

         
         <form noValidate autoComplete="off">

         
         
         <FormControl variant="filled" className={classes.formControl}>
          <TextField
               name="rentAge"
               label="ระยะเวลาสัญญา"
               variant="outlined"
               type="string"
               size="medium"
               
               value={Patientrights.rentAge}
               onChange={handleChange}
             />

          </FormControl>
          
          <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>หน่วยระยะเวลาสัญญา</InputLabel>
                <Select
                  name="roomage"
                  value={Patientrights.roomage}
                  onChange={handleChange}
                >
                  {Patientrights.map((item:any) => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.text}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>


        </form>
       </div>
                
       <div className={classes.root}>
         <form noValidate autoComplete="off">

         <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ห้อง</InputLabel>
                <Select
                  name="room"
                  value={Patientrights.room}
                  onChange={handleChange}
                >
                  {Room.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.roomName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
         
 
 
        
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ราคามัดจำ</InputLabel>
                <Select
                  name="insurance"
                  value={Patientrights.insurance}
                  onChange={handleChange}
                >
                  {Insurance.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.insurance}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>

           </form>
       </div>











        <ContentHeader title="ข้อมูลผู้เช่า"/>
         




        <div className={classes.root}>
         <form noValidate autoComplete="off">

         <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกพนักงาน</InputLabel>
                <Select
                 name="user"
                  value={Patientrights.user}
                  onChange={handleChange}
                >
                  {Medicalrecordstaff.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.id}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
        
         
         

        


        </form>
        
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