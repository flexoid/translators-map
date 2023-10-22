// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/flexoid/translators-map-go/ent/predicate"
	"github.com/flexoid/translators-map-go/ent/translator"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTranslator = "Translator"
)

// TranslatorMutation represents an operation that mutates the Translator nodes in the graph.
type TranslatorMutation struct {
	config
	op             Op
	typ            string
	id             *int
	external_id    *int
	addexternal_id *int
	name           *string
	language       *string
	address        *string
	address_sha    *[]byte
	details_url    *string
	latitude       *float64
	addlatitude    *float64
	longitude      *float64
	addlongitude   *float64
	created_at     *time.Time
	updated_at     *time.Time
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Translator, error)
	predicates     []predicate.Translator
}

var _ ent.Mutation = (*TranslatorMutation)(nil)

// translatorOption allows management of the mutation configuration using functional options.
type translatorOption func(*TranslatorMutation)

// newTranslatorMutation creates new mutation for the Translator entity.
func newTranslatorMutation(c config, op Op, opts ...translatorOption) *TranslatorMutation {
	m := &TranslatorMutation{
		config:        c,
		op:            op,
		typ:           TypeTranslator,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTranslatorID sets the ID field of the mutation.
func withTranslatorID(id int) translatorOption {
	return func(m *TranslatorMutation) {
		var (
			err   error
			once  sync.Once
			value *Translator
		)
		m.oldValue = func(ctx context.Context) (*Translator, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Translator.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTranslator sets the old Translator of the mutation.
func withTranslator(node *Translator) translatorOption {
	return func(m *TranslatorMutation) {
		m.oldValue = func(context.Context) (*Translator, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TranslatorMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TranslatorMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TranslatorMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TranslatorMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Translator.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetExternalID sets the "external_id" field.
func (m *TranslatorMutation) SetExternalID(i int) {
	m.external_id = &i
	m.addexternal_id = nil
}

// ExternalID returns the value of the "external_id" field in the mutation.
func (m *TranslatorMutation) ExternalID() (r int, exists bool) {
	v := m.external_id
	if v == nil {
		return
	}
	return *v, true
}

// OldExternalID returns the old "external_id" field's value of the Translator entity.
// If the Translator object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranslatorMutation) OldExternalID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldExternalID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldExternalID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExternalID: %w", err)
	}
	return oldValue.ExternalID, nil
}

// AddExternalID adds i to the "external_id" field.
func (m *TranslatorMutation) AddExternalID(i int) {
	if m.addexternal_id != nil {
		*m.addexternal_id += i
	} else {
		m.addexternal_id = &i
	}
}

// AddedExternalID returns the value that was added to the "external_id" field in this mutation.
func (m *TranslatorMutation) AddedExternalID() (r int, exists bool) {
	v := m.addexternal_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetExternalID resets all changes to the "external_id" field.
func (m *TranslatorMutation) ResetExternalID() {
	m.external_id = nil
	m.addexternal_id = nil
}

// SetName sets the "name" field.
func (m *TranslatorMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TranslatorMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Translator entity.
// If the Translator object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranslatorMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ClearName clears the value of the "name" field.
func (m *TranslatorMutation) ClearName() {
	m.name = nil
	m.clearedFields[translator.FieldName] = struct{}{}
}

// NameCleared returns if the "name" field was cleared in this mutation.
func (m *TranslatorMutation) NameCleared() bool {
	_, ok := m.clearedFields[translator.FieldName]
	return ok
}

// ResetName resets all changes to the "name" field.
func (m *TranslatorMutation) ResetName() {
	m.name = nil
	delete(m.clearedFields, translator.FieldName)
}

// SetLanguage sets the "language" field.
func (m *TranslatorMutation) SetLanguage(s string) {
	m.language = &s
}

// Language returns the value of the "language" field in the mutation.
func (m *TranslatorMutation) Language() (r string, exists bool) {
	v := m.language
	if v == nil {
		return
	}
	return *v, true
}

// OldLanguage returns the old "language" field's value of the Translator entity.
// If the Translator object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranslatorMutation) OldLanguage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLanguage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLanguage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLanguage: %w", err)
	}
	return oldValue.Language, nil
}

// ResetLanguage resets all changes to the "language" field.
func (m *TranslatorMutation) ResetLanguage() {
	m.language = nil
}

// SetAddress sets the "address" field.
func (m *TranslatorMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the value of the "address" field in the mutation.
func (m *TranslatorMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// OldAddress returns the old "address" field's value of the Translator entity.
// If the Translator object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranslatorMutation) OldAddress(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAddress is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAddress requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddress: %w", err)
	}
	return oldValue.Address, nil
}

// ClearAddress clears the value of the "address" field.
func (m *TranslatorMutation) ClearAddress() {
	m.address = nil
	m.clearedFields[translator.FieldAddress] = struct{}{}
}

// AddressCleared returns if the "address" field was cleared in this mutation.
func (m *TranslatorMutation) AddressCleared() bool {
	_, ok := m.clearedFields[translator.FieldAddress]
	return ok
}

// ResetAddress resets all changes to the "address" field.
func (m *TranslatorMutation) ResetAddress() {
	m.address = nil
	delete(m.clearedFields, translator.FieldAddress)
}

// SetAddressSha sets the "address_sha" field.
func (m *TranslatorMutation) SetAddressSha(b []byte) {
	m.address_sha = &b
}

// AddressSha returns the value of the "address_sha" field in the mutation.
func (m *TranslatorMutation) AddressSha() (r []byte, exists bool) {
	v := m.address_sha
	if v == nil {
		return
	}
	return *v, true
}

// OldAddressSha returns the old "address_sha" field's value of the Translator entity.
// If the Translator object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranslatorMutation) OldAddressSha(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAddressSha is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAddressSha requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddressSha: %w", err)
	}
	return oldValue.AddressSha, nil
}

// ResetAddressSha resets all changes to the "address_sha" field.
func (m *TranslatorMutation) ResetAddressSha() {
	m.address_sha = nil
}

// SetDetailsURL sets the "details_url" field.
func (m *TranslatorMutation) SetDetailsURL(s string) {
	m.details_url = &s
}

// DetailsURL returns the value of the "details_url" field in the mutation.
func (m *TranslatorMutation) DetailsURL() (r string, exists bool) {
	v := m.details_url
	if v == nil {
		return
	}
	return *v, true
}

// OldDetailsURL returns the old "details_url" field's value of the Translator entity.
// If the Translator object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranslatorMutation) OldDetailsURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDetailsURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDetailsURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDetailsURL: %w", err)
	}
	return oldValue.DetailsURL, nil
}

// ResetDetailsURL resets all changes to the "details_url" field.
func (m *TranslatorMutation) ResetDetailsURL() {
	m.details_url = nil
}

// SetLatitude sets the "latitude" field.
func (m *TranslatorMutation) SetLatitude(f float64) {
	m.latitude = &f
	m.addlatitude = nil
}

// Latitude returns the value of the "latitude" field in the mutation.
func (m *TranslatorMutation) Latitude() (r float64, exists bool) {
	v := m.latitude
	if v == nil {
		return
	}
	return *v, true
}

// OldLatitude returns the old "latitude" field's value of the Translator entity.
// If the Translator object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranslatorMutation) OldLatitude(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLatitude is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLatitude requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLatitude: %w", err)
	}
	return oldValue.Latitude, nil
}

