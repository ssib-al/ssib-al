// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"ssib.al/ssib-al-back/ent/link"
	"ssib.al/ssib-al-back/ent/predicate"
	"ssib.al/ssib-al-back/ent/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeLink = "Link"
	TypeUser = "User"
)

// LinkMutation represents an operation that mutates the Link nodes in the graph.
type LinkMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	domain        *string
	uri           *string
	target_url    *string
	clearedFields map[string]struct{}
	owner         *uuid.UUID
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*Link, error)
	predicates    []predicate.Link
}

var _ ent.Mutation = (*LinkMutation)(nil)

// linkOption allows management of the mutation configuration using functional options.
type linkOption func(*LinkMutation)

// newLinkMutation creates new mutation for the Link entity.
func newLinkMutation(c config, op Op, opts ...linkOption) *LinkMutation {
	m := &LinkMutation{
		config:        c,
		op:            op,
		typ:           TypeLink,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLinkID sets the ID field of the mutation.
func withLinkID(id uuid.UUID) linkOption {
	return func(m *LinkMutation) {
		var (
			err   error
			once  sync.Once
			value *Link
		)
		m.oldValue = func(ctx context.Context) (*Link, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Link.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLink sets the old Link of the mutation.
func withLink(node *Link) linkOption {
	return func(m *LinkMutation) {
		m.oldValue = func(context.Context) (*Link, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LinkMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LinkMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Link entities.
func (m *LinkMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *LinkMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *LinkMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Link.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetDomain sets the "domain" field.
func (m *LinkMutation) SetDomain(s string) {
	m.domain = &s
}

// Domain returns the value of the "domain" field in the mutation.
func (m *LinkMutation) Domain() (r string, exists bool) {
	v := m.domain
	if v == nil {
		return
	}
	return *v, true
}

// OldDomain returns the old "domain" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldDomain(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDomain is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDomain requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDomain: %w", err)
	}
	return oldValue.Domain, nil
}

// ResetDomain resets all changes to the "domain" field.
func (m *LinkMutation) ResetDomain() {
	m.domain = nil
}

// SetURI sets the "uri" field.
func (m *LinkMutation) SetURI(s string) {
	m.uri = &s
}

// URI returns the value of the "uri" field in the mutation.
func (m *LinkMutation) URI() (r string, exists bool) {
	v := m.uri
	if v == nil {
		return
	}
	return *v, true
}

// OldURI returns the old "uri" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldURI(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURI is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURI requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURI: %w", err)
	}
	return oldValue.URI, nil
}

// ResetURI resets all changes to the "uri" field.
func (m *LinkMutation) ResetURI() {
	m.uri = nil
}

// SetTargetURL sets the "target_url" field.
func (m *LinkMutation) SetTargetURL(s string) {
	m.target_url = &s
}

// TargetURL returns the value of the "target_url" field in the mutation.
func (m *LinkMutation) TargetURL() (r string, exists bool) {
	v := m.target_url
	if v == nil {
		return
	}
	return *v, true
}

// OldTargetURL returns the old "target_url" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldTargetURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTargetURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTargetURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTargetURL: %w", err)
	}
	return oldValue.TargetURL, nil
}

// ResetTargetURL resets all changes to the "target_url" field.
func (m *LinkMutation) ResetTargetURL() {
	m.target_url = nil
}

// SetOwnerID sets the "owner_id" field.
func (m *LinkMutation) SetOwnerID(u uuid.UUID) {
	m.owner = &u
}

// OwnerID returns the value of the "owner_id" field in the mutation.
func (m *LinkMutation) OwnerID() (r uuid.UUID, exists bool) {
	v := m.owner
	if v == nil {
		return
	}
	return *v, true
}

// OldOwnerID returns the old "owner_id" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldOwnerID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOwnerID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOwnerID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwnerID: %w", err)
	}
	return oldValue.OwnerID, nil
}

// ClearOwnerID clears the value of the "owner_id" field.
func (m *LinkMutation) ClearOwnerID() {
	m.owner = nil
	m.clearedFields[link.FieldOwnerID] = struct{}{}
}

// OwnerIDCleared returns if the "owner_id" field was cleared in this mutation.
func (m *LinkMutation) OwnerIDCleared() bool {
	_, ok := m.clearedFields[link.FieldOwnerID]
	return ok
}

// ResetOwnerID resets all changes to the "owner_id" field.
func (m *LinkMutation) ResetOwnerID() {
	m.owner = nil
	delete(m.clearedFields, link.FieldOwnerID)
}

// ClearOwner clears the "owner" edge to the User entity.
func (m *LinkMutation) ClearOwner() {
	m.clearedowner = true
	m.clearedFields[link.FieldOwnerID] = struct{}{}
}

// OwnerCleared reports if the "owner" edge to the User entity was cleared.
func (m *LinkMutation) OwnerCleared() bool {
	return m.OwnerIDCleared() || m.clearedowner
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *LinkMutation) OwnerIDs() (ids []uuid.UUID) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *LinkMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Where appends a list predicates to the LinkMutation builder.
func (m *LinkMutation) Where(ps ...predicate.Link) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the LinkMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *LinkMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Link, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *LinkMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *LinkMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Link).
func (m *LinkMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LinkMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.domain != nil {
		fields = append(fields, link.FieldDomain)
	}
	if m.uri != nil {
		fields = append(fields, link.FieldURI)
	}
	if m.target_url != nil {
		fields = append(fields, link.FieldTargetURL)
	}
	if m.owner != nil {
		fields = append(fields, link.FieldOwnerID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LinkMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case link.FieldDomain:
		return m.Domain()
	case link.FieldURI:
		return m.URI()
	case link.FieldTargetURL:
		return m.TargetURL()
	case link.FieldOwnerID:
		return m.OwnerID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LinkMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case link.FieldDomain:
		return m.OldDomain(ctx)
	case link.FieldURI:
		return m.OldURI(ctx)
	case link.FieldTargetURL:
		return m.OldTargetURL(ctx)
	case link.FieldOwnerID:
		return m.OldOwnerID(ctx)
	}
	return nil, fmt.Errorf("unknown Link field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LinkMutation) SetField(name string, value ent.Value) error {
	switch name {
	case link.FieldDomain:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDomain(v)
		return nil
	case link.FieldURI:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURI(v)
		return nil
	case link.FieldTargetURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTargetURL(v)
		return nil
	case link.FieldOwnerID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwnerID(v)
		return nil
	}
	return fmt.Errorf("unknown Link field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LinkMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LinkMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LinkMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Link numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LinkMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(link.FieldOwnerID) {
		fields = append(fields, link.FieldOwnerID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LinkMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LinkMutation) ClearField(name string) error {
	switch name {
	case link.FieldOwnerID:
		m.ClearOwnerID()
		return nil
	}
	return fmt.Errorf("unknown Link nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LinkMutation) ResetField(name string) error {
	switch name {
	case link.FieldDomain:
		m.ResetDomain()
		return nil
	case link.FieldURI:
		m.ResetURI()
		return nil
	case link.FieldTargetURL:
		m.ResetTargetURL()
		return nil
	case link.FieldOwnerID:
		m.ResetOwnerID()
		return nil
	}
	return fmt.Errorf("unknown Link field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LinkMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, link.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LinkMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case link.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LinkMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LinkMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LinkMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, link.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LinkMutation) EdgeCleared(name string) bool {
	switch name {
	case link.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LinkMutation) ClearEdge(name string) error {
	switch name {
	case link.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Link unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LinkMutation) ResetEdge(name string) error {
	switch name {
	case link.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Link edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	username      *string
	email         *string
	password_hash *string
	clearedFields map[string]struct{}
	links         map[uuid.UUID]struct{}
	removedlinks  map[uuid.UUID]struct{}
	clearedlinks  bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ClearEmail clears the value of the "email" field.
func (m *UserMutation) ClearEmail() {
	m.email = nil
	m.clearedFields[user.FieldEmail] = struct{}{}
}

// EmailCleared returns if the "email" field was cleared in this mutation.
func (m *UserMutation) EmailCleared() bool {
	_, ok := m.clearedFields[user.FieldEmail]
	return ok
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
	delete(m.clearedFields, user.FieldEmail)
}

// SetPasswordHash sets the "password_hash" field.
func (m *UserMutation) SetPasswordHash(s string) {
	m.password_hash = &s
}

// PasswordHash returns the value of the "password_hash" field in the mutation.
func (m *UserMutation) PasswordHash() (r string, exists bool) {
	v := m.password_hash
	if v == nil {
		return
	}
	return *v, true
}

// OldPasswordHash returns the old "password_hash" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPasswordHash(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPasswordHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPasswordHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPasswordHash: %w", err)
	}
	return oldValue.PasswordHash, nil
}

// ResetPasswordHash resets all changes to the "password_hash" field.
func (m *UserMutation) ResetPasswordHash() {
	m.password_hash = nil
}

// AddLinkIDs adds the "links" edge to the Link entity by ids.
func (m *UserMutation) AddLinkIDs(ids ...uuid.UUID) {
	if m.links == nil {
		m.links = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.links[ids[i]] = struct{}{}
	}
}

// ClearLinks clears the "links" edge to the Link entity.
func (m *UserMutation) ClearLinks() {
	m.clearedlinks = true
}

// LinksCleared reports if the "links" edge to the Link entity was cleared.
func (m *UserMutation) LinksCleared() bool {
	return m.clearedlinks
}

// RemoveLinkIDs removes the "links" edge to the Link entity by IDs.
func (m *UserMutation) RemoveLinkIDs(ids ...uuid.UUID) {
	if m.removedlinks == nil {
		m.removedlinks = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.links, ids[i])
		m.removedlinks[ids[i]] = struct{}{}
	}
}

// RemovedLinks returns the removed IDs of the "links" edge to the Link entity.
func (m *UserMutation) RemovedLinksIDs() (ids []uuid.UUID) {
	for id := range m.removedlinks {
		ids = append(ids, id)
	}
	return
}

// LinksIDs returns the "links" edge IDs in the mutation.
func (m *UserMutation) LinksIDs() (ids []uuid.UUID) {
	for id := range m.links {
		ids = append(ids, id)
	}
	return
}

// ResetLinks resets all changes to the "links" edge.
func (m *UserMutation) ResetLinks() {
	m.links = nil
	m.clearedlinks = false
	m.removedlinks = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password_hash != nil {
		fields = append(fields, user.FieldPasswordHash)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUsername:
		return m.Username()
	case user.FieldEmail:
		return m.Email()
	case user.FieldPasswordHash:
		return m.PasswordHash()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPasswordHash:
		return m.OldPasswordHash(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPasswordHash:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPasswordHash(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldEmail) {
		fields = append(fields, user.FieldEmail)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldEmail:
		m.ClearEmail()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPasswordHash:
		m.ResetPasswordHash()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.links != nil {
		edges = append(edges, user.EdgeLinks)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeLinks:
		ids := make([]ent.Value, 0, len(m.links))
		for id := range m.links {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedlinks != nil {
		edges = append(edges, user.EdgeLinks)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeLinks:
		ids := make([]ent.Value, 0, len(m.removedlinks))
		for id := range m.removedlinks {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedlinks {
		edges = append(edges, user.EdgeLinks)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeLinks:
		return m.clearedlinks
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeLinks:
		m.ResetLinks()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
