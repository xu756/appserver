// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "creator", Type: field.TypeInt64, Default: 0},
		{Name: "editor", Type: field.TypeInt64, Default: 1},
		{Name: "version", Type: field.TypeInt64, Default: 0},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "parent_id", Type: field.TypeInt64, Default: 0},
		{Name: "level", Type: field.TypeInt64, Default: 0},
		{Name: "name", Type: field.TypeString},
		{Name: "intro", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "creator", Type: field.TypeInt64, Default: 0},
		{Name: "editor", Type: field.TypeInt64, Default: 1},
		{Name: "version", Type: field.TypeInt64, Default: 0},
		{Name: "parent_id", Type: field.TypeInt64, Default: 0},
		{Name: "level", Type: field.TypeInt64, Default: 0},
		{Name: "role_name", Type: field.TypeString},
		{Name: "intro", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "creator", Type: field.TypeInt64, Default: 0},
		{Name: "editor", Type: field.TypeInt64, Default: 1},
		{Name: "version", Type: field.TypeInt64, Default: 0},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "mobile", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString, Default: "https://cos.imlogic.cn/appadmin/images/avatar.jpeg"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_uuid",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[7]},
			},
			{
				Name:    "user_username_mobile",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[8], UsersColumns[10]},
			},
		},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "creator", Type: field.TypeInt64, Default: 0},
		{Name: "editor", Type: field.TypeInt64, Default: 1},
		{Name: "version", Type: field.TypeInt64, Default: 0},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "group_id", Type: field.TypeInt64},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0]},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "creator", Type: field.TypeInt64, Default: 0},
		{Name: "editor", Type: field.TypeInt64, Default: 1},
		{Name: "version", Type: field.TypeInt64, Default: 0},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "role_id", Type: field.TypeInt64},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupsTable,
		RolesTable,
		UsersTable,
		UserGroupsTable,
		UserRolesTable,
	}
)

func init() {
}
