/*
 This file was autogenerated via
 ---------------------------------------------------------------------------------------
 gen-builtin --declarable --type Hex64 --name hex64 --handler AddUint64 --go-name uint64
 ---------------------------------------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Hex64("")

// Hex64 represents field of type hex64
type Hex64 string

// Name returns field name
func (i Hex64) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Hex64) TypeName() string {
	return "hex64"
}

// Register registers a field
func (i Hex64) Register(comment []string, registrator FieldRegistrator) {
	registrator.AddUint64(comment, i.Name())
}

// GoName returns Go's representation of this field's type
func (i Hex64) GoName() string {
	return "uint64"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["hex64"] = func(fieldName string) Field {
		return Hex64(fieldName)
	}

	if declarables == nil {
		declarables = map[string]struct{}{}
	}
	declarables["hex64"] = struct{}{}

	if backedBy == nil {
		backedBy = map[string]string{}
	}
	backedBy["hex64"] = "uint64"
}