// AddLatitude adds f to the "latitude" field.
func (m *TranslatorMutation) AddLatitude(f float64) {
	if m.addlatitude != nil {
		*m.addlatitude += f
	} else {
		m.addlatitude = &f
	}
}

// AddedLatitude returns the value that was added to the "latitude" field in this mutation.
func (m *TranslatorMutation) AddedLatitude() (r float64, exists bool) {
	v := m.addlatitude
	if v == nil {
		return
	}
	return *v, true
}

// ClearLatitude clears the value of the "latitude" field.
func (m *TranslatorMutation) ClearLatitude() {
	m.latitude = nil
	m.addlatitude = nil
	m.clearedFields[translator.FieldLatitude] = struct{}{}
}

// LatitudeCleared returns if the "latitude" field was cleared in this mutation.
func (m *TranslatorMutation) LatitudeCleared() bool {
	_, ok := m.clearedFields[translator.FieldLatitude]
	return ok
}

// ResetLatitude resets all changes to the "latitude" field.
func (m *TranslatorMutation) ResetLatitude() {
	m.latitude = nil
	m.addlatitude = nil
	delete(m.clearedFields, translator.FieldLatitude)
}

// SetLongitude sets the "longitude" field.
func (m *TranslatorMutation) SetLongitude(f float64) {
	m.longitude = &f
	m.addlongitude = nil
}

