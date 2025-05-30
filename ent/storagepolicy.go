// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cloudreve/Cloudreve/v4/ent/node"
	"github.com/cloudreve/Cloudreve/v4/ent/storagepolicy"
	"github.com/cloudreve/Cloudreve/v4/inventory/types"
)

// StoragePolicy is the model entity for the StoragePolicy schema.
type StoragePolicy struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Server holds the value of the "server" field.
	Server string `json:"server,omitempty"`
	// BucketName holds the value of the "bucket_name" field.
	BucketName string `json:"bucket_name,omitempty"`
	// IsPrivate holds the value of the "is_private" field.
	IsPrivate bool `json:"is_private,omitempty"`
	// AccessKey holds the value of the "access_key" field.
	AccessKey string `json:"access_key,omitempty"`
	// SecretKey holds the value of the "secret_key" field.
	SecretKey string `json:"secret_key,omitempty"`
	// MaxSize holds the value of the "max_size" field.
	MaxSize int64 `json:"max_size,omitempty"`
	// DirNameRule holds the value of the "dir_name_rule" field.
	DirNameRule string `json:"dir_name_rule,omitempty"`
	// FileNameRule holds the value of the "file_name_rule" field.
	FileNameRule string `json:"file_name_rule,omitempty"`
	// Settings holds the value of the "settings" field.
	Settings *types.PolicySetting `json:"settings,omitempty"`
	// NodeID holds the value of the "node_id" field.
	NodeID int `json:"node_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StoragePolicyQuery when eager-loading is set.
	Edges        StoragePolicyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StoragePolicyEdges holds the relations/edges for other nodes in the graph.
type StoragePolicyEdges struct {
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// Entities holds the value of the entities edge.
	Entities []*Entity `json:"entities,omitempty"`
	// Node holds the value of the node edge.
	Node *Node `json:"node,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e StoragePolicyEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[0] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e StoragePolicyEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[1] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// EntitiesOrErr returns the Entities value or an error if the edge
// was not loaded in eager-loading.
func (e StoragePolicyEdges) EntitiesOrErr() ([]*Entity, error) {
	if e.loadedTypes[2] {
		return e.Entities, nil
	}
	return nil, &NotLoadedError{edge: "entities"}
}

// NodeOrErr returns the Node value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StoragePolicyEdges) NodeOrErr() (*Node, error) {
	if e.loadedTypes[3] {
		if e.Node == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: node.Label}
		}
		return e.Node, nil
	}
	return nil, &NotLoadedError{edge: "node"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StoragePolicy) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case storagepolicy.FieldSettings:
			values[i] = new([]byte)
		case storagepolicy.FieldIsPrivate:
			values[i] = new(sql.NullBool)
		case storagepolicy.FieldID, storagepolicy.FieldMaxSize, storagepolicy.FieldNodeID:
			values[i] = new(sql.NullInt64)
		case storagepolicy.FieldName, storagepolicy.FieldType, storagepolicy.FieldServer, storagepolicy.FieldBucketName, storagepolicy.FieldAccessKey, storagepolicy.FieldSecretKey, storagepolicy.FieldDirNameRule, storagepolicy.FieldFileNameRule:
			values[i] = new(sql.NullString)
		case storagepolicy.FieldCreatedAt, storagepolicy.FieldUpdatedAt, storagepolicy.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StoragePolicy fields.
func (sp *StoragePolicy) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case storagepolicy.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sp.ID = int(value.Int64)
		case storagepolicy.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sp.CreatedAt = value.Time
			}
		case storagepolicy.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sp.UpdatedAt = value.Time
			}
		case storagepolicy.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sp.DeletedAt = new(time.Time)
				*sp.DeletedAt = value.Time
			}
		case storagepolicy.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sp.Name = value.String
			}
		case storagepolicy.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				sp.Type = value.String
			}
		case storagepolicy.FieldServer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field server", values[i])
			} else if value.Valid {
				sp.Server = value.String
			}
		case storagepolicy.FieldBucketName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bucket_name", values[i])
			} else if value.Valid {
				sp.BucketName = value.String
			}
		case storagepolicy.FieldIsPrivate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_private", values[i])
			} else if value.Valid {
				sp.IsPrivate = value.Bool
			}
		case storagepolicy.FieldAccessKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_key", values[i])
			} else if value.Valid {
				sp.AccessKey = value.String
			}
		case storagepolicy.FieldSecretKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_key", values[i])
			} else if value.Valid {
				sp.SecretKey = value.String
			}
		case storagepolicy.FieldMaxSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_size", values[i])
			} else if value.Valid {
				sp.MaxSize = value.Int64
			}
		case storagepolicy.FieldDirNameRule:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dir_name_rule", values[i])
			} else if value.Valid {
				sp.DirNameRule = value.String
			}
		case storagepolicy.FieldFileNameRule:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_name_rule", values[i])
			} else if value.Valid {
				sp.FileNameRule = value.String
			}
		case storagepolicy.FieldSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &sp.Settings); err != nil {
					return fmt.Errorf("unmarshal field settings: %w", err)
				}
			}
		case storagepolicy.FieldNodeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field node_id", values[i])
			} else if value.Valid {
				sp.NodeID = int(value.Int64)
			}
		default:
			sp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the StoragePolicy.
