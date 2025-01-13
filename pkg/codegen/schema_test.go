package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProperty_GoTypeDef(t *testing.T) {
	type fields struct {
		DisableRequiredReadOnlyAsPointer bool
		Schema                           Schema
		Required                         bool
		Nullable                         bool
		ReadOnly                         bool
		WriteOnly                        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			// When pointer is skipped by setting flag SkipOptionalPointer, the
			// flag will never be pointer irrespective of other flags.
			name: "Set skip optional pointer type for go type",
			fields: fields{
				Schema: Schema{
					SkipOptionalPointer: true,
					RefType:             "",
					GoType:              "int",
				},
			},
			want: "int",
		},

		{
			// if the field is optional, it will always be pointer irrespective of other
			// flags, given that pointer type is not skipped by setting SkipOptionalPointer
			// flag to true
			name: "When the field is optional",
			fields: fields{
				Schema: Schema{
					SkipOptionalPointer: false,
					RefType:             "",
					GoType:              "int",
				},
				Required: false,
			},
			want: "*int",
		},

		{
			// if the field(custom-type) is optional, it will NOT be a pointer if
			// SkipOptionalPointer flag is set to true
			name: "Set skip optional pointer type for ref type",
			fields: fields{
				Schema: Schema{
					SkipOptionalPointer: true,
					RefType:             "CustomType",
					GoType:              "int",
				},
				Required: false,
			},
			want: "CustomType",
		},

		// For the following test cases, SkipOptionalPointer flag is false.
		{
			name: "When field is required and not nullable",
			fields: fields{
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				Required: true,
				Nullable: false,
			},
			want: "int",
		},

		{
			name: "When field is required and nullable",
			fields: fields{
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				Required: true,
				Nullable: true,
			},
			want: "*int",
		},

		{
			name: "When field is optional and not nullable",
			fields: fields{
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				Required: false,
				Nullable: false,
			},
			want: "*int",
		},

		{
			name: "When field is optional and nullable",
			fields: fields{
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				Required: false,
				Nullable: true,
			},
			want: "*int",
		},

		// Following tests cases for non-nullable and required; and skip pointer is not opted
		{
			name: "When field is readOnly it will always be pointer",
			fields: fields{
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				ReadOnly: true,
				Required: true,
			},
			want: "*int",
		},

		{
			name: "When field is readOnly and read only pointer disabled",
			fields: fields{
				DisableRequiredReadOnlyAsPointer: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				ReadOnly: true,
				Required: true,
			},
			want: "int",
		},

		{
			name: "When field is readOnly and optional",
			fields: fields{
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				ReadOnly: true,
				Required: false,
			},
			want: "*int",
		},
		{
			name: "When field is readOnly and optional and read only pointer disabled",
			fields: fields{
				DisableRequiredReadOnlyAsPointer: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				ReadOnly: true,
				Required: false,
			},
			want: "*int",
		},

		// When field is write only, it will always be pointer unless pointer is
		// skipped by setting SkipOptionalPointer flag
		{
			name: "When field is write only and read only pointer disabled",
			fields: fields{
				DisableRequiredReadOnlyAsPointer: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				WriteOnly: true,
			},
			want: "*int",
		},

		{
			name: "When field is write only and read only pointer enabled",
			fields: fields{
				DisableRequiredReadOnlyAsPointer: false,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				WriteOnly: true,
			},
			want: "*int",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := &State{}
			gen.options.Compatibility.DisableRequiredReadOnlyAsPointer = tt.fields.DisableRequiredReadOnlyAsPointer
			p := Property{
				Schema:    tt.fields.Schema,
				Required:  tt.fields.Required,
				Nullable:  tt.fields.Nullable,
				ReadOnly:  tt.fields.ReadOnly,
				WriteOnly: tt.fields.WriteOnly,
				state:     gen,
			}
			assert.Equal(t, tt.want, p.GoTypeDef())
		})
	}
}

