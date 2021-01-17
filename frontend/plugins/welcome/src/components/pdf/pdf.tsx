import * as React from "react";


import PdfDocument from "./PdfDocument";
import TestDocument from "./TestDocument";
import {
  FormControl,
Button,
} from '@material-ui/core';



const PDFLink: React.FC = () => {



  return (
    <div>
               <form noValidate autoComplete="off">
              
      <Button>
     
        <PdfDocument
          title="Cost Disclosure Document"
          document={<TestDocument  />}
        />
  </Button>
 

 </form>
  
  </div>
  );
};

export default PDFLink;
