// Code generated by entc, DO NOT EDIT.

package messagewithstrings

const (
	// Label holds the string label denoting the messagewithstrings type in the database.
	Label = "message_with_strings"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStrings holds the string denoting the strings field in the database.
	FieldStrings = "strings"
	// Table holds the table name of the messagewithstrings in the database.
	Table = "message_with_strings"
)

// Columns holds all SQL columns for messagewithstrings fields.
var Columns = []string{
	FieldID,
	FieldStrings,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}