// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/PON/app/ent/insurance"
	"github.com/PON/app/ent/medicalrecordstaff"
	"github.com/PON/app/ent/patientrecord"
	"github.com/PON/app/ent/patientrights"
	"github.com/PON/app/ent/patientrightstype"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Patientrights is the model entity for the Patientrights schema.
type Patientrights struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PermissionDate holds the value of the "PermissionDate" field.
	PermissionDate string `json:"PermissionDate,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientrightsQuery when eager-loading is set.
	Edges                     PatientrightsEdges `json:"edges"`
	InsurancePatientrights_id *int
	Medicalrecordstaff_id     *int
	Patientrecord_id          *int
	Patientrightstype_id      *int
}

// PatientrightsEdges holds the relations/edges for other nodes in the graph.
type PatientrightsEdges struct {
	// PatientrightsPatientrightstype holds the value of the PatientrightsPatientrightstype edge.
	PatientrightsPatientrightstype *Patientrightstype
	// PatientrightsInsurance holds the value of the PatientrightsInsurance edge.
	PatientrightsInsurance *Insurance
	// PatientrightsPatientrecord holds the value of the PatientrightsPatientrecord edge.
	PatientrightsPatientrecord *Patientrecord
	// PatientrightsMedicalrecordstaff holds the value of the PatientrightsMedicalrecordstaff edge.
	PatientrightsMedicalrecordstaff *Medicalrecordstaff
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PatientrightsPatientrightstypeOrErr returns the PatientrightsPatientrightstype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientrightsEdges) PatientrightsPatientrightstypeOrErr() (*Patientrightstype, error) {
	if e.loadedTypes[0] {
		if e.PatientrightsPatientrightstype == nil {
			// The edge PatientrightsPatientrightstype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patientrightstype.Label}
		}
		return e.PatientrightsPatientrightstype, nil
	}
	return nil, &NotLoadedError{edge: "PatientrightsPatientrightstype"}
}

// PatientrightsInsuranceOrErr returns the PatientrightsInsurance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientrightsEdges) PatientrightsInsuranceOrErr() (*Insurance, error) {
	if e.loadedTypes[1] {
		if e.PatientrightsInsurance == nil {
			// The edge PatientrightsInsurance was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: insurance.Label}
		}
		return e.PatientrightsInsurance, nil
	}
	return nil, &NotLoadedError{edge: "PatientrightsInsurance"}
}

// PatientrightsPatientrecordOrErr returns the PatientrightsPatientrecord value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientrightsEdges) PatientrightsPatientrecordOrErr() (*Patientrecord, error) {
	if e.loadedTypes[2] {
		if e.PatientrightsPatientrecord == nil {
			// The edge PatientrightsPatientrecord was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patientrecord.Label}
		}
		return e.PatientrightsPatientrecord, nil
	}
	return nil, &NotLoadedError{edge: "PatientrightsPatientrecord"}
}

// PatientrightsMedicalrecordstaffOrErr returns the PatientrightsMedicalrecordstaff value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientrightsEdges) PatientrightsMedicalrecordstaffOrErr() (*Medicalrecordstaff, error) {
	if e.loadedTypes[3] {
		if e.PatientrightsMedicalrecordstaff == nil {
			// The edge PatientrightsMedicalrecordstaff was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: medicalrecordstaff.Label}
		}
		return e.PatientrightsMedicalrecordstaff, nil
	}
	return nil, &NotLoadedError{edge: "PatientrightsMedicalrecordstaff"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Patientrights) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // PermissionDate
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Patientrights) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // InsurancePatientrights_id
		&sql.NullInt64{}, // Medicalrecordstaff_id
		&sql.NullInt64{}, // Patientrecord_id
		&sql.NullInt64{}, // Patientrightstype_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Patientrights fields.
func (pa *Patientrights) assignValues(values ...interface{}) error {
	if m, n := len(values), len(patientrights.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PermissionDate", values[0])
	} else if value.Valid {
		pa.PermissionDate = value.String
	}
	values = values[1:]
	if len(values) == len(patientrights.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field InsurancePatientrights_id", value)
		} else if value.Valid {
			pa.InsurancePatientrights_id = new(int)
			*pa.InsurancePatientrights_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Medicalrecordstaff_id", value)
		} else if value.Valid {
			pa.Medicalrecordstaff_id = new(int)
			*pa.Medicalrecordstaff_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Patientrecord_id", value)
		} else if value.Valid {
			pa.Patientrecord_id = new(int)
			*pa.Patientrecord_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Patientrightstype_id", value)
		} else if value.Valid {
			pa.Patientrightstype_id = new(int)
			*pa.Patientrightstype_id = int(value.Int64)
		}
	}
	return nil
}

// QueryPatientrightsPatientrightstype queries the PatientrightsPatientrightstype edge of the Patientrights.
func (pa *Patientrights) QueryPatientrightsPatientrightstype() *PatientrightstypeQuery {
	return (&PatientrightsClient{config: pa.config}).QueryPatientrightsPatientrightstype(pa)
}

// QueryPatientrightsInsurance queries the PatientrightsInsurance edge of the Patientrights.
func (pa *Patientrights) QueryPatientrightsInsurance() *InsuranceQuery {
	return (&PatientrightsClient{config: pa.config}).QueryPatientrightsInsurance(pa)
}

// QueryPatientrightsPatientrecord queries the PatientrightsPatientrecord edge of the Patientrights.
func (pa *Patientrights) QueryPatientrightsPatientrecord() *PatientrecordQuery {
	return (&PatientrightsClient{config: pa.config}).QueryPatientrightsPatientrecord(pa)
}

// QueryPatientrightsMedicalrecordstaff queries the PatientrightsMedicalrecordstaff edge of the Patientrights.
func (pa *Patientrights) QueryPatientrightsMedicalrecordstaff() *MedicalrecordstaffQuery {
	return (&PatientrightsClient{config: pa.config}).QueryPatientrightsMedicalrecordstaff(pa)
}

// Update returns a builder for updating this Patientrights.
// Note that, you need to call Patientrights.Unwrap() before calling this method, if this Patientrights
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Patientrights) Update() *PatientrightsUpdateOne {
	return (&PatientrightsClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Patientrights) Unwrap() *Patientrights {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Patientrights is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Patientrights) String() string {
	var builder strings.Builder
	builder.WriteString("Patientrights(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", PermissionDate=")
	builder.WriteString(pa.PermissionDate)
	builder.WriteByte(')')
	return builder.String()
}

// PatientrightsSlice is a parsable slice of Patientrights.
type PatientrightsSlice []*Patientrights

func (pa PatientrightsSlice) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
