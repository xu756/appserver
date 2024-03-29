// Code generated by ent, DO NOT EDIT.

package usergroup

import (
	"server/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldUpdatedAt, v))
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v bool) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldDeleted, v))
}

// Creator applies equality check predicate on the "creator" field. It's identical to CreatorEQ.
func Creator(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldCreator, v))
}

// Editor applies equality check predicate on the "editor" field. It's identical to EditorEQ.
func Editor(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldEditor, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldVersion, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldUserID, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldGroupID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v bool) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldDeleted, v))
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v bool) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldDeleted, v))
}

// CreatorEQ applies the EQ predicate on the "creator" field.
func CreatorEQ(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldCreator, v))
}

// CreatorNEQ applies the NEQ predicate on the "creator" field.
func CreatorNEQ(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldCreator, v))
}

// CreatorIn applies the In predicate on the "creator" field.
func CreatorIn(vs ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldCreator, vs...))
}

// CreatorNotIn applies the NotIn predicate on the "creator" field.
func CreatorNotIn(vs ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldCreator, vs...))
}

// CreatorGT applies the GT predicate on the "creator" field.
func CreatorGT(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldCreator, v))
}

// CreatorGTE applies the GTE predicate on the "creator" field.
func CreatorGTE(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldCreator, v))
}

// CreatorLT applies the LT predicate on the "creator" field.
func CreatorLT(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldCreator, v))
}

// CreatorLTE applies the LTE predicate on the "creator" field.
func CreatorLTE(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldCreator, v))
}

// EditorEQ applies the EQ predicate on the "editor" field.
func EditorEQ(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldEditor, v))
}

// EditorNEQ applies the NEQ predicate on the "editor" field.
func EditorNEQ(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldEditor, v))
}

// EditorIn applies the In predicate on the "editor" field.
func EditorIn(vs ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldEditor, vs...))
}

// EditorNotIn applies the NotIn predicate on the "editor" field.
func EditorNotIn(vs ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldEditor, vs...))
}

// EditorGT applies the GT predicate on the "editor" field.
func EditorGT(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldEditor, v))
}

// EditorGTE applies the GTE predicate on the "editor" field.
func EditorGTE(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldEditor, v))
}

// EditorLT applies the LT predicate on the "editor" field.
func EditorLT(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldEditor, v))
}

// EditorLTE applies the LTE predicate on the "editor" field.
func EditorLTE(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldEditor, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldVersion, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldUserID, v))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldGroupID, vs...))
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldGroupID, v))
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldGroupID, v))
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldGroupID, v))
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v int64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldGroupID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserGroup) predicate.UserGroup {
	return predicate.UserGroup(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserGroup) predicate.UserGroup {
	return predicate.UserGroup(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserGroup) predicate.UserGroup {
	return predicate.UserGroup(sql.NotPredicates(p))
}
