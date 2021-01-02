// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/PON/app/ent/insurance"
	"github.com/PON/app/ent/patientrights"
	"github.com/PON/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// InsuranceUpdate is the builder for updating Insurance entities.
type InsuranceUpdate struct {
	config
	hooks      []Hook
	mutation   *InsuranceMutation
	predicates []predicate.Insurance
}

// Where adds a new predicate for the builder.
func (iu *InsuranceUpdate) Where(ps ...predicate.Insurance) *InsuranceUpdate {
	iu.predicates = append(iu.predicates, ps...)
	return iu
}

// SetInsurancecompany sets the Insurancecompany field.
func (iu *InsuranceUpdate) SetInsurancecompany(s string) *InsuranceUpdate {
	iu.mutation.SetInsurancecompany(s)
	return iu
}

// AddInsurancePatientrightIDs adds the InsurancePatientrights edge to Patientrights by ids.
func (iu *InsuranceUpdate) AddInsurancePatientrightIDs(ids ...int) *InsuranceUpdate {
	iu.mutation.AddInsurancePatientrightIDs(ids...)
	return iu
}

// AddInsurancePatientrights adds the InsurancePatientrights edges to Patientrights.
func (iu *InsuranceUpdate) AddInsurancePatientrights(p ...*Patientrights) *InsuranceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iu.AddInsurancePatientrightIDs(ids...)
}

// Mutation returns the InsuranceMutation object of the builder.
func (iu *InsuranceUpdate) Mutation() *InsuranceMutation {
	return iu.mutation
}

// RemoveInsurancePatientrightIDs removes the InsurancePatientrights edge to Patientrights by ids.
func (iu *InsuranceUpdate) RemoveInsurancePatientrightIDs(ids ...int) *InsuranceUpdate {
	iu.mutation.RemoveInsurancePatientrightIDs(ids...)
	return iu
}

// RemoveInsurancePatientrights removes InsurancePatientrights edges to Patientrights.
func (iu *InsuranceUpdate) RemoveInsurancePatientrights(p ...*Patientrights) *InsuranceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iu.RemoveInsurancePatientrightIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (iu *InsuranceUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InsuranceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InsuranceUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InsuranceUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InsuranceUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *InsuranceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   insurance.Table,
			Columns: insurance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: insurance.FieldID,
			},
		},
	}
	if ps := iu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Insurancecompany(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: insurance.FieldInsurancecompany,
		})
	}
	if nodes := iu.mutation.RemovedInsurancePatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   insurance.InsurancePatientrightsTable,
			Columns: []string{insurance.InsurancePatientrightsColumn},
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
	if nodes := iu.mutation.InsurancePatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   insurance.InsurancePatientrightsTable,
			Columns: []string{insurance.InsurancePatientrightsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{insurance.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// InsuranceUpdateOne is the builder for updating a single Insurance entity.
type InsuranceUpdateOne struct {
	config
	hooks    []Hook
	mutation *InsuranceMutation
}

// SetInsurancecompany sets the Insurancecompany field.
func (iuo *InsuranceUpdateOne) SetInsurancecompany(s string) *InsuranceUpdateOne {
	iuo.mutation.SetInsurancecompany(s)
	return iuo
}

// AddInsurancePatientrightIDs adds the InsurancePatientrights edge to Patientrights by ids.
func (iuo *InsuranceUpdateOne) AddInsurancePatientrightIDs(ids ...int) *InsuranceUpdateOne {
	iuo.mutation.AddInsurancePatientrightIDs(ids...)
	return iuo
}

// AddInsurancePatientrights adds the InsurancePatientrights edges to Patientrights.
func (iuo *InsuranceUpdateOne) AddInsurancePatientrights(p ...*Patientrights) *InsuranceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iuo.AddInsurancePatientrightIDs(ids...)
}

// Mutation returns the InsuranceMutation object of the builder.
func (iuo *InsuranceUpdateOne) Mutation() *InsuranceMutation {
	return iuo.mutation
}

// RemoveInsurancePatientrightIDs removes the InsurancePatientrights edge to Patientrights by ids.
func (iuo *InsuranceUpdateOne) RemoveInsurancePatientrightIDs(ids ...int) *InsuranceUpdateOne {
	iuo.mutation.RemoveInsurancePatientrightIDs(ids...)
	return iuo
}

// RemoveInsurancePatientrights removes InsurancePatientrights edges to Patientrights.
func (iuo *InsuranceUpdateOne) RemoveInsurancePatientrights(p ...*Patientrights) *InsuranceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iuo.RemoveInsurancePatientrightIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (iuo *InsuranceUpdateOne) Save(ctx context.Context) (*Insurance, error) {

	var (
		err  error
		node *Insurance
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InsuranceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InsuranceUpdateOne) SaveX(ctx context.Context) *Insurance {
	i, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return i
}

// Exec executes the query on the entity.
func (iuo *InsuranceUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InsuranceUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *InsuranceUpdateOne) sqlSave(ctx context.Context) (i *Insurance, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   insurance.Table,
			Columns: insurance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: insurance.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Insurance.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := iuo.mutation.Insurancecompany(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: insurance.FieldInsurancecompany,
		})
	}
	if nodes := iuo.mutation.RemovedInsurancePatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   insurance.InsurancePatientrightsTable,
			Columns: []string{insurance.InsurancePatientrightsColumn},
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
	if nodes := iuo.mutation.InsurancePatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   insurance.InsurancePatientrightsTable,
			Columns: []string{insurance.InsurancePatientrightsColumn},
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
	i = &Insurance{config: iuo.config}
	_spec.Assign = i.assignValues
	_spec.ScanValues = i.scanValues()
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{insurance.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return i, nil
}