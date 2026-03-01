package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ninestems/go-proxy-gen/pkg/entity"
)

func Test_importSpecification_Build(t *testing.T) {
	type fields struct {
		alias string
		path  string
	}
	tests := []struct {
		name   string
		fields fields
		want   *entity.Import
	}{
		{
			name: "build import",
			fields: fields{
				alias: "alias",
				path:  "int",
			},
			want: entity.NewImport("alias", "int"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &importSpecification{
				alias: tt.fields.alias,
				path:  tt.fields.path,
			}
			require.Equal(t, tt.want, s.Build())
		})
	}
}
