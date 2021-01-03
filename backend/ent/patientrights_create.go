// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/PON/app/ent/insurance"
	"github.com/PON/app/ent/medicalrecordstaff"
	"github.com/PON/app/ent/patientrecord"
	"github.com/PON/app/ent/patientrights"
	"github.com/PON/app/ent/patientrightstype"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PatientrightsCreate is the builder for creating a Patientrights entity.
type PatientrightsCreate struct {
	config
	mutation *PatientrightsMutation
	hooks    []Hook
}

// SetPermissionDate sets the PermissionDate field.
func (pc *PatientrightsCreate) SetPermissionDate(t time.Time) *PatientrightsCreate {
	pc.mutation.SetPermissionDate(t)
	return pc
}

// SetPatientrightsPatientrightstypeID sets the PatientrightsPatientrightstype edge to Patientrightstype by id.
func (pc *PatientrightsCreate) SetPatientrightsPatientrightstypeID(id int) *PatientrightsCreate {
	pc.mutation.SetPatientrightsPatientrightstypeID(id)
	return pc
}

// SetNillablePatientrightsPatientrightstypeID sets the PatientrightsPatientrightstype edge to Patientrightstype by id if the given value is not nil.
func (pc *PatientrightsCreate) SetNillablePatientrightsPatientrightstypeID(id *int) *PatientrightsCreate {
	if id != nil {
		pc = pc.SetPatientrightsPatientrightstypeID(*id)
	}
	return pc
}

// SetPatientrightsPatientrightstype sets the PatientrightsPatientrightstype edge to Patientrightstype.
func (pc *PatientrightsCreate) SetPatientrightsPatientrightstype(p *Patientrightstype) *PatientrightsCreate {
	return pc.SetPatientrightsPatientrightstypeID(p.ID)
}

// SetPatientrightsInsuranceID sets the PatientrightsInsurance edge to Insurance by id.
func (pc *PatientrightsCreate) SetPatientrightsInsuranceID(id int) *PatientrightsCreate {
	pc.mutation.SetPatientrightsInsuranceID(id)
	return pc
}

// SetNillablePatientrightsInsuranceID sets the PatientrightsInsurance edge to Insurance by id if the given value is not nil.
func (pc *PatientrightsCreate) SetNillablePatientrightsInsuranceID(id *int) *PatientrightsCreate {
	if id != nil {
		pc = pc.SetPatientrightsInsuranceID(*id)
	}
	return pc
}

// SetPatientrightsInsurance sets the PatientrightsInsurance edge to Insurance.
func (pc *PatientrightsCreate) SetPatientrightsInsurance(i *Insurance) *PatientrightsCreate {
	return pc.SetPatientrightsInsuranceID(i.ID)
}

// SetPatientrightsPatientrecordID sets the PatientrightsPatientrecord edge to Patientrecord by id.
func (pc *PatientrightsCreate) SetPatientrightsPatientrecordID(id int) *PatientrightsCreate {
	pc.mutation.SetPatientrightsPatientrecordID(id)
	return pc
}

// SetNillablePatientrightsPatientrecordID sets the PatientrightsPatientrecord edge to Patientrecord by id if the given value is not nil.
func (pc *PatientrightsCreate) SetNillablePatientrightsPatientrecordID(id *int) *PatientrightsCreate {
	if id != nil {
		pc = pc.SetPatientrightsPatientrecordID(*id)
	}
	return pc
}

// SetPatientrightsPatientrecord sets the PatientrightsPatientrecord edge to Patientrecord.
func (pc *PatientrightsCreate) SetPatientrightsPatientrecord(p *Patientrecord) *PatientrightsCreate {
	return pc.SetPatientrightsPatientrecordID(p.ID)
}

// SetPatientrightsMedicalrecordstaffID sets the PatientrightsMedicalrecordstaff edge to Medicalrecordstaff by id.
func (pc *PatientrightsCreate) SetPatientrightsMedicalrecordstaffID(id int) *PatientrightsCreate {
	pc.mutation.SetPatientrightsMedicalrecordstaffID(id)
	return pc
}

// SetNillablePatientrightsMedicalrecordstaffID sets the PatientrightsMedicalrecordstaff edge to Medicalrecordstaff by id if the given value is not nil.
func (pc *PatientrightsCreate) SetNillablePatientrightsMedicalrecordstaffID(id *int) *PatientrightsCreate {
	if id != nil {
		pc = pc.SetPatientrightsMedicalrecordstaffID(*id)
	}
	return pc
}

// SetPatientrightsMedicalrecordstaff sets the PatientrightsMedicalrecordstaff edge to Medicalrecordstaff.
func (pc *PatientrightsCreate) SetPatientrightsMedicalrecordstaff(m *Medicalrecordstaff) *PatientrightsCreate {
	return pc.SetPatientrightsMedicalrecordstaffID(m.ID)
}

// Mutation returns the PatientrightsMutation object of the builder.
func (pc *PatientrightsCreate) Mutation() *PatientrightsMutation {
	return pc.mutation
}

// Save creates the Patientrights in the database.
func (pc *PatientrightsCreate) Save(ctx context.Context) (*Patientrights, error) {
	if _, ok := pc.mutation.PermissionDate(); !ok {
		return nil, &ValidationError{Name: "PermissionDate", err: errors.New("ent: missing required field \"PermissionDate\"")}
	}
	var (
		err  error
		node *Patientrights
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientrightsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PatientrightsCreate) SaveX(ctx context.Context) *Patientrights {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PatientrightsCreate) sqlSave(ctx context.Context) (*Patientrights, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PatientrightsCreate) createSpec() (*Patientrights, *sqlgraph.CreateSpec) {
	var (
		pa    = &Patientrights{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: patientrights.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientrights.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PermissionDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: patientrights.FieldPermissionDate,
		})
		pa.PermissionDate = value
	}
	if nodes := pc.mutation.PatientrightsPatientrightstypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrights.PatientrightsPatientrightstypeTable,
			Columns: []string{patientrights.PatientrightsPatientrightstypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrightstype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PatientrightsInsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrights.PatientrightsInsuranceTable,
			Columns: []string{patientrights.PatientrightsInsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: insurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PatientrightsPatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrights.PatientrightsPatientrecordTable,
			Columns: []string{patientrights.PatientrightsPatientrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PatientrightsMedicalrecordstaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrights.PatientrightsMedicalrecordstaffTable,
			Columns: []string{patientrights.PatientrightsMedicalrecordstaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalrecordstaff.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pa, _spec
}
