/*
 This file was autogenerated via
 -------------------------------------------------------------------------------
 gen-builtin --declarable --type Oct --name oct --handler AddUint --go-name uint
 -------------------------------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Oct("")

// Oct represents field of type oct
type Oct string

// Name returns field name
func (i Oct) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Oct) TypeName() string {
	return "oct"
}

// Register registers a field
func (i Oct) Register(comment []string, registrator FieldRegistrator) {
	registrator.AddUint(comment, i.Name())
}

// GoName returns Go's representation of this field's type
func (i Oct) GoName() string {
	return "uint"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["oct"] = func(fieldName string) Field {
		return Oct(fieldName)
	}

	if declarables == nil {
		declarables = map[string]struct{}{}
	}
	declarables["oct"] = struct{}{}

	if backedBy == nil {
		backedBy = map[string]string{}
	}
	backedBy["oct"] = "uint"
}
