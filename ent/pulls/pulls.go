// Code generated by ent, DO NOT EDIT.

package pulls

const (
	// Label holds the string label denoting the pulls type in the database.
	Label = "pulls"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOwner holds the string denoting the owner field in the database.
	FieldOwner = "owner"
	// FieldPrID holds the string denoting the prid field in the database.
	FieldPrID = "pr_id"
	// FieldRepo holds the string denoting the repo field in the database.
	FieldRepo = "repo"
	// FieldRepoID holds the string denoting the repoid field in the database.
	FieldRepoID = "repo_id"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldMergedAt holds the string denoting the mergedat field in the database.
	FieldMergedAt = "merged_at"
	// FieldCheckRunID holds the string denoting the checkrunid field in the database.
	FieldCheckRunID = "check_run_id"
	// FieldCheckRunResult holds the string denoting the checkrunresult field in the database.
	FieldCheckRunResult = "check_run_result"
	// FieldHeadSha holds the string denoting the headsha field in the database.
	FieldHeadSha = "head_sha"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "Creator"
	// Table holds the table name of the pulls in the database.
	Table = "pulls"
	// CreatorTable is the table that holds the Creator relation/edge.
	CreatorTable = "pulls"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "users"
	// CreatorColumn is the table column denoting the Creator relation/edge.
	CreatorColumn = "user_pull_requests"
)

// Columns holds all SQL columns for pulls fields.
var Columns = []string{
	FieldID,
	FieldOwner,
	FieldPrID,
	FieldRepo,
	FieldRepoID,
	FieldNumber,
	FieldComment,
	FieldCreatedAt,
	FieldMergedAt,
	FieldCheckRunID,
	FieldCheckRunResult,
	FieldHeadSha,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "pulls"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_pull_requests",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultPrID holds the default value on creation for the "prID" field.
	DefaultPrID int64
	// DefaultRepoID holds the default value on creation for the "repoID" field.
	DefaultRepoID int64
	// DefaultComment holds the default value on creation for the "comment" field.
	DefaultComment int64
	// DefaultCheckRunID holds the default value on creation for the "checkRunID" field.
	DefaultCheckRunID int64
	// DefaultCheckRunResult holds the default value on creation for the "checkRunResult" field.
	DefaultCheckRunResult string
	// DefaultHeadSha holds the default value on creation for the "headSha" field.
	DefaultHeadSha string
)
