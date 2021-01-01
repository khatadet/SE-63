package schema

import (
    
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Abilitypatientrights holds the schema definition for the Abilitypatientrights entity.
type Abilitypatientrights struct {
	ent.Schema
}

// Fields of the Abilitypatientrights.
func (Abilitypatientrights) Fields() []ent.Field {
	return []ent.Field{
		field.String("Operative"),
		field.String("MedicalSupplies"),
		field.String("Examine"),
			
    }
}

// Edges of the Abilitypatientrights.
func (Abilitypatientrights) Edges() []ent.Edge {
	return []ent.Edge{
		
		edge.To("AbilitypatientrightsPatientrightstype", Patientrightstype.Type).StorageKey(edge.Column("Abilitypatientrights_id")),
    }
}
