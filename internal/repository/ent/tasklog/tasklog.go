// Code generated by ent, DO NOT EDIT.

package tasklog

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tasklog type in the database.
	Label = "task_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// EdgeTaskRecords holds the string denoting the task_records edge name in mutations.
	EdgeTaskRecords = "task_records"
	// Table holds the table name of the tasklog in the database.
	Table = "task_logs"
	// TaskRecordsTable is the table that holds the task_records relation/edge. The primary key declared below.
	TaskRecordsTable = "task_record_task_logs"
	// TaskRecordsInverseTable is the table name for the TaskRecord entity.
	// It exists in this package in order to avoid circular dependency with the "taskrecord" package.
	TaskRecordsInverseTable = "task_records"
)

// Columns holds all SQL columns for tasklog fields.
var Columns = []string{
	FieldID,
	FieldMessage,
}

var (
	// TaskRecordsPrimaryKey and TaskRecordsColumn2 are the table columns denoting the
	// primary key for the task_records relation (M2M).
	TaskRecordsPrimaryKey = []string{"task_record_id", "task_log_id"}
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
	// DefaultMessage holds the default value on creation for the "message" field.
	DefaultMessage string
)

// OrderOption defines the ordering options for the TaskLog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByTaskRecordsCount orders the results by task_records count.
func ByTaskRecordsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTaskRecordsStep(), opts...)
	}
}

// ByTaskRecords orders the results by task_records terms.
func ByTaskRecords(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTaskRecordsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTaskRecordsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TaskRecordsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TaskRecordsTable, TaskRecordsPrimaryKey...),
	)
}