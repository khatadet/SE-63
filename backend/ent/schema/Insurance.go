package schema
//เป็นส่วนประกอบของ OUTPUT หลัก
import (
    
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Insurance holds the schema definition for the Insurance entity.
type Insurance struct {
	ent.Schema
}

// Fields of the INSURANCE.
func (Insurance) Fields() []ent.Field {
	return []ent.Field{


		field.String("Insurancecompany").Unique(),//ชื่อบริษัท ที่รับผิดชอบ
		

    }
}

// Edges of the INSURANCE.
func (Insurance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("InsurancePatientrights", Patientrights.Type).StorageKey(edge.Column("Insurance_id")),//เป็นส่วนประกอบของ OUTPUT หลัก เพื่อบอกว่าใครจะแป็นคนออกค้าใช้จ่าย
	
    }
}
