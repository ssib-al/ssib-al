// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"ssib.al/ssib-al-back/ent/link"
	"ssib.al/ssib-al-back/ent/user"
)

// Link is the model entity for the Link schema.
type Link struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Domain holds the value of the "domain" field.
	Domain string `json:"domain,omitempty"`
	// URI holds the value of the "uri" field.
	URI string `json:"uri,omitempty"`
	// TargetURL holds the value of the "target_url" field.
	TargetURL string `json:"target_url,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID uuid.UUID `json:"owner_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LinkQuery when eager-loading is set.
	Edges        LinkEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LinkEdges holds the relations/edges for other nodes in the graph.
type LinkEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LinkEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Link) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case link.FieldDomain, link.FieldURI, link.FieldTargetURL:
			values[i] = new(sql.NullString)
		case link.FieldID, link.FieldOwnerID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Link fields.
func (l *Link) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case link.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				l.ID = *value
			}
		case link.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				l.Domain = value.String
			}
		case link.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				l.URI = value.String
			}
		case link.FieldTargetURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field target_url", values[i])
			} else if value.Valid {
				l.TargetURL = value.String
			}
		case link.FieldOwnerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value != nil {
				l.OwnerID = *value
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Link.
// This includes values selected through modifiers, order, etc.
func (l *Link) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Link entity.
func (l *Link) QueryOwner() *UserQuery {
	return NewLinkClient(l.config).QueryOwner(l)
}

// Update returns a builder for updating this Link.
// Note that you need to call Link.Unwrap() before calling this method if this Link
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Link) Update() *LinkUpdateOne {
	return NewLinkClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Link entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Link) Unwrap() *Link {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Link is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Link) String() string {
	var builder strings.Builder
	builder.WriteString("Link(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("domain=")
	builder.WriteString(l.Domain)
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(l.URI)
	builder.WriteString(", ")
	builder.WriteString("target_url=")
	builder.WriteString(l.TargetURL)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", l.OwnerID))
	builder.WriteByte(')')
	return builder.String()
}

// Links is a parsable slice of Link.
type Links []*Link
