// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PullsColumns holds the columns for the "pulls" table.
	PullsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "owner", Type: field.TypeString},
		{Name: "pr_id", Type: field.TypeInt64, Default: 0},
		{Name: "repo", Type: field.TypeString},
		{Name: "repo_id", Type: field.TypeInt64, Default: 0},
		{Name: "number", Type: field.TypeInt, Unique: true},
		{Name: "comment", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "merged_at", Type: field.TypeTime, Nullable: true},
		{Name: "check_run_id", Type: field.TypeInt64, Default: 0},
		{Name: "check_run_result", Type: field.TypeString, Default: ""},
		{Name: "head_sha", Type: field.TypeString, Default: ""},
		{Name: "user_pull_requests", Type: field.TypeInt},
	}
	// PullsTable holds the schema information for the "pulls" table.
	PullsTable = &schema.Table{
		Name:       "pulls",
		Columns:    PullsColumns,
		PrimaryKey: []*schema.Column{PullsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pulls_users_pull_requests",
				Columns:    []*schema.Column{PullsColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "github_id", Type: field.TypeInt64, Unique: true},
		{Name: "bangumi_id", Type: field.TypeInt64, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_github_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PullsTable,
		UsersTable,
	}
)

func init() {
	PullsTable.ForeignKeys[0].RefTable = UsersTable
}
