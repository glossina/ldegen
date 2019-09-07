/*
 This file was autogenerated via
 ----------------------------------------------------------------------------------
 gen-builtin --decimal --type Dec32 --name dec32 --handler AddInt32 --go-name int32
 ----------------------------------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Dec32("")

// Dec32 represents field of type dec32
type Dec32 string

// Name returns field name
func (i Dec32) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Dec32) TypeName() string {
	return "dec32"
}

// Register registers a field
func (i Dec32) Register(comment []string, registrator FieldRegistrator) {
	registrator.AddInt32(comment, i.Name())
}

// GoName returns Go's representation of this field's type
func (i Dec32) GoName() string {
	return "int32"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["dec32"] = func(fieldName string) Field {
		return Dec32(fieldName)
	}

	if decimals == nil {
		decimals = map[string]struct{}{}
	}
	decimals["dec32"] = struct{}{}
	if backedBy == nil {
		backedBy = map[string]string{}
	}
	backedBy["dec32"] = "int32"
}
