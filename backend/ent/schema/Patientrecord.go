package schema
// ของแพร
import (
    
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Patientrecord holds the schema definition for the Patientrecord entity.
type Patientrecord struct {
	ent.Schema
}

// Fields of the Patientrecord.
func (Patientrecord) Fields() []ent.Field {
	return []ent.Field{

        field.String("Name"),

    }
}

// Edges of the Patientrecord.
func (Patientrecord) Edges() []ent.Edge {
	return []ent.Edge{
		
		edge.To("PatientrecordPatientrights", Patientrights.Type).StorageKey(edge.Column("Patientrecord_id")),

    }
}
