package scanner

import (
	"go/ast"

	"github.com/ninestems/go-proxy-gen/entity"
)

// iface extracts interface from ast tree.
func iface(in *ast.TypeSpec) *entity.Interface {
	ifa, ok := in.Type.(*ast.InterfaceType)
	if !ok {
		return nil
	}

	return entity.NewInterface(in.Name.Name, functions(ifa.Methods.List))
}

// ifaces extracts interfaces from ast tree.
func ifaces(in []ast.Spec) []*entity.Interface {
	var out = make([]*entity.Interface, 0, len(in))

	for _, spec := range in {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		if ifa := iface(typeSpec); ifa != nil {
			out = append(out, ifa)
		}
	}

	return out
}
