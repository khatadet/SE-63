package schema
//เป็น OUTPUT หลักของ B6103866
import (
    
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Patientrights holds the schema definition for the Patientrights entity.
type Patientrights struct {
	ent.Schema
}

// Fields of the Patientrights.
func (Patientrights) Fields() []ent.Field {
	return []ent.Field{
		field.Time("PermissionDate"),//เวลาที่สร้างสิทธิ์
		
			
    }
}
// Edges of the Patientrights.
func (Patientrights) Edges() []ent.Edge {
	return []ent.Edge{

		
		edge.From("PatientrightsPatientrightstype", Patientrightstype.Type).
			Ref("PatientrightstypePatientrights").
			Unique(),
			//รูปแบบสิทธิ์

		/*
		edge.To("Patientrights___", ___.Type).StorageKey(edge.Column("Patientrights_id")),

		*/

		edge.From("PatientrightsInsurance", Insurance.Type).
			Ref("InsurancePatientrights").
			Unique(),
			//ผู้จ่ายเงิน

		edge.From("PatientrightsPatientrecord", Patientrecord.Type).
			Ref("PatientrecordPatientrights").
			Unique(),
			//ผู้ป่วย ผู้รับสิทธิ์

		edge.From("PatientrightsMedicalrecordstaff", Medicalrecordstaff.Type).
			Ref("MedicalrecordstaffPatientrights").
			Unique(),
			//พนักงานผู้กรอกข้อมูล

	}
}
