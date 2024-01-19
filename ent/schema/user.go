package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"server/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Mixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).
			Unique().Immutable().Comment("用户uuid"),
		field.String("username").Comment("姓名"),
		field.String("password").Comment("密码"),
		field.String("mobile").Comment("手机号"),
		field.String("avatar").Comment("头像").Default("https://cos.imlogic.cn/appadmin/images/avatar.jpeg"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
