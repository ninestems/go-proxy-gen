package scanner

import (
	"go/ast"
	"go/token"

	"github.com/ninestems/go-proxy-gen/entity"
)

// declaration extracts interfaces from ast tree.
func declaration(in ast.Decl) []*entity.Interface {
	genDecl, ok := in.(*ast.GenDecl)
	if !ok || genDecl.Tok != token.TYPE {
		return nil
	}

	return ifaces(genDecl.Specs)
}

// declarations extracts interfaces from ast tree.
func declarations(in []ast.Decl) []*entity.Interface {
	var out = make([]*entity.Interface, 0, len(in))
	for _, spec := range in {
		if ifas := declaration(spec); len(ifas) > 0 {
			out = append(out, ifas...)
		}
	}

	return out
}
