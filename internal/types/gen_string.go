/*
 This file was autogenerated via
 -------------------------------------------------------------
 gen-builtin --declarable --native --type String --name string
 -------------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = String("")

// String represents field of type string
type String string

// Name returns field name
func (i String) Name() string {
	return string(i)
}

// TypeName name of the type
func (i String) TypeName() string {
	return "string"
}

// Register registers a field
func (i String) Register(comment []string, registrator FieldRegistrator) {
	registrator.AddString(comment, i.Name())
}

// GoName returns Go's representation of this field's type
func (i String) GoName() string {
	return "string"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["string"] = func(fieldName string) Field {
		return String(fieldName)
	}
	if natives == nil {
		natives = map[string]struct{}{}
	}
	natives["string"] = struct{}{}
	if declarables == nil {
		declarables = map[string]struct{}{}
	}
	declarables["string"] = struct{}{}

}
