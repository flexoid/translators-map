// Code generated by ent, DO NOT EDIT.

package translator

const (
	// Label holds the string label denoting the translator type in the database.
	Label = "translator"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldContacts holds the string denoting the contacts field in the database.
	FieldContacts = "contacts"
	// FieldDetailsURL holds the string denoting the details_url field in the database.
	FieldDetailsURL = "details_url"
	// FieldLatitude holds the string denoting the latitude field in the database.
	FieldLatitude = "latitude"
	// FieldLongitude holds the string denoting the longitude field in the database.
	FieldLongitude = "longitude"
	// Table holds the table name of the translator in the database.
	Table = "translators"
)

// Columns holds all SQL columns for translator fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldLanguage,
	FieldAddress,
	FieldContacts,
	FieldDetailsURL,
	FieldLatitude,
	FieldLongitude,
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