// This includes values selected through modifiers, order, etc.
func (sp *StoragePolicy) Value(name string) (ent.Value, error) {
	return sp.selectValues.Get(name)
}

// QueryGroups queries the "groups" edge of the StoragePolicy entity.
func (sp *StoragePolicy) QueryGroups() *GroupQuery {
	return NewStoragePolicyClient(sp.config).QueryGroups(sp)
}

// QueryFiles queries the "files" edge of the StoragePolicy entity.
func (sp *StoragePolicy) QueryFiles() *FileQuery {
	return NewStoragePolicyClient(sp.config).QueryFiles(sp)
}

// QueryEntities queries the "entities" edge of the StoragePolicy entity.
func (sp *StoragePolicy) QueryEntities() *EntityQuery {
	return NewStoragePolicyClient(sp.config).QueryEntities(sp)
}

// QueryNode queries the "node" edge of the StoragePolicy entity.
func (sp *StoragePolicy) QueryNode() *NodeQuery {
	return NewStoragePolicyClient(sp.config).QueryNode(sp)
}

// Update returns a builder for updating this StoragePolicy.
// Note that you need to call StoragePolicy.Unwrap() before calling this method if this StoragePolicy
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *StoragePolicy) Update() *StoragePolicyUpdateOne {
	return NewStoragePolicyClient(sp.config).UpdateOne(sp)
}

// Unwrap unwraps the StoragePolicy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *StoragePolicy) Unwrap() *StoragePolicy {
	_tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: StoragePolicy is not a transactional entity")
	}
	sp.config.driver = _tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *StoragePolicy) String() string {
	var builder strings.Builder
	builder.WriteString("StoragePolicy(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := sp.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sp.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(sp.Type)
	builder.WriteString(", ")
	builder.WriteString("server=")
	builder.WriteString(sp.Server)
	builder.WriteString(", ")
	builder.WriteString("bucket_name=")
	builder.WriteString(sp.BucketName)
	builder.WriteString(", ")
	builder.WriteString("is_private=")
	builder.WriteString(fmt.Sprintf("%v", sp.IsPrivate))
	builder.WriteString(", ")
	builder.WriteString("access_key=")
	builder.WriteString(sp.AccessKey)
	builder.WriteString(", ")
	builder.WriteString("secret_key=")
	builder.WriteString(sp.SecretKey)
	builder.WriteString(", ")
	builder.WriteString("max_size=")
	builder.WriteString(fmt.Sprintf("%v", sp.MaxSize))
	builder.WriteString(", ")
	builder.WriteString("dir_name_rule=")
	builder.WriteString(sp.DirNameRule)
	builder.WriteString(", ")
	builder.WriteString("file_name_rule=")
	builder.WriteString(sp.FileNameRule)
	builder.WriteString(", ")
	builder.WriteString("settings=")
	builder.WriteString(fmt.Sprintf("%v", sp.Settings))
	builder.WriteString(", ")
	builder.WriteString("node_id=")
	builder.WriteString(fmt.Sprintf("%v", sp.NodeID))
	builder.WriteByte(')')
	return builder.String()
}

// SetGroups manually set the edge as loaded state.
func (e *StoragePolicy) SetGroups(v []*Group) {
	e.Edges.Groups = v
	e.Edges.loadedTypes[0] = true
}

// SetFiles manually set the edge as loaded state.
func (e *StoragePolicy) SetFiles(v []*File) {
	e.Edges.Files = v
	e.Edges.loadedTypes[1] = true
}

// SetEntities manually set the edge as loaded state.
func (e *StoragePolicy) SetEntities(v []*Entity) {
	e.Edges.Entities = v
	e.Edges.loadedTypes[2] = true
}

// SetNode manually set the edge as loaded state.
func (e *StoragePolicy) SetNode(v *Node) {
	e.Edges.Node = v
	e.Edges.loadedTypes[3] = true
}

// StoragePolicies is a parsable slice of StoragePolicy.
type StoragePolicies []*StoragePolicy
