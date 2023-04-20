// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/fukunokaze/mfs_backend/MFSRepository/ent/mfsuser"
)

// MFSUser is the model entity for the MFSUser schema.
type MFSUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserId holds the value of the "userId" field.
	UserId int `json:"userId,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Address holds the value of the "address" field.
	Address      string `json:"address,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MFSUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mfsuser.FieldID, mfsuser.FieldUserId:
			values[i] = new(sql.NullInt64)
		case mfsuser.FieldUsername, mfsuser.FieldPassword, mfsuser.FieldAddress:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MFSUser fields.
func (mu *MFSUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mfsuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mu.ID = int(value.Int64)
		case mfsuser.FieldUserId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				mu.UserId = int(value.Int64)
			}
		case mfsuser.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				mu.Username = value.String
			}
		case mfsuser.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				mu.Password = value.String
			}
		case mfsuser.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				mu.Address = value.String
			}
		default:
			mu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MFSUser.
// This includes values selected through modifiers, order, etc.
func (mu *MFSUser) Value(name string) (ent.Value, error) {
	return mu.selectValues.Get(name)
}

// Update returns a builder for updating this MFSUser.
// Note that you need to call MFSUser.Unwrap() before calling this method if this MFSUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (mu *MFSUser) Update() *MFSUserUpdateOne {
	return NewMFSUserClient(mu.config).UpdateOne(mu)
}

// Unwrap unwraps the MFSUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mu *MFSUser) Unwrap() *MFSUser {
	_tx, ok := mu.config.driver.(*txDriver)
	if !ok {
		panic("ent: MFSUser is not a transactional entity")
	}
	mu.config.driver = _tx.drv
	return mu
}

// String implements the fmt.Stringer.
func (mu *MFSUser) String() string {
	var builder strings.Builder
	builder.WriteString("MFSUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mu.ID))
	builder.WriteString("userId=")
	builder.WriteString(fmt.Sprintf("%v", mu.UserId))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(mu.Username)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(mu.Password)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(mu.Address)
	builder.WriteByte(')')
	return builder.String()
}

// MFSUsers is a parsable slice of MFSUser.
type MFSUsers []*MFSUser
