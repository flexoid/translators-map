// Code generated by ent, DO NOT EDIT.

package translator

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/flexoid/translators-map-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ExternalID applies equality check predicate on the "external_id" field. It's identical to ExternalIDEQ.
func ExternalID(v int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExternalID), v))
	})
}

// Language applies equality check predicate on the "language" field. It's identical to LanguageEQ.
func Language(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLanguage), v))
	})
}

// AddressSha applies equality check predicate on the "address_sha" field. It's identical to AddressShaEQ.
func AddressSha(v []byte) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddressSha), v))
	})
}

// DetailsURL applies equality check predicate on the "details_url" field. It's identical to DetailsURLEQ.
func DetailsURL(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetailsURL), v))
	})
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatitude), v))
	})
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitude), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// ExternalIDEQ applies the EQ predicate on the "external_id" field.
func ExternalIDEQ(v int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExternalID), v))
	})
}

// ExternalIDNEQ applies the NEQ predicate on the "external_id" field.
func ExternalIDNEQ(v int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExternalID), v))
	})
}

// ExternalIDIn applies the In predicate on the "external_id" field.
func ExternalIDIn(vs ...int) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExternalID), v...))
	})
}

// ExternalIDNotIn applies the NotIn predicate on the "external_id" field.
func ExternalIDNotIn(vs ...int) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExternalID), v...))
	})
}

// ExternalIDGT applies the GT predicate on the "external_id" field.
func ExternalIDGT(v int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExternalID), v))
	})
}

// ExternalIDGTE applies the GTE predicate on the "external_id" field.
func ExternalIDGTE(v int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExternalID), v))
	})
}

// ExternalIDLT applies the LT predicate on the "external_id" field.
func ExternalIDLT(v int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExternalID), v))
	})
}

// ExternalIDLTE applies the LTE predicate on the "external_id" field.
func ExternalIDLTE(v int) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExternalID), v))
	})
}

// LanguageEQ applies the EQ predicate on the "language" field.
func LanguageEQ(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLanguage), v))
	})
}

// LanguageNEQ applies the NEQ predicate on the "language" field.
func LanguageNEQ(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLanguage), v))
	})
}

// LanguageIn applies the In predicate on the "language" field.
func LanguageIn(vs ...string) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLanguage), v...))
	})
}

// LanguageNotIn applies the NotIn predicate on the "language" field.
func LanguageNotIn(vs ...string) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLanguage), v...))
	})
}

// LanguageGT applies the GT predicate on the "language" field.
func LanguageGT(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLanguage), v))
	})
}

// LanguageGTE applies the GTE predicate on the "language" field.
func LanguageGTE(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLanguage), v))
	})
}

// LanguageLT applies the LT predicate on the "language" field.
func LanguageLT(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLanguage), v))
	})
}

// LanguageLTE applies the LTE predicate on the "language" field.
func LanguageLTE(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLanguage), v))
	})
}

// LanguageContains applies the Contains predicate on the "language" field.
func LanguageContains(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLanguage), v))
	})
}

// LanguageHasPrefix applies the HasPrefix predicate on the "language" field.
func LanguageHasPrefix(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLanguage), v))
	})
}

// LanguageHasSuffix applies the HasSuffix predicate on the "language" field.
func LanguageHasSuffix(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLanguage), v))
	})
}

// LanguageEqualFold applies the EqualFold predicate on the "language" field.
func LanguageEqualFold(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLanguage), v))
	})
}

// LanguageContainsFold applies the ContainsFold predicate on the "language" field.
func LanguageContainsFold(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLanguage), v))
	})
}

// AddressShaEQ applies the EQ predicate on the "address_sha" field.
func AddressShaEQ(v []byte) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddressSha), v))
	})
}

// AddressShaNEQ applies the NEQ predicate on the "address_sha" field.
func AddressShaNEQ(v []byte) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddressSha), v))
	})
}

// AddressShaIn applies the In predicate on the "address_sha" field.
func AddressShaIn(vs ...[]byte) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAddressSha), v...))
	})
}

