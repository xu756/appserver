// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/ent/predicate"
	"server/ent/role"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RoleUpdate) SetUpdatedAt(t time.Time) *RoleUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetDeleted sets the "deleted" field.
func (ru *RoleUpdate) SetDeleted(b bool) *RoleUpdate {
	ru.mutation.SetDeleted(b)
	return ru
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableDeleted(b *bool) *RoleUpdate {
	if b != nil {
		ru.SetDeleted(*b)
	}
	return ru
}

// SetCreator sets the "creator" field.
func (ru *RoleUpdate) SetCreator(i int64) *RoleUpdate {
	ru.mutation.ResetCreator()
	ru.mutation.SetCreator(i)
	return ru
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableCreator(i *int64) *RoleUpdate {
	if i != nil {
		ru.SetCreator(*i)
	}
	return ru
}

// AddCreator adds i to the "creator" field.
func (ru *RoleUpdate) AddCreator(i int64) *RoleUpdate {
	ru.mutation.AddCreator(i)
	return ru
}

// SetEditor sets the "editor" field.
func (ru *RoleUpdate) SetEditor(i int64) *RoleUpdate {
	ru.mutation.ResetEditor()
	ru.mutation.SetEditor(i)
	return ru
}

// SetNillableEditor sets the "editor" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableEditor(i *int64) *RoleUpdate {
	if i != nil {
		ru.SetEditor(*i)
	}
	return ru
}

// AddEditor adds i to the "editor" field.
func (ru *RoleUpdate) AddEditor(i int64) *RoleUpdate {
	ru.mutation.AddEditor(i)
	return ru
}

// SetVersion sets the "version" field.
func (ru *RoleUpdate) SetVersion(i int64) *RoleUpdate {
	ru.mutation.ResetVersion()
	ru.mutation.SetVersion(i)
	return ru
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableVersion(i *int64) *RoleUpdate {
	if i != nil {
		ru.SetVersion(*i)
	}
	return ru
}

// AddVersion adds i to the "version" field.
func (ru *RoleUpdate) AddVersion(i int64) *RoleUpdate {
	ru.mutation.AddVersion(i)
	return ru
}

// SetParentID sets the "parent_id" field.
func (ru *RoleUpdate) SetParentID(i int64) *RoleUpdate {
	ru.mutation.ResetParentID()
	ru.mutation.SetParentID(i)
	return ru
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableParentID(i *int64) *RoleUpdate {
	if i != nil {
		ru.SetParentID(*i)
	}
	return ru
}

// AddParentID adds i to the "parent_id" field.
func (ru *RoleUpdate) AddParentID(i int64) *RoleUpdate {
	ru.mutation.AddParentID(i)
	return ru
}

// SetLevel sets the "level" field.
func (ru *RoleUpdate) SetLevel(i int64) *RoleUpdate {
	ru.mutation.ResetLevel()
	ru.mutation.SetLevel(i)
	return ru
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableLevel(i *int64) *RoleUpdate {
	if i != nil {
		ru.SetLevel(*i)
	}
	return ru
}

// AddLevel adds i to the "level" field.
func (ru *RoleUpdate) AddLevel(i int64) *RoleUpdate {
	ru.mutation.AddLevel(i)
	return ru
}

// SetRoleName sets the "role_name" field.
func (ru *RoleUpdate) SetRoleName(s string) *RoleUpdate {
	ru.mutation.SetRoleName(s)
	return ru
}

// SetNillableRoleName sets the "role_name" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableRoleName(s *string) *RoleUpdate {
	if s != nil {
		ru.SetRoleName(*s)
	}
	return ru
}

// SetIntro sets the "intro" field.
func (ru *RoleUpdate) SetIntro(s string) *RoleUpdate {
	ru.mutation.SetIntro(s)
	return ru
}

// SetNillableIntro sets the "intro" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableIntro(s *string) *RoleUpdate {
	if s != nil {
		ru.SetIntro(*s)
	}
	return ru
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RoleUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := role.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoleUpdate) check() error {
	if v, ok := ru.mutation.RoleName(); ok {
		if err := role.RoleNameValidator(v); err != nil {
			return &ValidationError{Name: "role_name", err: fmt.Errorf(`ent: validator failed for field "Role.role_name": %w`, err)}
		}
	}
	return nil
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Deleted(); ok {
		_spec.SetField(role.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := ru.mutation.Creator(); ok {
		_spec.SetField(role.FieldCreator, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedCreator(); ok {
		_spec.AddField(role.FieldCreator, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.Editor(); ok {
		_spec.SetField(role.FieldEditor, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedEditor(); ok {
		_spec.AddField(role.FieldEditor, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.Version(); ok {
		_spec.SetField(role.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedVersion(); ok {
		_spec.AddField(role.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.ParentID(); ok {
		_spec.SetField(role.FieldParentID, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedParentID(); ok {
		_spec.AddField(role.FieldParentID, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.Level(); ok {
		_spec.SetField(role.FieldLevel, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedLevel(); ok {
		_spec.AddField(role.FieldLevel, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.RoleName(); ok {
		_spec.SetField(role.FieldRoleName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Intro(); ok {
		_spec.SetField(role.FieldIntro, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RoleUpdateOne) SetUpdatedAt(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetDeleted sets the "deleted" field.
func (ruo *RoleUpdateOne) SetDeleted(b bool) *RoleUpdateOne {
	ruo.mutation.SetDeleted(b)
	return ruo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableDeleted(b *bool) *RoleUpdateOne {
	if b != nil {
		ruo.SetDeleted(*b)
	}
	return ruo
}

// SetCreator sets the "creator" field.
func (ruo *RoleUpdateOne) SetCreator(i int64) *RoleUpdateOne {
	ruo.mutation.ResetCreator()
	ruo.mutation.SetCreator(i)
	return ruo
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableCreator(i *int64) *RoleUpdateOne {
	if i != nil {
		ruo.SetCreator(*i)
	}
	return ruo
}

// AddCreator adds i to the "creator" field.
func (ruo *RoleUpdateOne) AddCreator(i int64) *RoleUpdateOne {
	ruo.mutation.AddCreator(i)
	return ruo
}

// SetEditor sets the "editor" field.
func (ruo *RoleUpdateOne) SetEditor(i int64) *RoleUpdateOne {
	ruo.mutation.ResetEditor()
	ruo.mutation.SetEditor(i)
	return ruo
}

// SetNillableEditor sets the "editor" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableEditor(i *int64) *RoleUpdateOne {
	if i != nil {
		ruo.SetEditor(*i)
	}
	return ruo
}

// AddEditor adds i to the "editor" field.
func (ruo *RoleUpdateOne) AddEditor(i int64) *RoleUpdateOne {
	ruo.mutation.AddEditor(i)
	return ruo
}

// SetVersion sets the "version" field.
func (ruo *RoleUpdateOne) SetVersion(i int64) *RoleUpdateOne {
	ruo.mutation.ResetVersion()
	ruo.mutation.SetVersion(i)
	return ruo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableVersion(i *int64) *RoleUpdateOne {
	if i != nil {
		ruo.SetVersion(*i)
	}
	return ruo
}

// AddVersion adds i to the "version" field.
func (ruo *RoleUpdateOne) AddVersion(i int64) *RoleUpdateOne {
	ruo.mutation.AddVersion(i)
	return ruo
}

// SetParentID sets the "parent_id" field.
func (ruo *RoleUpdateOne) SetParentID(i int64) *RoleUpdateOne {
	ruo.mutation.ResetParentID()
	ruo.mutation.SetParentID(i)
	return ruo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableParentID(i *int64) *RoleUpdateOne {
	if i != nil {
		ruo.SetParentID(*i)
	}
	return ruo
}

// AddParentID adds i to the "parent_id" field.
func (ruo *RoleUpdateOne) AddParentID(i int64) *RoleUpdateOne {
	ruo.mutation.AddParentID(i)
	return ruo
}

// SetLevel sets the "level" field.
func (ruo *RoleUpdateOne) SetLevel(i int64) *RoleUpdateOne {
	ruo.mutation.ResetLevel()
	ruo.mutation.SetLevel(i)
	return ruo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableLevel(i *int64) *RoleUpdateOne {
	if i != nil {
		ruo.SetLevel(*i)
	}
	return ruo
}

// AddLevel adds i to the "level" field.
func (ruo *RoleUpdateOne) AddLevel(i int64) *RoleUpdateOne {
	ruo.mutation.AddLevel(i)
	return ruo
}

// SetRoleName sets the "role_name" field.
func (ruo *RoleUpdateOne) SetRoleName(s string) *RoleUpdateOne {
	ruo.mutation.SetRoleName(s)
	return ruo
}

// SetNillableRoleName sets the "role_name" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableRoleName(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetRoleName(*s)
	}
	return ruo
}

// SetIntro sets the "intro" field.
func (ruo *RoleUpdateOne) SetIntro(s string) *RoleUpdateOne {
	ruo.mutation.SetIntro(s)
	return ruo
}

// SetNillableIntro sets the "intro" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableIntro(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetIntro(*s)
	}
	return ruo
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ruo *RoleUpdateOne) Where(ps ...predicate.Role) *RoleUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoleUpdateOne) Select(field string, fields ...string) *RoleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RoleUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := role.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoleUpdateOne) check() error {
	if v, ok := ruo.mutation.RoleName(); ok {
		if err := role.RoleNameValidator(v); err != nil {
			return &ValidationError{Name: "role_name", err: fmt.Errorf(`ent: validator failed for field "Role.role_name": %w`, err)}
		}
	}
	return nil
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Role.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for _, f := range fields {
			if !role.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Deleted(); ok {
		_spec.SetField(role.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.Creator(); ok {
		_spec.SetField(role.FieldCreator, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedCreator(); ok {
		_spec.AddField(role.FieldCreator, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.Editor(); ok {
		_spec.SetField(role.FieldEditor, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedEditor(); ok {
		_spec.AddField(role.FieldEditor, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.Version(); ok {
		_spec.SetField(role.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedVersion(); ok {
		_spec.AddField(role.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.ParentID(); ok {
		_spec.SetField(role.FieldParentID, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedParentID(); ok {
		_spec.AddField(role.FieldParentID, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.Level(); ok {
		_spec.SetField(role.FieldLevel, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedLevel(); ok {
		_spec.AddField(role.FieldLevel, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.RoleName(); ok {
		_spec.SetField(role.FieldRoleName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Intro(); ok {
		_spec.SetField(role.FieldIntro, field.TypeString, value)
	}
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
