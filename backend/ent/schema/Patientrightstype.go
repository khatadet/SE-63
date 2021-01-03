package schema
//เป็นส่วนประกอบของ OUTPUT หลัก
import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Patientrightstype holds the schema definition for the Patientrightstype entity.
type Patientrightstype struct {
	ent.Schema
}

// Fields of the Patientrightstype.
func (Patientrightstype) Fields() []ent.Field {
	return []ent.Field{
		field.String("Permission"),//
		field.String("PermissionArea"),//พื้นที่ หรือเขต หรือโรงพยาบาล ที่สามารถใช้สิทได้
		field.String("Responsible"),//
	}
}

// Edges of the Patientrightstype.
func (Patientrightstype) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("PatientrightstypePatientrights", Patientrights.Type).StorageKey(edge.Column("Patientrightstype_id")),//เป็นส่วนประกอบของ OUTPUT หลัก

		edge.From("PatientrightstypeAbilitypatientrights", Abilitypatientrights.Type).
			Ref("AbilitypatientrightsPatientrightstype").
			Unique(),
			//บอกว่า รูปแบบสิทธิ์นี้มีความสามารถอะไรบ้าง

	}
}
