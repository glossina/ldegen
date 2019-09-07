/*
 This file was autogenerated via
 -----------------------------------------------------------------------------------
 gen-builtin --declarable --type Hex8 --name hex8 --handler AddUint8 --go-name uint8
 -----------------------------------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Hex8("")

// Hex8 represents field of type hex8
type Hex8 string

// Name returns field name
func (i Hex8) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Hex8) TypeName() string {
	return "hex8"
}

// Register registers a field
func (i Hex8) Register(comment []string, registrator FieldRegistrator) {
	registrator.AddUint8(comment, i.Name())
}

// GoName returns Go's representation of this field's type
func (i Hex8) GoName() string {
	return "uint8"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["hex8"] = func(fieldName string) Field {
		return Hex8(fieldName)
	}

	if declarables == nil {
		declarables = map[string]struct{}{}
	}
	declarables["hex8"] = struct{}{}

	if backedBy == nil {
		backedBy = map[string]string{}
	}
	backedBy["hex8"] = "uint8"
}
