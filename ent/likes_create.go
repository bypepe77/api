// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bypepe77/api/ent/likes"
	"github.com/bypepe77/api/ent/post"
	"github.com/bypepe77/api/ent/user"
)

// LikesCreate is the builder for creating a Likes entity.
type LikesCreate struct {
	config
	mutation *LikesMutation
	hooks    []Hook
}

// SetLikedAt sets the "liked_at" field.
func (lc *LikesCreate) SetLikedAt(t time.Time) *LikesCreate {
	lc.mutation.SetLikedAt(t)
	return lc
}

// SetNillableLikedAt sets the "liked_at" field if the given value is not nil.
func (lc *LikesCreate) SetNillableLikedAt(t *time.Time) *LikesCreate {
	if t != nil {
		lc.SetLikedAt(*t)
	}
	return lc
}

// SetUserID sets the "user_id" field.
func (lc *LikesCreate) SetUserID(i int) *LikesCreate {
	lc.mutation.SetUserID(i)
	return lc
}

// SetPostID sets the "post_id" field.
func (lc *LikesCreate) SetPostID(i int) *LikesCreate {
	lc.mutation.SetPostID(i)
	return lc
}

// SetUser sets the "user" edge to the User entity.
func (lc *LikesCreate) SetUser(u *User) *LikesCreate {
	return lc.SetUserID(u.ID)
}

// SetTweetID sets the "tweet" edge to the Post entity by ID.
func (lc *LikesCreate) SetTweetID(id int) *LikesCreate {
	lc.mutation.SetTweetID(id)
	return lc
}

// SetTweet sets the "tweet" edge to the Post entity.
func (lc *LikesCreate) SetTweet(p *Post) *LikesCreate {
	return lc.SetTweetID(p.ID)
}

// Mutation returns the LikesMutation object of the builder.
func (lc *LikesCreate) Mutation() *LikesMutation {
	return lc.mutation
}

// Save creates the Likes in the database.
func (lc *LikesCreate) Save(ctx context.Context) (*Likes, error) {
	var (
		err  error
		node *Likes
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LikesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Likes)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LikesMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LikesCreate) SaveX(ctx context.Context) *Likes {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LikesCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LikesCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LikesCreate) defaults() {
	if _, ok := lc.mutation.LikedAt(); !ok {
		v := likes.DefaultLikedAt()
		lc.mutation.SetLikedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LikesCreate) check() error {
	if _, ok := lc.mutation.LikedAt(); !ok {
		return &ValidationError{Name: "liked_at", err: errors.New(`ent: missing required field "Likes.liked_at"`)}
	}
	if _, ok := lc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Likes.user_id"`)}
	}
	if _, ok := lc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post_id", err: errors.New(`ent: missing required field "Likes.post_id"`)}
	}
	if _, ok := lc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Likes.user"`)}
	}
	if _, ok := lc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet", err: errors.New(`ent: missing required edge "Likes.tweet"`)}
	}
	return nil
}

func (lc *LikesCreate) sqlSave(ctx context.Context) (*Likes, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (lc *LikesCreate) createSpec() (*Likes, *sqlgraph.CreateSpec) {
	var (
		_node = &Likes{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: likes.Table,
		}
	)
	if value, ok := lc.mutation.LikedAt(); ok {
		_spec.SetField(likes.FieldLikedAt, field.TypeTime, value)
		_node.LikedAt = value
	}
	if nodes := lc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   likes.UserTable,
			Columns: []string{likes.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   likes.TweetTable,
			Columns: []string{likes.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PostID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LikesCreateBulk is the builder for creating many Likes entities in bulk.
type LikesCreateBulk struct {
	config
	builders []*LikesCreate
}

// Save creates the Likes entities in the database.
func (lcb *LikesCreateBulk) Save(ctx context.Context) ([]*Likes, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Likes, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LikesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LikesCreateBulk) SaveX(ctx context.Context) []*Likes {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LikesCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LikesCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
