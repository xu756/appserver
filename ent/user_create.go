// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetDeleted sets the "deleted" field.
func (uc *UserCreate) SetDeleted(b bool) *UserCreate {
	uc.mutation.SetDeleted(b)
	return uc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeleted(b *bool) *UserCreate {
	if b != nil {
		uc.SetDeleted(*b)
	}
	return uc
}

// SetCreator sets the "creator" field.
func (uc *UserCreate) SetCreator(i int64) *UserCreate {
	uc.mutation.SetCreator(i)
	return uc
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreator(i *int64) *UserCreate {
	if i != nil {
		uc.SetCreator(*i)
	}
	return uc
}

// SetEditor sets the "editor" field.
func (uc *UserCreate) SetEditor(i int64) *UserCreate {
	uc.mutation.SetEditor(i)
	return uc
}

// SetNillableEditor sets the "editor" field if the given value is not nil.
func (uc *UserCreate) SetNillableEditor(i *int64) *UserCreate {
	if i != nil {
		uc.SetEditor(*i)
	}
	return uc
}

// SetVersion sets the "version" field.
func (uc *UserCreate) SetVersion(i int64) *UserCreate {
	uc.mutation.SetVersion(i)
	return uc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (uc *UserCreate) SetNillableVersion(i *int64) *UserCreate {
	if i != nil {
		uc.SetVersion(*i)
	}
	return uc
}

// SetUUID sets the "uuid" field.
func (uc *UserCreate) SetUUID(s string) *UserCreate {
	uc.mutation.SetUUID(s)
	return uc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (uc *UserCreate) SetNillableUUID(s *string) *UserCreate {
	if s != nil {
		uc.SetUUID(*s)
	}
	return uc
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetMobile sets the "mobile" field.
func (uc *UserCreate) SetMobile(s string) *UserCreate {
	uc.mutation.SetMobile(s)
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatar(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatar(*s)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int64) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.Deleted(); !ok {
		v := user.DefaultDeleted
		uc.mutation.SetDeleted(v)
	}
	if _, ok := uc.mutation.Creator(); !ok {
		v := user.DefaultCreator
		uc.mutation.SetCreator(v)
	}
	if _, ok := uc.mutation.Editor(); !ok {
		v := user.DefaultEditor
		uc.mutation.SetEditor(v)
	}
	if _, ok := uc.mutation.Version(); !ok {
		v := user.DefaultVersion
		uc.mutation.SetVersion(v)
	}
	if _, ok := uc.mutation.UUID(); !ok {
		v := user.DefaultUUID()
		uc.mutation.SetUUID(v)
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		v := user.DefaultAvatar
		uc.mutation.SetAvatar(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "User.deleted"`)}
	}
	if _, ok := uc.mutation.Creator(); !ok {
		return &ValidationError{Name: "creator", err: errors.New(`ent: missing required field "User.creator"`)}
	}
	if _, ok := uc.mutation.Editor(); !ok {
		return &ValidationError{Name: "editor", err: errors.New(`ent: missing required field "User.editor"`)}
	}
	if _, ok := uc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "User.version"`)}
	}
	if _, ok := uc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "User.uuid"`)}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if _, ok := uc.mutation.Mobile(); !ok {
		return &ValidationError{Name: "mobile", err: errors.New(`ent: missing required field "User.mobile"`)}
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "User.avatar"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.Deleted(); ok {
		_spec.SetField(user.FieldDeleted, field.TypeBool, value)
		_node.Deleted = value
	}
	if value, ok := uc.mutation.Creator(); ok {
		_spec.SetField(user.FieldCreator, field.TypeInt64, value)
		_node.Creator = value
	}
	if value, ok := uc.mutation.Editor(); ok {
		_spec.SetField(user.FieldEditor, field.TypeInt64, value)
		_node.Editor = value
	}
	if value, ok := uc.mutation.Version(); ok {
		_spec.SetField(user.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := uc.mutation.UUID(); ok {
		_spec.SetField(user.FieldUUID, field.TypeString, value)
		_node.UUID = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