// Longitude returns the value of the "longitude" field in the mutation.
func (m *TranslatorMutation) Longitude() (r float64, exists bool) {
	v := m.longitude
	if v == nil {
		return
	}
	return *v, true
}

// OldLongitude returns the old "longitude" field's value of the Translator entity.
// If the Translator object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranslatorMutation) OldLongitude(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLongitude is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLongitude requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLongitude: %w", err)
	}
	return oldValue.Longitude, nil
}

// AddLongitude adds f to the "longitude" field.
func (m *TranslatorMutation) AddLongitude(f float64) {
	if m.addlongitude != nil {
		*m.addlongitude += f
	} else {
		m.addlongitude = &f
	}
}

// AddedLongitude returns the value that was added to the "longitude" field in this mutation.
func (m *TranslatorMutation) AddedLongitude() (r float64, exists bool) {
	v := m.addlongitude
	if v == nil {
		return
	}
	return *v, true
}

// ClearLongitude clears the value of the "longitude" field.
func (m *TranslatorMutation) ClearLongitude() {
	m.longitude = nil
	m.addlongitude = nil
	m.clearedFields[translator.FieldLongitude] = struct{}{}
}

// LongitudeCleared returns if the "longitude" field was cleared in this mutation.
func (m *TranslatorMutation) LongitudeCleared() bool {
	_, ok := m.clearedFields[translator.FieldLongitude]
	return ok
}

// ResetLongitude resets all changes to the "longitude" field.
func (m *TranslatorMutation) ResetLongitude() {
	m.longitude = nil
	m.addlongitude = nil
	delete(m.clearedFields, translator.FieldLongitude)
}