func TestProperty_GoTypeDef_nullable(t *testing.T) {
	type fields struct {
		DisableRequiredReadOnlyAsPointer bool
		NullableType                     bool
		Schema                           Schema
		Required                         bool
		Nullable                         bool
		ReadOnly                         bool
		WriteOnly                        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			// Field not nullable.
			// When pointer is skipped by setting flag SkipOptionalPointer, the
			// flag will never be pointer irrespective of other flags.
			name: "Set skip optional pointer type for go type",
			fields: fields{
				NullableType: true,
				Schema: Schema{
					SkipOptionalPointer: true,
					RefType:             "",
					GoType:              "int",
				},
			},
			want: "int",
		},

		{
			// Field not nullable.
			// if the field is optional, it will always be pointer irrespective of other
			// flags, given that pointer type is not skipped by setting SkipOptionalPointer
			// flag to true
			name: "When the field is optional",
			fields: fields{
				NullableType: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					RefType:             "",
					GoType:              "int",
				},
				Required: false,
			},
			want: "*int",
		},

		{
			// Field not nullable.
			// if the field(custom type) is optional, it will NOT be a pointer if
			// SkipOptionalPointer flag is set to true
			name: "Set skip optional pointer type for ref type",
			fields: fields{
				NullableType: true,
				Schema: Schema{
					SkipOptionalPointer: true,
					RefType:             "CustomType",
					GoType:              "int",
				},
				Required: false,
			},
			want: "CustomType",
		},

		// Field not nullable.
		// For the following test case, SkipOptionalPointer flag is false.
		{
			name: "When field is required and not nullable",
			fields: fields{
				NullableType: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				Required: true,
				Nullable: false,
			},
			want: "int",
		},

		{
			name: "When field is required and nullable",
			fields: fields{
				NullableType: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				Required: true,
				Nullable: true,
			},
			want: "nullable.Nullable[int]",
		},

		{
			name: "When field is optional and not nullable",
			fields: fields{
				NullableType: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				Required: false,
				Nullable: false,
			},
			want: "*int",
		},

		{
			name: "When field is optional and nullable",
			fields: fields{
				NullableType: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				Required: false,
				Nullable: true,
			},
			want: "nullable.Nullable[int]",
		},

		{
			name: "When field is readOnly, non-nullable and required and skip pointer is not opted",
			fields: fields{
				NullableType: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				ReadOnly: true,
				Required: true,
			},
			want: "*int",
		},

		{
			name: "When field is readOnly, required, non-nullable and read only pointer disabled",
			fields: fields{
				NullableType:                     true,
				DisableRequiredReadOnlyAsPointer: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				ReadOnly: true,
				Required: true,
			},
			want: "int",
		},

		{
			name: "When field is readOnly, optional and non nullable",
			fields: fields{
				NullableType: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				ReadOnly: true,
				Required: false,
			},
			want: "*int",
		},
		{
			name: "When field is readOnly and optional and read only pointer disabled",
			fields: fields{
				NullableType:                     true,
				DisableRequiredReadOnlyAsPointer: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				ReadOnly: true,
				Required: false,
			},
			want: "*int",
		},

		{
			name: "When field is write only and non nullable",
			fields: fields{
				NullableType:                     true,
				DisableRequiredReadOnlyAsPointer: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				WriteOnly: true,
			},
			want: "*int",
		},

		{
			name: "When field is write only and nullable",
			fields: fields{
				NullableType:                     true,
				DisableRequiredReadOnlyAsPointer: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				WriteOnly: true,
				Nullable:  true,
			},
			want: "nullable.Nullable[int]",
		},

		{
			name: "When field is write only, nullable and read only pointer enabled",
			fields: fields{
				NullableType: true,
				Schema: Schema{
					SkipOptionalPointer: false,
					GoType:              "int",
				},
				WriteOnly: true,
				Nullable:  true,
			},
			want: "nullable.Nullable[int]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := &State{}
			gen.options.Compatibility.DisableRequiredReadOnlyAsPointer = tt.fields.DisableRequiredReadOnlyAsPointer
			gen.options.OutputOptions.NullableType = tt.fields.NullableType
			p := Property{
				Schema:    tt.fields.Schema,
				Required:  tt.fields.Required,
				Nullable:  tt.fields.Nullable,
				ReadOnly:  tt.fields.ReadOnly,
				WriteOnly: tt.fields.WriteOnly,
				state:     gen,
			}
			assert.Equal(t, tt.want, p.GoTypeDef())
		})
	}
}
