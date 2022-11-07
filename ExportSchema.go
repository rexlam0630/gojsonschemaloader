package gojsonschemaloader

import "math/big"

type ExportSchema struct {
	// Root Properties
	Draft       *string `json:"$schema,omitempty"`
	Id          *string `json:"id,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`

	// Common Properties
	Types                []string                 `json:"type,omitempty"`
	Properties           map[string]*ExportSchema `json:"properties,omitempty"`
	Items                map[string]*ExportSchema `json:"items,omitempty"`
	MultipleOf           *big.Rat                 `json:"multipleOf,omitempty"`
	Maximum              *big.Rat                 `json:"maximum,omitempty"`
	ExclusiveMaximum     *big.Rat                 `json:"exclusiveMaximum,omitempty"`
	Minimum              *big.Rat                 `json:"minimum,omitempty"`
	ExclusiveMinimum     *big.Rat                 `json:"exclusiveMinimum,omitempty"`
	MinLength            *int                     `json:"minLength,omitempty"`
	MaxLength            *int                     `json:"maxLength,omitempty"`
	Pattern              *string                  `json:"pattern,omitempty"`
	Format               *string                  `json:"format,omitempty"`
	MinProperties        *int                     `json:"minProperties,omitempty"`
	MaxProperties        *int                     `json:"maxProperties,omitempty"`
	Required             []string                 `json:"required,omitempty"`
	Dependencies         map[string]interface{}   `json:"dependencies,omitempty"`
	AdditionalProperties interface{}              `json:"additionalProperties,omitempty"`
	PatternProperties    map[string]*ExportSchema `json:"patternProperties,omitempty"`
	PropertyNames        *ExportSchema            `json:"propertyNames,omitempty"`
	MinItems             *int                     `json:"minItems,omitempty"`
	MaxItems             *int                     `json:"maxItems,omitempty"`
	UniqueItems          *bool                    `json:"uniqueItems,omitempty"`
	Contains             *ExportSchema            `json:"contains,omitempty"`
	AdditionalItems      interface{}              `json:"additionalItems,omitempty"`
	Const                *string                  `json:"const,omitempty"`
	Enum                 []string                 `json:"enum,omitempty"`
	OneOf                []*ExportSchema          `json:"oneOf,omitempty"`
	AnyOf                []*ExportSchema          `json:"anyOf,omitempty"`
	AllOf                []*ExportSchema          `json:"allOf,omitempty"`
	Not                  *ExportSchema            `json:"not,omitempty"`
	If                   *ExportSchema            `json:"if,omitempty"`
	Then                 *ExportSchema            `json:"then,omitempty"`
	Else                 *ExportSchema            `json:"else,omitempty"`
}
