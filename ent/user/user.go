// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"

	// EdgeResponses holds the string denoting the responses edge name in mutations.
	EdgeResponses = "responses"
	// EdgeQuizes holds the string denoting the quizes edge name in mutations.
	EdgeQuizes = "quizes"

	// Table holds the table name of the user in the database.
	Table = "users"
	// ResponsesTable is the table the holds the responses relation/edge.
	ResponsesTable = "responses"
	// ResponsesInverseTable is the table name for the Response entity.
	// It exists in this package in order to avoid circular dependency with the "response" package.
	ResponsesInverseTable = "responses"
	// ResponsesColumn is the table column denoting the responses relation/edge.
	ResponsesColumn = "user_responses"
	// QuizesTable is the table the holds the quizes relation/edge. The primary key declared below.
	QuizesTable = "user_quizes"
	// QuizesInverseTable is the table name for the Quiz entity.
	// It exists in this package in order to avoid circular dependency with the "quiz" package.
	QuizesInverseTable = "quizs"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldCreatedAt,
}

var (
	// QuizesPrimaryKey and QuizesColumn2 are the table columns denoting the
	// primary key for the quizes relation (M2M).
	QuizesPrimaryKey = []string{"user_id", "quiz_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