// SetCreatedAt sets the "created_at" field.
func (m *TranslatorMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TranslatorMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Translator entity.
// If the Translator object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranslatorMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ClearCreatedAt clears the value of the "created_at" field.
func (m *TranslatorMutation) ClearCreatedAt() {
	m.created_at = nil
	m.clearedFields[translator.FieldCreatedAt] = struct{}{}
}

// CreatedAtCleared returns if the "created_at" field was cleared in this mutation.
func (m *TranslatorMutation) CreatedAtCleared() bool {
	_, ok := m.clearedFields[translator.FieldCreatedAt]
	return ok
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TranslatorMutation) ResetCreatedAt() {
	m.created_at = nil
	delete(m.clearedFields, translator.FieldCreatedAt)
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TranslatorMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TranslatorMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Translator entity.
// If the Translator object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranslatorMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (m *TranslatorMutation) ClearUpdatedAt() {
	m.updated_at = nil
	m.clearedFields[translator.FieldUpdatedAt] = struct{}{}
}

// UpdatedAtCleared returns if the "updated_at" field was cleared in this mutation.
func (m *TranslatorMutation) UpdatedAtCleared() bool {
	_, ok := m.clearedFields[translator.FieldUpdatedAt]
	return ok
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TranslatorMutation) ResetUpdatedAt() {
	m.updated_at = nil
	delete(m.clearedFields, translator.FieldUpdatedAt)
}

// Where appends a list predicates to the TranslatorMutation builder.
func (m *TranslatorMutation) Where(ps ...predicate.Translator) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TranslatorMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TranslatorMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Translator, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TranslatorMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TranslatorMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Translator).
func (m *TranslatorMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TranslatorMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m.external_id != nil {
		fields = append(fields, translator.FieldExternalID)
	}
	if m.name != nil {
		fields = append(fields, translator.FieldName)
	}
	if m.language != nil {
		fields = append(fields, translator.FieldLanguage)
	}
	if m.address != nil {
		fields = append(fields, translator.FieldAddress)
	}
	if m.address_sha != nil {
		fields = append(fields, translator.FieldAddressSha)
	}
	if m.details_url != nil {
		fields = append(fields, translator.FieldDetailsURL)
	}
	if m.latitude != nil {
		fields = append(fields, translator.FieldLatitude)
	}
	if m.longitude != nil {
		fields = append(fields, translator.FieldLongitude)
	}
	if m.created_at != nil {
		fields = append(fields, translator.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, translator.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TranslatorMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case translator.FieldExternalID:
		return m.ExternalID()
	case translator.FieldName:
		return m.Name()
	case translator.FieldLanguage:
		return m.Language()
	case translator.FieldAddress:
		return m.Address()
	case translator.FieldAddressSha:
		return m.AddressSha()
	case translator.FieldDetailsURL:
		return m.DetailsURL()
	case translator.FieldLatitude:
		return m.Latitude()
	case translator.FieldLongitude:
		return m.Longitude()
	case translator.FieldCreatedAt:
		return m.CreatedAt()
	case translator.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TranslatorMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case translator.FieldExternalID:
		return m.OldExternalID(ctx)
	case translator.FieldName:
		return m.OldName(ctx)
	case translator.FieldLanguage:
		return m.OldLanguage(ctx)
	case translator.FieldAddress:
		return m.OldAddress(ctx)
	case translator.FieldAddressSha:
		return m.OldAddressSha(ctx)
	case translator.FieldDetailsURL:
		return m.OldDetailsURL(ctx)
	case translator.FieldLatitude:
		return m.OldLatitude(ctx)
	case translator.FieldLongitude:
		return m.OldLongitude(ctx)
	case translator.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case translator.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Translator field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TranslatorMutation) SetField(name string, value ent.Value) error {
	switch name {
	case translator.FieldExternalID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExternalID(v)
		return nil
	case translator.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case translator.FieldLanguage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLanguage(v)
		return nil
	case translator.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	case translator.FieldAddressSha:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddressSha(v)
		return nil
	case translator.FieldDetailsURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDetailsURL(v)
		return nil
	case translator.FieldLatitude:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLatitude(v)
		return nil
	case translator.FieldLongitude:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLongitude(v)
		return nil
	case translator.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case translator.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Translator field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TranslatorMutation) AddedFields() []string {
	var fields []string
	if m.addexternal_id != nil {
		fields = append(fields, translator.FieldExternalID)
	}
	if m.addlatitude != nil {
		fields = append(fields, translator.FieldLatitude)
	}
	if m.addlongitude != nil {
		fields = append(fields, translator.FieldLongitude)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TranslatorMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case translator.FieldExternalID:
		return m.AddedExternalID()
	case translator.FieldLatitude:
		return m.AddedLatitude()
	case translator.FieldLongitude:
		return m.AddedLongitude()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TranslatorMutation) AddField(name string, value ent.Value) error {
	switch name {
	case translator.FieldExternalID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddExternalID(v)
		return nil
	case translator.FieldLatitude:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLatitude(v)
		return nil
	case translator.FieldLongitude:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLongitude(v)
		return nil
	}
	return fmt.Errorf("unknown Translator numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TranslatorMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(translator.FieldName) {
		fields = append(fields, translator.FieldName)
	}
	if m.FieldCleared(translator.FieldAddress) {
		fields = append(fields, translator.FieldAddress)
	}
	if m.FieldCleared(translator.FieldLatitude) {
		fields = append(fields, translator.FieldLatitude)
	}
	if m.FieldCleared(translator.FieldLongitude) {
		fields = append(fields, translator.FieldLongitude)
	}
	if m.FieldCleared(translator.FieldCreatedAt) {
		fields = append(fields, translator.FieldCreatedAt)
	}
	if m.FieldCleared(translator.FieldUpdatedAt) {
		fields = append(fields, translator.FieldUpdatedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TranslatorMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TranslatorMutation) ClearField(name string) error {
	switch name {
	case translator.FieldName:
		m.ClearName()
		return nil
	case translator.FieldAddress:
		m.ClearAddress()
		return nil
	case translator.FieldLatitude:
		m.ClearLatitude()
		return nil
	case translator.FieldLongitude:
		m.ClearLongitude()
		return nil
	case translator.FieldCreatedAt:
		m.ClearCreatedAt()
		return nil
	case translator.FieldUpdatedAt:
		m.ClearUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Translator nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TranslatorMutation) ResetField(name string) error {
	switch name {
	case translator.FieldExternalID:
		m.ResetExternalID()
		return nil
	case translator.FieldName:
		m.ResetName()
		return nil
	case translator.FieldLanguage:
		m.ResetLanguage()
		return nil
	case translator.FieldAddress:
		m.ResetAddress()
		return nil
	case translator.FieldAddressSha:
		m.ResetAddressSha()
		return nil
	case translator.FieldDetailsURL:
		m.ResetDetailsURL()
		return nil
	case translator.FieldLatitude:
		m.ResetLatitude()
		return nil
	case translator.FieldLongitude:
		m.ResetLongitude()
		return nil
	case translator.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case translator.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Translator field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TranslatorMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TranslatorMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TranslatorMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TranslatorMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TranslatorMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TranslatorMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TranslatorMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Translator unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TranslatorMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Translator edge %s", name)
}
