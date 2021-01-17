import {
  Document,
  Font,
  Page,
  StyleSheet,
  Text,
  View,
  Image
} from "@react-pdf/renderer";
import * as React from "react";
import CardMedia from '@material-ui/core/CardMedia';
import { Image1Base64Function } from './Image1';
import ComponanceTable from '../Table_Patientrightstype';
export const styles = StyleSheet.create({
  font: { fontFamily: "Oswald" }
});
const fontSrc =
  "https://fonts.gstatic.com/s/oswald/v13/Y_TKV6o8WovbUd3m_X9aAA.ttf";
Font.register({ family: "Oswald", src: fontSrc });

const TestDocument: React.FC<{}> = () => {


  return (
    <Document>
      <Page size="A4" style={styles.font}>
        <View>
       
           
          
          
            <Image  source ={ Image1Base64Function} />
          <Text>Email:sona1</Text>
          <Text>Email:sona2</Text>
          <Text>Email:sona3</Text>
          <Text>Email:sona4</Text>
          <Text>Email:sona5</Text>
          </View>
      </Page>
    
      <Page size="A4" style={styles.font}>
        <View>
       
          <Text>Email:sona6</Text>
          <Text>Email:sona7</Text>
          <Text>Email:sona8</Text>
          <Text>Email:sona9</Text>
          <Text>Email:sona10</Text>
          <Text>Email:sona11</Text>
          <Text>Email:sona12</Text>
          <Text>Email:sona13</Text>
          <Text>Email:sona14</Text>
         

        </View>
      </Page>
    </Document>
  );

};

export default TestDocument;