// AddressShaNotIn applies the NotIn predicate on the "address_sha" field.
func AddressShaNotIn(vs ...[]byte) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAddressSha), v...))
	})
}

// AddressShaGT applies the GT predicate on the "address_sha" field.
func AddressShaGT(v []byte) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddressSha), v))
	})
}

// AddressShaGTE applies the GTE predicate on the "address_sha" field.
func AddressShaGTE(v []byte) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddressSha), v))
	})
}

// AddressShaLT applies the LT predicate on the "address_sha" field.
func AddressShaLT(v []byte) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddressSha), v))
	})
}

// AddressShaLTE applies the LTE predicate on the "address_sha" field.
func AddressShaLTE(v []byte) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddressSha), v))
	})
}

// DetailsURLEQ applies the EQ predicate on the "details_url" field.
func DetailsURLEQ(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetailsURL), v))
	})
}

// DetailsURLNEQ applies the NEQ predicate on the "details_url" field.
func DetailsURLNEQ(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetailsURL), v))
	})
}

// DetailsURLIn applies the In predicate on the "details_url" field.
func DetailsURLIn(vs ...string) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDetailsURL), v...))
	})
}

// DetailsURLNotIn applies the NotIn predicate on the "details_url" field.
func DetailsURLNotIn(vs ...string) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDetailsURL), v...))
	})
}

// DetailsURLGT applies the GT predicate on the "details_url" field.
func DetailsURLGT(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetailsURL), v))
	})
}

// DetailsURLGTE applies the GTE predicate on the "details_url" field.
func DetailsURLGTE(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetailsURL), v))
	})
}

// DetailsURLLT applies the LT predicate on the "details_url" field.
func DetailsURLLT(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetailsURL), v))
	})
}

// DetailsURLLTE applies the LTE predicate on the "details_url" field.
func DetailsURLLTE(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetailsURL), v))
	})
}

// DetailsURLContains applies the Contains predicate on the "details_url" field.
func DetailsURLContains(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetailsURL), v))
	})
}

// DetailsURLHasPrefix applies the HasPrefix predicate on the "details_url" field.
func DetailsURLHasPrefix(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetailsURL), v))
	})
}

// DetailsURLHasSuffix applies the HasSuffix predicate on the "details_url" field.
func DetailsURLHasSuffix(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetailsURL), v))
	})
}

// DetailsURLEqualFold applies the EqualFold predicate on the "details_url" field.
func DetailsURLEqualFold(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetailsURL), v))
	})
}

// DetailsURLContainsFold applies the ContainsFold predicate on the "details_url" field.
func DetailsURLContainsFold(v string) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetailsURL), v))
	})
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatitude), v))
	})
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLatitude), v))
	})
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLatitude), v...))
	})
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLatitude), v...))
	})
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLatitude), v))
	})
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLatitude), v))
	})
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLatitude), v))
	})
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLatitude), v))
	})
}

// LatitudeIsNil applies the IsNil predicate on the "latitude" field.
func LatitudeIsNil() predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLatitude)))
	})
}

// LatitudeNotNil applies the NotNil predicate on the "latitude" field.
func LatitudeNotNil() predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLatitude)))
	})
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitude), v))
	})
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLongitude), v))
	})
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLongitude), v...))
	})
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLongitude), v...))
	})
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLongitude), v))
	})
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLongitude), v))
	})
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLongitude), v))
	})
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLongitude), v))
	})
}

// LongitudeIsNil applies the IsNil predicate on the "longitude" field.
func LongitudeIsNil() predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLongitude)))
	})
}

// LongitudeNotNil applies the NotNil predicate on the "longitude" field.
func LongitudeNotNil() predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLongitude)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedAt)))
	})
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedAt)))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Translator {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Translator) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Translator) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Translator) predicate.Translator {
	return predicate.Translator(func(s *sql.Selector) {
		p(s.Not())
	})
}
