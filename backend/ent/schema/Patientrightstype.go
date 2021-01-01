package schema

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
		field.String("Permission"),
		field.String("PermissionArea"),
		field.String("Responsible"),
	}
}

// Edges of the Patientrightstype.
func (Patientrightstype) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("PatientrightstypePatientrights", Patientrights.Type).StorageKey(edge.Column("Patientrightstype_id")),

		edge.From("PatientrightstypeAbilitypatientrights", Abilitypatientrights.Type).
			Ref("AbilitypatientrightsPatientrightstype").
			Unique(),

	}
}
