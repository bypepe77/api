// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/bypepe77/api/ent/post"
	"github.com/bypepe77/api/ent/predicate"
	"github.com/bypepe77/api/ent/user"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetAuthorID sets the "author_id" field.
func (pu *PostUpdate) SetAuthorID(i int) *PostUpdate {
	pu.mutation.SetAuthorID(i)
	return pu
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillableAuthorID(i *int) *PostUpdate {
	if i != nil {
		pu.SetAuthorID(*i)
	}
	return pu
}

// ClearAuthorID clears the value of the "author_id" field.
func (pu *PostUpdate) ClearAuthorID() *PostUpdate {
	pu.mutation.ClearAuthorID()
	return pu
}

// SetContent sets the "content" field.
func (pu *PostUpdate) SetContent(s string) *PostUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetHashtags sets the "hashtags" field.
func (pu *PostUpdate) SetHashtags(s []string) *PostUpdate {
	pu.mutation.SetHashtags(s)
	return pu
}

// AppendHashtags appends s to the "hashtags" field.
func (pu *PostUpdate) AppendHashtags(s []string) *PostUpdate {
	pu.mutation.AppendHashtags(s)
	return pu
}

// ClearHashtags clears the value of the "hashtags" field.
func (pu *PostUpdate) ClearHashtags() *PostUpdate {
	pu.mutation.ClearHashtags()
	return pu
}

// SetLinks sets the "links" field.
func (pu *PostUpdate) SetLinks(s []string) *PostUpdate {
	pu.mutation.SetLinks(s)
	return pu
}

// AppendLinks appends s to the "links" field.
func (pu *PostUpdate) AppendLinks(s []string) *PostUpdate {
	pu.mutation.AppendLinks(s)
	return pu
}

// ClearLinks clears the value of the "links" field.
func (pu *PostUpdate) ClearLinks() *PostUpdate {
	pu.mutation.ClearLinks()
	return pu
}

// SetShared sets the "shared" field.
func (pu *PostUpdate) SetShared(b bool) *PostUpdate {
	pu.mutation.SetShared(b)
	return pu
}

// SetNillableShared sets the "shared" field if the given value is not nil.
func (pu *PostUpdate) SetNillableShared(b *bool) *PostUpdate {
	if b != nil {
		pu.SetShared(*b)
	}
	return pu
}

// SetAuthor sets the "author" edge to the User entity.
func (pu *PostUpdate) SetAuthor(u *User) *PostUpdate {
	return pu.SetAuthorID(u.ID)
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (pu *PostUpdate) AddLikedUserIDs(ids ...int) *PostUpdate {
	pu.mutation.AddLikedUserIDs(ids...)
	return pu
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (pu *PostUpdate) AddLikedUsers(u ...*User) *PostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddLikedUserIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (pu *PostUpdate) ClearAuthor() *PostUpdate {
	pu.mutation.ClearAuthor()
	return pu
}

// ClearLikedUsers clears all "liked_users" edges to the User entity.
func (pu *PostUpdate) ClearLikedUsers() *PostUpdate {
	pu.mutation.ClearLikedUsers()
	return pu
}

// RemoveLikedUserIDs removes the "liked_users" edge to User entities by IDs.
func (pu *PostUpdate) RemoveLikedUserIDs(ids ...int) *PostUpdate {
	pu.mutation.RemoveLikedUserIDs(ids...)
	return pu
}

// RemoveLikedUsers removes "liked_users" edges to User entities.
func (pu *PostUpdate) RemoveLikedUsers(u ...*User) *PostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveLikedUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PostUpdate) check() error {
	if v, ok := pu.mutation.Content(); ok {
		if err := post.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Post.content": %w`, err)}
		}
	}
	return nil
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: post.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if value, ok := pu.mutation.Hashtags(); ok {
		_spec.SetField(post.FieldHashtags, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedHashtags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, post.FieldHashtags, value)
		})
	}
	if pu.mutation.HashtagsCleared() {
		_spec.ClearField(post.FieldHashtags, field.TypeJSON)
	}
	if value, ok := pu.mutation.Links(); ok {
		_spec.SetField(post.FieldLinks, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, post.FieldLinks, value)
		})
	}
	if pu.mutation.LinksCleared() {
		_spec.ClearField(post.FieldLinks, field.TypeJSON)
	}
	if value, ok := pu.mutation.Shared(); ok {
		_spec.SetField(post.FieldShared, field.TypeBool, value)
	}
	if pu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedUsersTable,
			Columns: post.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		createE := &LikesCreate{config: pu.config, mutation: newLikesMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedLikedUsersIDs(); len(nodes) > 0 && !pu.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedUsersTable,
			Columns: post.LikedUsersPrimaryKey,
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
		createE := &LikesCreate{config: pu.config, mutation: newLikesMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedUsersTable,
			Columns: post.LikedUsersPrimaryKey,
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
		createE := &LikesCreate{config: pu.config, mutation: newLikesMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetAuthorID sets the "author_id" field.
func (puo *PostUpdateOne) SetAuthorID(i int) *PostUpdateOne {
	puo.mutation.SetAuthorID(i)
	return puo
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableAuthorID(i *int) *PostUpdateOne {
	if i != nil {
		puo.SetAuthorID(*i)
	}
	return puo
}

// ClearAuthorID clears the value of the "author_id" field.
func (puo *PostUpdateOne) ClearAuthorID() *PostUpdateOne {
	puo.mutation.ClearAuthorID()
	return puo
}

// SetContent sets the "content" field.
func (puo *PostUpdateOne) SetContent(s string) *PostUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetHashtags sets the "hashtags" field.
func (puo *PostUpdateOne) SetHashtags(s []string) *PostUpdateOne {
	puo.mutation.SetHashtags(s)
	return puo
}

// AppendHashtags appends s to the "hashtags" field.
func (puo *PostUpdateOne) AppendHashtags(s []string) *PostUpdateOne {
	puo.mutation.AppendHashtags(s)
	return puo
}

// ClearHashtags clears the value of the "hashtags" field.
func (puo *PostUpdateOne) ClearHashtags() *PostUpdateOne {
	puo.mutation.ClearHashtags()
	return puo
}

// SetLinks sets the "links" field.
func (puo *PostUpdateOne) SetLinks(s []string) *PostUpdateOne {
	puo.mutation.SetLinks(s)
	return puo
}

// AppendLinks appends s to the "links" field.
func (puo *PostUpdateOne) AppendLinks(s []string) *PostUpdateOne {
	puo.mutation.AppendLinks(s)
	return puo
}

// ClearLinks clears the value of the "links" field.
func (puo *PostUpdateOne) ClearLinks() *PostUpdateOne {
	puo.mutation.ClearLinks()
	return puo
}

// SetShared sets the "shared" field.
func (puo *PostUpdateOne) SetShared(b bool) *PostUpdateOne {
	puo.mutation.SetShared(b)
	return puo
}

// SetNillableShared sets the "shared" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableShared(b *bool) *PostUpdateOne {
	if b != nil {
		puo.SetShared(*b)
	}
	return puo
}

// SetAuthor sets the "author" edge to the User entity.
func (puo *PostUpdateOne) SetAuthor(u *User) *PostUpdateOne {
	return puo.SetAuthorID(u.ID)
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (puo *PostUpdateOne) AddLikedUserIDs(ids ...int) *PostUpdateOne {
	puo.mutation.AddLikedUserIDs(ids...)
	return puo
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (puo *PostUpdateOne) AddLikedUsers(u ...*User) *PostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddLikedUserIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (puo *PostUpdateOne) ClearAuthor() *PostUpdateOne {
	puo.mutation.ClearAuthor()
	return puo
}

// ClearLikedUsers clears all "liked_users" edges to the User entity.
func (puo *PostUpdateOne) ClearLikedUsers() *PostUpdateOne {
	puo.mutation.ClearLikedUsers()
	return puo
}

// RemoveLikedUserIDs removes the "liked_users" edge to User entities by IDs.
func (puo *PostUpdateOne) RemoveLikedUserIDs(ids ...int) *PostUpdateOne {
	puo.mutation.RemoveLikedUserIDs(ids...)
	return puo
}

// RemoveLikedUsers removes "liked_users" edges to User entities.
func (puo *PostUpdateOne) RemoveLikedUsers(u ...*User) *PostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveLikedUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	var (
		err  error
		node *Post
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Post)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PostMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PostUpdateOne) check() error {
	if v, ok := puo.mutation.Content(); ok {
		if err := post.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Post.content": %w`, err)}
		}
	}
	return nil
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: post.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if value, ok := puo.mutation.Hashtags(); ok {
		_spec.SetField(post.FieldHashtags, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedHashtags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, post.FieldHashtags, value)
		})
	}
	if puo.mutation.HashtagsCleared() {
		_spec.ClearField(post.FieldHashtags, field.TypeJSON)
	}
	if value, ok := puo.mutation.Links(); ok {
		_spec.SetField(post.FieldLinks, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedLinks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, post.FieldLinks, value)
		})
	}
	if puo.mutation.LinksCleared() {
		_spec.ClearField(post.FieldLinks, field.TypeJSON)
	}
	if value, ok := puo.mutation.Shared(); ok {
		_spec.SetField(post.FieldShared, field.TypeBool, value)
	}
	if puo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedUsersTable,
			Columns: post.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		createE := &LikesCreate{config: puo.config, mutation: newLikesMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedLikedUsersIDs(); len(nodes) > 0 && !puo.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedUsersTable,
			Columns: post.LikedUsersPrimaryKey,
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
		createE := &LikesCreate{config: puo.config, mutation: newLikesMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedUsersTable,
			Columns: post.LikedUsersPrimaryKey,
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
		createE := &LikesCreate{config: puo.config, mutation: newLikesMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
