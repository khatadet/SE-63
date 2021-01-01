// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/PON/app/ent/abilitypatientrights"
	"github.com/PON/app/ent/patientrights"
	"github.com/PON/app/ent/patientrightstype"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PatientrightstypeCreate is the builder for creating a Patientrightstype entity.
type PatientrightstypeCreate struct {
	config
	mutation *PatientrightstypeMutation
	hooks    []Hook
}

// SetPermission sets the Permission field.
func (pc *PatientrightstypeCreate) SetPermission(s string) *PatientrightstypeCreate {
	pc.mutation.SetPermission(s)
	return pc
}

// SetPermissionArea sets the PermissionArea field.
func (pc *PatientrightstypeCreate) SetPermissionArea(s string) *PatientrightstypeCreate {
	pc.mutation.SetPermissionArea(s)
	return pc
}

// SetResponsible sets the Responsible field.
func (pc *PatientrightstypeCreate) SetResponsible(s string) *PatientrightstypeCreate {
	pc.mutation.SetResponsible(s)
	return pc
}

// AddPatientrightstypePatientrightIDs adds the PatientrightstypePatientrights edge to Patientrights by ids.
func (pc *PatientrightstypeCreate) AddPatientrightstypePatientrightIDs(ids ...int) *PatientrightstypeCreate {
	pc.mutation.AddPatientrightstypePatientrightIDs(ids...)
	return pc
}

// AddPatientrightstypePatientrights adds the PatientrightstypePatientrights edges to Patientrights.
func (pc *PatientrightstypeCreate) AddPatientrightstypePatientrights(p ...*Patientrights) *PatientrightstypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPatientrightstypePatientrightIDs(ids...)
}

// SetPatientrightstypeAbilitypatientrightsID sets the PatientrightstypeAbilitypatientrights edge to Abilitypatientrights by id.
func (pc *PatientrightstypeCreate) SetPatientrightstypeAbilitypatientrightsID(id int) *PatientrightstypeCreate {
	pc.mutation.SetPatientrightstypeAbilitypatientrightsID(id)
	return pc
}

// SetNillablePatientrightstypeAbilitypatientrightsID sets the PatientrightstypeAbilitypatientrights edge to Abilitypatientrights by id if the given value is not nil.
func (pc *PatientrightstypeCreate) SetNillablePatientrightstypeAbilitypatientrightsID(id *int) *PatientrightstypeCreate {
	if id != nil {
		pc = pc.SetPatientrightstypeAbilitypatientrightsID(*id)
	}
	return pc
}

// SetPatientrightstypeAbilitypatientrights sets the PatientrightstypeAbilitypatientrights edge to Abilitypatientrights.
func (pc *PatientrightstypeCreate) SetPatientrightstypeAbilitypatientrights(a *Abilitypatientrights) *PatientrightstypeCreate {
	return pc.SetPatientrightstypeAbilitypatientrightsID(a.ID)
}

// Mutation returns the PatientrightstypeMutation object of the builder.
func (pc *PatientrightstypeCreate) Mutation() *PatientrightstypeMutation {
	return pc.mutation
}

// Save creates the Patientrightstype in the database.
func (pc *PatientrightstypeCreate) Save(ctx context.Context) (*Patientrightstype, error) {
	if _, ok := pc.mutation.Permission(); !ok {
		return nil, &ValidationError{Name: "Permission", err: errors.New("ent: missing required field \"Permission\"")}
	}
	if _, ok := pc.mutation.PermissionArea(); !ok {
		return nil, &ValidationError{Name: "PermissionArea", err: errors.New("ent: missing required field \"PermissionArea\"")}
	}
	if _, ok := pc.mutation.Responsible(); !ok {
		return nil, &ValidationError{Name: "Responsible", err: errors.New("ent: missing required field \"Responsible\"")}
	}
	var (
		err  error
		node *Patientrightstype
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientrightstypeMutation)
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
func (pc *PatientrightstypeCreate) SaveX(ctx context.Context) *Patientrightstype {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PatientrightstypeCreate) sqlSave(ctx context.Context) (*Patientrightstype, error) {
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

func (pc *PatientrightstypeCreate) createSpec() (*Patientrightstype, *sqlgraph.CreateSpec) {
	var (
		pa    = &Patientrightstype{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: patientrightstype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientrightstype.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Permission(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrightstype.FieldPermission,
		})
		pa.Permission = value
	}
	if value, ok := pc.mutation.PermissionArea(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrightstype.FieldPermissionArea,
		})
		pa.PermissionArea = value
	}
	if value, ok := pc.mutation.Responsible(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrightstype.FieldResponsible,
		})
		pa.Responsible = value
	}
	if nodes := pc.mutation.PatientrightstypePatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrightstype.PatientrightstypePatientrightsTable,
			Columns: []string{patientrightstype.PatientrightstypePatientrightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrights.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PatientrightstypeAbilitypatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patientrightstype.PatientrightstypeAbilitypatientrightsTable,
			Columns: []string{patientrightstype.PatientrightstypeAbilitypatientrightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: abilitypatientrights.FieldID,
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
