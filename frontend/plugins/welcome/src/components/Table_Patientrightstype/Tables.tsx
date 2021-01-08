import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntPatientrightstype } from '../../api/models/EntPatientrightstype';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 


export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [Patientrightstypes, setPatientrightstypes] = useState<EntPatientrightstype[]>(Array);




 const [loading, setLoading] = useState(true);
 
 const getPatientrightstypes = async () => {
    const res = await api.listPatientrightstype({ limit: 10, offset: 0 });
  
  setLoading(false);
  setPatientrightstypes(res);
};

//console.log(Patientrightss)


 useEffect(() => {
   
  getPatientrightstypes();



   
 }, [loading]);
 
 
 const deletePatientrightstypes = async (id: number) => {
  const res = await api.deletePatientrightstype({ id: id });
  setLoading(true);
};

 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">permission</TableCell>
           <TableCell align="center">permissionArea</TableCell>
           <TableCell align="center">responsible</TableCell>
           <TableCell align="center">abilitypatientrightsId</TableCell>
           <TableCell align="center">operative-medicalSupplies-examine</TableCell>
           <TableCell align="center">Manage</TableCell>
           
         </TableRow>
       </TableHead>
       <TableBody>
       {Patientrightstypes === undefined 
          ? null
          : Patientrightstypes.map((item :any)=> (
           <TableRow>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.permission}</TableCell>
             
             <TableCell align="center">{item.permissionArea}</TableCell>
             <TableCell align="center">{item.responsible}</TableCell>
             <TableCell align="center">{item.edges.patientrightstypeAbilitypatientrights.id}</TableCell>
             <TableCell align="center">{item.edges.patientrightstypeAbilitypatientrights.operative }-{item.edges.patientrightstypeAbilitypatientrights.medicalSupplies}-{item.edges.patientrightstypeAbilitypatientrights.examine}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                  deletePatientrightstypes(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>

               
               

           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
 
