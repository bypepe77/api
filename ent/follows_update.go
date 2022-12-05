// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bypepe77/api/ent/follows"
	"github.com/bypepe77/api/ent/predicate"
)

// FollowsUpdate is the builder for updating Follows entities.
type FollowsUpdate struct {
	config
	hooks    []Hook
	mutation *FollowsMutation
}

// Where appends a list predicates to the FollowsUpdate builder.
func (fu *FollowsUpdate) Where(ps ...predicate.Follows) *FollowsUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetFollowedby sets the "followedby" field.
func (fu *FollowsUpdate) SetFollowedby(i int) *FollowsUpdate {
	fu.mutation.ResetFollowedby()
	fu.mutation.SetFollowedby(i)
	return fu
}

// AddFollowedby adds i to the "followedby" field.
func (fu *FollowsUpdate) AddFollowedby(i int) *FollowsUpdate {
	fu.mutation.AddFollowedby(i)
	return fu
}

// SetFollower sets the "follower" field.
func (fu *FollowsUpdate) SetFollower(i int) *FollowsUpdate {
	fu.mutation.ResetFollower()
	fu.mutation.SetFollower(i)
	return fu
}

// AddFollower adds i to the "follower" field.
func (fu *FollowsUpdate) AddFollower(i int) *FollowsUpdate {
	fu.mutation.AddFollower(i)
	return fu
}

// Mutation returns the FollowsMutation object of the builder.
func (fu *FollowsUpdate) Mutation() *FollowsMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FollowsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		if err = fu.check(); err != nil {
			return 0, err
		}
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FollowsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fu.check(); err != nil {
				return 0, err
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FollowsUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FollowsUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FollowsUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FollowsUpdate) check() error {
	if v, ok := fu.mutation.Followedby(); ok {
		if err := follows.FollowedbyValidator(v); err != nil {
			return &ValidationError{Name: "followedby", err: fmt.Errorf(`ent: validator failed for field "Follows.followedby": %w`, err)}
		}
	}
	if v, ok := fu.mutation.Follower(); ok {
		if err := follows.FollowerValidator(v); err != nil {
			return &ValidationError{Name: "follower", err: fmt.Errorf(`ent: validator failed for field "Follows.follower": %w`, err)}
		}
	}
	return nil
}

func (fu *FollowsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   follows.Table,
			Columns: follows.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: follows.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Followedby(); ok {
		_spec.SetField(follows.FieldFollowedby, field.TypeInt, value)
	}
	if value, ok := fu.mutation.AddedFollowedby(); ok {
		_spec.AddField(follows.FieldFollowedby, field.TypeInt, value)
	}
	if value, ok := fu.mutation.Follower(); ok {
		_spec.SetField(follows.FieldFollower, field.TypeInt, value)
	}
	if value, ok := fu.mutation.AddedFollower(); ok {
		_spec.AddField(follows.FieldFollower, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{follows.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FollowsUpdateOne is the builder for updating a single Follows entity.
type FollowsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FollowsMutation
}

// SetFollowedby sets the "followedby" field.
func (fuo *FollowsUpdateOne) SetFollowedby(i int) *FollowsUpdateOne {
	fuo.mutation.ResetFollowedby()
	fuo.mutation.SetFollowedby(i)
	return fuo
}

// AddFollowedby adds i to the "followedby" field.
func (fuo *FollowsUpdateOne) AddFollowedby(i int) *FollowsUpdateOne {
	fuo.mutation.AddFollowedby(i)
	return fuo
}

// SetFollower sets the "follower" field.
func (fuo *FollowsUpdateOne) SetFollower(i int) *FollowsUpdateOne {
	fuo.mutation.ResetFollower()
	fuo.mutation.SetFollower(i)
	return fuo
}

// AddFollower adds i to the "follower" field.
func (fuo *FollowsUpdateOne) AddFollower(i int) *FollowsUpdateOne {
	fuo.mutation.AddFollower(i)
	return fuo
}

// Mutation returns the FollowsMutation object of the builder.
func (fuo *FollowsUpdateOne) Mutation() *FollowsMutation {
	return fuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FollowsUpdateOne) Select(field string, fields ...string) *FollowsUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Follows entity.
func (fuo *FollowsUpdateOne) Save(ctx context.Context) (*Follows, error) {
	var (
		err  error
		node *Follows
	)
	if len(fuo.hooks) == 0 {
		if err = fuo.check(); err != nil {
			return nil, err
		}
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FollowsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fuo.check(); err != nil {
				return nil, err
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Follows)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FollowsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FollowsUpdateOne) SaveX(ctx context.Context) *Follows {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FollowsUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FollowsUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FollowsUpdateOne) check() error {
	if v, ok := fuo.mutation.Followedby(); ok {
		if err := follows.FollowedbyValidator(v); err != nil {
			return &ValidationError{Name: "followedby", err: fmt.Errorf(`ent: validator failed for field "Follows.followedby": %w`, err)}
		}
	}
	if v, ok := fuo.mutation.Follower(); ok {
		if err := follows.FollowerValidator(v); err != nil {
			return &ValidationError{Name: "follower", err: fmt.Errorf(`ent: validator failed for field "Follows.follower": %w`, err)}
		}
	}
	return nil
}

func (fuo *FollowsUpdateOne) sqlSave(ctx context.Context) (_node *Follows, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   follows.Table,
			Columns: follows.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: follows.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Follows.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, follows.FieldID)
		for _, f := range fields {
			if !follows.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != follows.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Followedby(); ok {
		_spec.SetField(follows.FieldFollowedby, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.AddedFollowedby(); ok {
		_spec.AddField(follows.FieldFollowedby, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.Follower(); ok {
		_spec.SetField(follows.FieldFollower, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.AddedFollower(); ok {
		_spec.AddField(follows.FieldFollower, field.TypeInt, value)
	}
	_node = &Follows{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{follows.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}