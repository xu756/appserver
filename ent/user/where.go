// Code generated by ent, DO NOT EDIT.

package user

import (
	"server/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeleted, v))
}

// Creator applies equality check predicate on the "creator" field. It's identical to CreatorEQ.
func Creator(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreator, v))
}

// Editor applies equality check predicate on the "editor" field. It's identical to EditorEQ.
func Editor(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEditor, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldVersion, v))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUUID, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatar, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeleted, v))
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDeleted, v))
}

// CreatorEQ applies the EQ predicate on the "creator" field.
func CreatorEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreator, v))
}

// CreatorNEQ applies the NEQ predicate on the "creator" field.
func CreatorNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreator, v))
}

// CreatorIn applies the In predicate on the "creator" field.
func CreatorIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreator, vs...))
}

// CreatorNotIn applies the NotIn predicate on the "creator" field.
func CreatorNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreator, vs...))
}

// CreatorGT applies the GT predicate on the "creator" field.
func CreatorGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreator, v))
}

// CreatorGTE applies the GTE predicate on the "creator" field.
func CreatorGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreator, v))
}

// CreatorLT applies the LT predicate on the "creator" field.
func CreatorLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreator, v))
}

// CreatorLTE applies the LTE predicate on the "creator" field.
func CreatorLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreator, v))
}

// EditorEQ applies the EQ predicate on the "editor" field.
func EditorEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEditor, v))
}

// EditorNEQ applies the NEQ predicate on the "editor" field.
func EditorNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEditor, v))
}

// EditorIn applies the In predicate on the "editor" field.
func EditorIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldEditor, vs...))
}

// EditorNotIn applies the NotIn predicate on the "editor" field.
func EditorNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEditor, vs...))
}

// EditorGT applies the GT predicate on the "editor" field.
func EditorGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldEditor, v))
}

// EditorGTE applies the GTE predicate on the "editor" field.
func EditorGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEditor, v))
}

// EditorLT applies the LT predicate on the "editor" field.
func EditorLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldEditor, v))
}

// EditorLTE applies the LTE predicate on the "editor" field.
func EditorLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEditor, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldVersion, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUUID, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldMobile, v))
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAvatar, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
