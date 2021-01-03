package main
 
import (
   "context"
   "fmt"
   "log"
   "time"
  
 
   "github.com/PON/app/controllers"
   _ "github.com/PON/app/docs"
   "github.com/PON/app/ent"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"

   "github.com/PON/app/ent/abilitypatientrights"


   "github.com/PON/app/ent/insurance"
	"github.com/PON/app/ent/patientrecord"
	"github.com/PON/app/ent/patientrightstype"
	"github.com/PON/app/ent/medicalrecordstaff"


)

//-------------------------------------------------------------------



// Patientrightstypes defines the struct for the Patientrightstypes
type Patientrightstypes struct {
	Patientrightstype []Patientrightstype
}

// Patientrightstype defines the struct for the Patientrightstype
type Patientrightstype struct {
	Permission          string
    PermissionArea      string
    Responsible         string
    Abilitypatientrights    int
}

// Abilitypatientrightss defines the struct for the Abilitypatientrightss
type Abilitypatientrightss struct {
	Abilitypatientrights []Abilitypatientrights
}

// Abilitypatientrights defines the struct for the Abilitypatientrights
type Abilitypatientrights struct {
	Operative       string
    MedicalSupplies string
    Examine         string
}

// Patientrightss defines the struct for the Patientrightss
type Patientrightss struct {
	Patientrights []Patientrights
}

// Patientrights defines the struct for the Patientrights
type Patientrights struct {
	
	Patientrightstype    int
	Patientrecord        int
	Insurance            int
	Medicalrecordstaff   int
}




//-------------------------------------------------------------------




// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
 
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
 
// @host localhost:8080
// @BasePath /api/v1
 
// @securityDefinitions.basic BasicAuth
 
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
 
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
   router := gin.Default()
   router.Use(cors.Default())
 
   client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
   if err != nil {
       log.Fatalf("fail to open sqlite3: %v", err)
   }
   defer client.Close()
 
   if err := client.Schema.Create(context.Background()); err != nil {
       log.Fatalf("failed creating schema resources: %v", err)
   }
 
   v1 := router.Group("/api/v1")
   controllers.NewPatientrightsController(v1, client)
   controllers.NewPatientrightstypeController(v1, client)
   controllers.NewAbilitypatientrightsController(v1, client)
   controllers.NewMedicalrecordstaffController(v1, client)
   controllers.NewInsuranceController(v1, client)
   controllers.NewPatientrecordController(v1, client)

   //-------------------------------------------------------------------

   // Set medicalrecordstaff Data
   	Patientrecord := []string{"Khatadet khianchainat","nara haru","morani rode","faratell yova","pulla visan","omaha yad",}
	for _, r := range Patientrecord {
		client.Patientrecord.
			Create().
			SetName(r).
			Save(context.Background())
	}

	// Set medicalrecordstaff Data
	
	for i := 1; i < 5; i++ {
		client.Medicalrecordstaff.
			Create().
			
			Save(context.Background())
	}

    // Set insurance Data
	Insurance := []string{"เมืองไทยประกันภัย", "ไทยสมุทรประกันชีวิต", "อื่นๆ", "กรมบัญชีกลาง", "AIA"}
	for _, r := range Insurance {
		client.Insurance.
			Create().
			SetInsurancecompany(r).
			Save(context.Background())
    }

    // Set Abilitypatientrights Data
	Abilitypatientrights := Abilitypatientrightss{
		Abilitypatientrights: []Abilitypatientrights{
			Abilitypatientrights{"100", "100","100"},
			Abilitypatientrights{"50", "100","100"},
			Abilitypatientrights{"50", "100","50"},
		},
	}

	for _, a := range Abilitypatientrights.Abilitypatientrights {
		client.Abilitypatientrights.
			Create().
			SetOperative(a.Operative).
		    SetMedicalSupplies(a.MedicalSupplies).
		    SetExamine(a.Examine).
			Save(context.Background())
    }
    

    // Set Patientrightstypes Data
	patientrightstypes := Patientrightstypes{
		Patientrightstype : []Patientrightstype{
			Patientrightstype{"จ่ายตรง", "ทั่วประเทศ","ราชการ", 1},
			Patientrightstype{"ผู้สูงอายุ", "ทั่วประเทศ","นาย แหวง" , 2},
			Patientrightstype{"สุขภาพถ้วนหน้า", "โคราช","ราชการ" ,2},
			Patientrightstype{"อุบัติเหตุสบายใจ", "ทั่วประเทศ","นาย มา", 2},
		},
	}

	for _, p := range patientrightstypes.Patientrightstype {

		a, err := client.Abilitypatientrights.
		Query().
		Where(abilitypatientrights.IDEQ(int(p.Abilitypatientrights))).
		Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Patientrightstype.
		Create().
		SetPermission(p.Permission).
		SetPermissionArea(p.PermissionArea).
		SetResponsible(p.Responsible).
		SetPatientrightstypeAbilitypatientrights(a).
		Save(context.Background())
	}
	//***********************************************************************************
	// Set Patientrightstypes Data
	patientrightss := Patientrightss{
		Patientrights : []Patientrights{
			Patientrights{1,1,1,1},
			Patientrights{1,1,1,1},
			Patientrights{1,1,1,1},
			Patientrights{1,1,1,1},
		},
	}

	for _, p := range patientrightss.Patientrights {

		
	Patientrightstype, err := client.Patientrightstype.
		Query().
		Where(patientrightstype.IDEQ(int(p.Patientrightstype))).
		Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		Insurance, err := client.Insurance.
		Query().
		Where(insurance.IDEQ(int(p.Insurance))).
		Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

	Patientrecord, err := client.Patientrecord.
		Query().
		Where(patientrecord.IDEQ(int(p.Patientrecord))).
		Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

	Medicalrecordstaff, err := client.Medicalrecordstaff.
		Query().
		Where(medicalrecordstaff.IDEQ(int(p.Medicalrecordstaff))).
		Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

	
	t := time.Now().Local()
	client.Patientrights.
		Create().
		SetPermissionDate(t).
		SetPatientrightsPatientrightstype(Patientrightstype).
		SetPatientrightsPatientrecord(Patientrecord).
		SetPatientrightsMedicalrecordstaff(Medicalrecordstaff).
		SetPatientrightsInsurance(Insurance).
		Save(context.Background())

	
	}
    //***********************************************************************
   //-------------------------------------------------------------------

   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Run()
}
 
