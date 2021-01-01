package schema
// รวม
import (
    
    "github.com/facebookincubator/ent"
	_ "github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Medicalrecordstaff holds the schema definition for the Medicalrecordstaff entity.
type Medicalrecordstaff struct {
	ent.Schema
}

// Fields of the Medicalrecordstaff.
func (Medicalrecordstaff) Fields() []ent.Field {
	return nil
}

// Edges of the Medicalrecordstaff.
func (Medicalrecordstaff) Edges() []ent.Edge {
	return []ent.Edge{
		
		edge.To("MedicalrecordstaffPatientrights", Patientrights.Type).StorageKey(edge.Column("Medicalrecordstaff_id")),

    }
}
