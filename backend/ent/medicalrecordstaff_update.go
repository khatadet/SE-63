// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/PON/app/ent/medicalrecordstaff"
	"github.com/PON/app/ent/patientrights"
	"github.com/PON/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// MedicalrecordstaffUpdate is the builder for updating Medicalrecordstaff entities.
type MedicalrecordstaffUpdate struct {
	config
	hooks      []Hook
	mutation   *MedicalrecordstaffMutation
	predicates []predicate.Medicalrecordstaff
}

// Where adds a new predicate for the builder.
func (mu *MedicalrecordstaffUpdate) Where(ps ...predicate.Medicalrecordstaff) *MedicalrecordstaffUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// AddMedicalrecordstaffPatientrightIDs adds the MedicalrecordstaffPatientrights edge to Patientrights by ids.
func (mu *MedicalrecordstaffUpdate) AddMedicalrecordstaffPatientrightIDs(ids ...int) *MedicalrecordstaffUpdate {
	mu.mutation.AddMedicalrecordstaffPatientrightIDs(ids...)
	return mu
}

// AddMedicalrecordstaffPatientrights adds the MedicalrecordstaffPatientrights edges to Patientrights.
func (mu *MedicalrecordstaffUpdate) AddMedicalrecordstaffPatientrights(p ...*Patientrights) *MedicalrecordstaffUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddMedicalrecordstaffPatientrightIDs(ids...)
}

// Mutation returns the MedicalrecordstaffMutation object of the builder.
func (mu *MedicalrecordstaffUpdate) Mutation() *MedicalrecordstaffMutation {
	return mu.mutation
}

// RemoveMedicalrecordstaffPatientrightIDs removes the MedicalrecordstaffPatientrights edge to Patientrights by ids.
func (mu *MedicalrecordstaffUpdate) RemoveMedicalrecordstaffPatientrightIDs(ids ...int) *MedicalrecordstaffUpdate {
	mu.mutation.RemoveMedicalrecordstaffPatientrightIDs(ids...)
	return mu
}

// RemoveMedicalrecordstaffPatientrights removes MedicalrecordstaffPatientrights edges to Patientrights.
func (mu *MedicalrecordstaffUpdate) RemoveMedicalrecordstaffPatientrights(p ...*Patientrights) *MedicalrecordstaffUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemoveMedicalrecordstaffPatientrightIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MedicalrecordstaffUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicalrecordstaffMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MedicalrecordstaffUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MedicalrecordstaffUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MedicalrecordstaffUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MedicalrecordstaffUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicalrecordstaff.Table,
			Columns: medicalrecordstaff.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalrecordstaff.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nodes := mu.mutation.RemovedMedicalrecordstaffPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalrecordstaff.MedicalrecordstaffPatientrightsTable,
			Columns: []string{medicalrecordstaff.MedicalrecordstaffPatientrightsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MedicalrecordstaffPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalrecordstaff.MedicalrecordstaffPatientrightsTable,
			Columns: []string{medicalrecordstaff.MedicalrecordstaffPatientrightsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicalrecordstaff.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MedicalrecordstaffUpdateOne is the builder for updating a single Medicalrecordstaff entity.
type MedicalrecordstaffUpdateOne struct {
	config
	hooks    []Hook
	mutation *MedicalrecordstaffMutation
}

// AddMedicalrecordstaffPatientrightIDs adds the MedicalrecordstaffPatientrights edge to Patientrights by ids.
func (muo *MedicalrecordstaffUpdateOne) AddMedicalrecordstaffPatientrightIDs(ids ...int) *MedicalrecordstaffUpdateOne {
	muo.mutation.AddMedicalrecordstaffPatientrightIDs(ids...)
	return muo
}

// AddMedicalrecordstaffPatientrights adds the MedicalrecordstaffPatientrights edges to Patientrights.
func (muo *MedicalrecordstaffUpdateOne) AddMedicalrecordstaffPatientrights(p ...*Patientrights) *MedicalrecordstaffUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddMedicalrecordstaffPatientrightIDs(ids...)
}

// Mutation returns the MedicalrecordstaffMutation object of the builder.
func (muo *MedicalrecordstaffUpdateOne) Mutation() *MedicalrecordstaffMutation {
	return muo.mutation
}

// RemoveMedicalrecordstaffPatientrightIDs removes the MedicalrecordstaffPatientrights edge to Patientrights by ids.
func (muo *MedicalrecordstaffUpdateOne) RemoveMedicalrecordstaffPatientrightIDs(ids ...int) *MedicalrecordstaffUpdateOne {
	muo.mutation.RemoveMedicalrecordstaffPatientrightIDs(ids...)
	return muo
}

// RemoveMedicalrecordstaffPatientrights removes MedicalrecordstaffPatientrights edges to Patientrights.
func (muo *MedicalrecordstaffUpdateOne) RemoveMedicalrecordstaffPatientrights(p ...*Patientrights) *MedicalrecordstaffUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemoveMedicalrecordstaffPatientrightIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MedicalrecordstaffUpdateOne) Save(ctx context.Context) (*Medicalrecordstaff, error) {

	var (
		err  error
		node *Medicalrecordstaff
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicalrecordstaffMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MedicalrecordstaffUpdateOne) SaveX(ctx context.Context) *Medicalrecordstaff {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *MedicalrecordstaffUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MedicalrecordstaffUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MedicalrecordstaffUpdateOne) sqlSave(ctx context.Context) (m *Medicalrecordstaff, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicalrecordstaff.Table,
			Columns: medicalrecordstaff.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalrecordstaff.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Medicalrecordstaff.ID for update")}
	}
	_spec.Node.ID.Value = id
	if nodes := muo.mutation.RemovedMedicalrecordstaffPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalrecordstaff.MedicalrecordstaffPatientrightsTable,
			Columns: []string{medicalrecordstaff.MedicalrecordstaffPatientrightsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MedicalrecordstaffPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalrecordstaff.MedicalrecordstaffPatientrightsTable,
			Columns: []string{medicalrecordstaff.MedicalrecordstaffPatientrightsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	m = &Medicalrecordstaff{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicalrecordstaff.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}
