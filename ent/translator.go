// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/flexoid/translators-map-go/ent/translator"
)

// Translator is the model entity for the Translator schema.
type Translator struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ExternalID holds the value of the "external_id" field.
	ExternalID int `json:"external_id,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// AddressSha holds the value of the "address_sha" field.
	AddressSha []byte `json:"address_sha,omitempty"`
	// DetailsURL holds the value of the "details_url" field.
	DetailsURL string `json:"details_url,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude float64 `json:"longitude,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Translator) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case translator.FieldAddressSha:
			values[i] = new([]byte)
		case translator.FieldLatitude, translator.FieldLongitude:
			values[i] = new(sql.NullFloat64)
		case translator.FieldID, translator.FieldExternalID:
			values[i] = new(sql.NullInt64)
		case translator.FieldLanguage, translator.FieldAddress, translator.FieldDetailsURL:
			values[i] = new(sql.NullString)
		case translator.FieldCreatedAt, translator.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Translator", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Translator fields.
func (t *Translator) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case translator.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case translator.FieldExternalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field external_id", values[i])
			} else if value.Valid {
				t.ExternalID = int(value.Int64)
			}
		case translator.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				t.Language = value.String
			}
		case translator.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				t.Address = value.String
			}
		case translator.FieldAddressSha:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field address_sha", values[i])
			} else if value != nil {
				t.AddressSha = *value
			}
		case translator.FieldDetailsURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field details_url", values[i])
			} else if value.Valid {
				t.DetailsURL = value.String
			}
		case translator.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				t.Latitude = value.Float64
			}
		case translator.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				t.Longitude = value.Float64
			}
		case translator.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case translator.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Translator.
// Note that you need to call Translator.Unwrap() before calling this method if this Translator
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Translator) Update() *TranslatorUpdateOne {
	return (&TranslatorClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Translator entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Translator) Unwrap() *Translator {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Translator is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Translator) String() string {
	var builder strings.Builder
	builder.WriteString("Translator(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("external_id=")
	builder.WriteString(fmt.Sprintf("%v", t.ExternalID))
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(t.Language)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(t.Address)
	builder.WriteString(", ")
	builder.WriteString("address_sha=")
	builder.WriteString(fmt.Sprintf("%v", t.AddressSha))
	builder.WriteString(", ")
	builder.WriteString("details_url=")
	builder.WriteString(t.DetailsURL)
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(fmt.Sprintf("%v", t.Latitude))
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(fmt.Sprintf("%v", t.Longitude))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Translators is a parsable slice of Translator.
type Translators []*Translator

func (t Translators) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
