package validators

import (
	"strings"

	structpb "github.com/golang/protobuf/ptypes/struct"

	nebula "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
)

type typeChecker interface {
	CastsFrom(*nebula.LiteralType) bool
}

type trivialChecker struct {
	literalType *nebula.LiteralType
}

// CastsFrom is a trivial type checker merely checks if types match exactly.
func (t trivialChecker) CastsFrom(upstreamType *nebula.LiteralType) bool {
	// If upstream is an enum, it can be consumed as a string downstream
	if upstreamType.GetEnumType() != nil {
		if t.literalType.GetSimple() == nebula.SimpleType_STRING {
			return true
		}
	}
	// If t is an enum, it can be created from a string as Enums as just constrained String aliases
	if t.literalType.GetEnumType() != nil {
		if upstreamType.GetSimple() == nebula.SimpleType_STRING {
			return true
		}
	}

	if GetTagForType(upstreamType) != "" && GetTagForType(t.literalType) != GetTagForType(upstreamType) {
		return false
	}

	// Ignore metadata when comparing types.
	upstreamTypeCopy := *upstreamType
	downstreamTypeCopy := *t.literalType
	upstreamTypeCopy.Structure = &nebula.TypeStructure{}
	downstreamTypeCopy.Structure = &nebula.TypeStructure{}
	upstreamTypeCopy.Metadata = &structpb.Struct{}
	downstreamTypeCopy.Metadata = &structpb.Struct{}
	upstreamTypeCopy.Annotation = &nebula.TypeAnnotation{}
	downstreamTypeCopy.Annotation = &nebula.TypeAnnotation{}
	return upstreamTypeCopy.String() == downstreamTypeCopy.String()
}

type noneTypeChecker struct{}

// CastsFrom matches only void
func (t noneTypeChecker) CastsFrom(upstreamType *nebula.LiteralType) bool {
	return isNoneType(upstreamType)
}

type mapTypeChecker struct {
	literalType *nebula.LiteralType
}

// CastsFrom checks that the target map type can be cast to the current map type. We need to ensure both the key types
// and value types match.
func (t mapTypeChecker) CastsFrom(upstreamType *nebula.LiteralType) bool {
	// Empty maps should match any collection.
	mapLiteralType := upstreamType.GetMapValueType()
	if isNoneType(mapLiteralType) {
		return true
	} else if mapLiteralType != nil {
		return getTypeChecker(t.literalType.GetMapValueType()).CastsFrom(mapLiteralType)
	}

	return false
}

type collectionTypeChecker struct {
	literalType *nebula.LiteralType
}

// CastsFrom checks whether two collection types match. We need to ensure that the nesting is correct and the final
// subtypes match.
func (t collectionTypeChecker) CastsFrom(upstreamType *nebula.LiteralType) bool {
	// Empty collections should match any collection.
	collectionType := upstreamType.GetCollectionType()
	if isNoneType(upstreamType.GetCollectionType()) {
		return true
	} else if collectionType != nil {
		return getTypeChecker(t.literalType.GetCollectionType()).CastsFrom(collectionType)
	}

	return false
}

type schemaTypeChecker struct {
	literalType *nebula.LiteralType
}

// CastsFrom handles type casting to the underlying schema type.
// Schemas are more complex types in the Nebula ecosystem. A schema is considered castable in the following
// cases.
//
//  1. The downstream schema has no column types specified.  In such a case, it accepts all schema input since it is
//     generic.
//
//  2. The downstream schema has a subset of the upstream columns and they match perfectly.
//
//  3. The upstream type can be Schema type or structured dataset type
func (t schemaTypeChecker) CastsFrom(upstreamType *nebula.LiteralType) bool {
	schemaType := upstreamType.GetSchema()
	structuredDatasetType := upstreamType.GetStructuredDatasetType()
	if structuredDatasetType == nil && schemaType == nil {
		return false
	}

	if schemaType != nil {
		return schemaCastFromSchema(schemaType, t.literalType.GetSchema())
	}

	// Nebula Schema can only be serialized to parquet
	if len(structuredDatasetType.Format) != 0 && !strings.EqualFold(structuredDatasetType.Format, "parquet") {
		return false
	}

	return schemaCastFromStructuredDataset(structuredDatasetType, t.literalType.GetSchema())
}

type structuredDatasetChecker struct {
	literalType *nebula.LiteralType
}

// CastsFrom for Structured dataset are more complex types in the Nebula ecosystem. A structured dataset is considered
// castable in the following cases:
//
//  1. The downstream structured dataset has no column types specified.  In such a case, it accepts all structured dataset input since it is
//     generic.
//
//  2. The downstream structured dataset has a subset of the upstream structured dataset columns and they match perfectly.
//
//  3. The upstream type can be Schema type or structured dataset type
func (t structuredDatasetChecker) CastsFrom(upstreamType *nebula.LiteralType) bool {
	// structured datasets are nullable
	if isNoneType(upstreamType) {
		return true
	}
	structuredDatasetType := upstreamType.GetStructuredDatasetType()
	schemaType := upstreamType.GetSchema()
	if structuredDatasetType == nil && schemaType == nil {
		return false
	}
	if schemaType != nil {
		// Nebula Schema can only be serialized to parquet
		format := t.literalType.GetStructuredDatasetType().Format
		if len(format) != 0 && !strings.EqualFold(format, "parquet") {
			return false
		}
		return structuredDatasetCastFromSchema(schemaType, t.literalType.GetStructuredDatasetType())
	}
	return structuredDatasetCastFromStructuredDataset(structuredDatasetType, t.literalType.GetStructuredDatasetType())
}

// Upstream (schema) -> downstream (schema)
func schemaCastFromSchema(upstream *nebula.SchemaType, downstream *nebula.SchemaType) bool {
	if len(upstream.Columns) == 0 || len(downstream.Columns) == 0 {
		return true
	}

	nameToTypeMap := make(map[string]nebula.SchemaType_SchemaColumn_SchemaColumnType)
	for _, column := range upstream.Columns {
		nameToTypeMap[column.Name] = column.Type
	}

	// Check that the downstream schema is a strict sub-set of the upstream schema.
	for _, column := range downstream.Columns {
		upstreamType, ok := nameToTypeMap[column.Name]
		if !ok {
			return false
		}
		if upstreamType != column.Type {
			return false
		}
	}
	return true
}

type unionTypeChecker struct {
	literalType *nebula.LiteralType
}

func (t unionTypeChecker) CastsFrom(upstreamType *nebula.LiteralType) bool {
	unionType := t.literalType.GetUnionType()

	upstreamUnionType := upstreamType.GetUnionType()
	if upstreamUnionType != nil {
		// For each upstream variant we must find a compatible downstream variant
		for _, u := range upstreamUnionType.GetVariants() {
			found := false
			for _, d := range unionType.GetVariants() {
				if AreTypesCastable(u, d) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}

		return true
	}

	// Matches iff we can unambiguously select a variant
	foundOne := false
	for _, x := range unionType.GetVariants() {
		if AreTypesCastable(upstreamType, x) {
			if foundOne {
				return false
			}
			foundOne = true
		}
	}

	return foundOne
}

// Upstream (structuredDatasetType) -> downstream (structuredDatasetType)
func structuredDatasetCastFromStructuredDataset(upstream *nebula.StructuredDatasetType, downstream *nebula.StructuredDatasetType) bool {
	// Skip the format check here when format is empty. https://github.com/nebulaclouds/nebula/issues/2864
	if len(upstream.Format) != 0 && len(downstream.Format) != 0 && !strings.EqualFold(upstream.Format, downstream.Format) {
		return false
	}

	if len(upstream.Columns) == 0 || len(downstream.Columns) == 0 {
		return true
	}

	nameToTypeMap := make(map[string]*nebula.LiteralType)
	for _, column := range upstream.Columns {
		nameToTypeMap[column.Name] = column.LiteralType
	}

	// Check that the downstream structured dataset is a strict sub-set of the upstream structured dataset.
	for _, column := range downstream.Columns {
		upstreamType, ok := nameToTypeMap[column.Name]
		if !ok {
			return false
		}
		if !getTypeChecker(column.LiteralType).CastsFrom(upstreamType) {
			return false
		}
	}
	return true
}

// Upstream (schemaType) -> downstream (structuredDatasetType)
func structuredDatasetCastFromSchema(upstream *nebula.SchemaType, downstream *nebula.StructuredDatasetType) bool {
	if len(upstream.Columns) == 0 || len(downstream.Columns) == 0 {
		return true
	}
	nameToTypeMap := make(map[string]nebula.SchemaType_SchemaColumn_SchemaColumnType)
	for _, column := range upstream.Columns {
		nameToTypeMap[column.Name] = column.GetType()
	}

	// Check that the downstream structuredDataset is a strict sub-set of the upstream schema.
	for _, column := range downstream.Columns {
		upstreamType, ok := nameToTypeMap[column.Name]
		if !ok {
			return false
		}
		if !schemaTypeIsMatchStructuredDatasetType(upstreamType, column.LiteralType.GetSimple()) {
			return false
		}
	}
	return true
}

// Upstream (structuredDatasetType) -> downstream (schemaType)
func schemaCastFromStructuredDataset(upstream *nebula.StructuredDatasetType, downstream *nebula.SchemaType) bool {
	if len(upstream.Columns) == 0 || len(downstream.Columns) == 0 {
		return true
	}
	nameToTypeMap := make(map[string]nebula.SimpleType)
	for _, column := range upstream.Columns {
		nameToTypeMap[column.Name] = column.LiteralType.GetSimple()
	}

	// Check that the downstream schema is a strict sub-set of the upstream structuredDataset.
	for _, column := range downstream.Columns {
		upstreamType, ok := nameToTypeMap[column.Name]
		if !ok {
			return false
		}
		if !schemaTypeIsMatchStructuredDatasetType(column.GetType(), upstreamType) {
			return false
		}
	}
	return true
}

func schemaTypeIsMatchStructuredDatasetType(schemaType nebula.SchemaType_SchemaColumn_SchemaColumnType, structuredDatasetType nebula.SimpleType) bool {
	switch schemaType {
	case nebula.SchemaType_SchemaColumn_INTEGER:
		return structuredDatasetType == nebula.SimpleType_INTEGER
	case nebula.SchemaType_SchemaColumn_FLOAT:
		return structuredDatasetType == nebula.SimpleType_FLOAT
	case nebula.SchemaType_SchemaColumn_STRING:
		return structuredDatasetType == nebula.SimpleType_STRING
	case nebula.SchemaType_SchemaColumn_BOOLEAN:
		return structuredDatasetType == nebula.SimpleType_BOOLEAN
	case nebula.SchemaType_SchemaColumn_DATETIME:
		return structuredDatasetType == nebula.SimpleType_DATETIME
	case nebula.SchemaType_SchemaColumn_DURATION:
		return structuredDatasetType == nebula.SimpleType_DURATION
	}
	return false
}

func isNoneType(t *nebula.LiteralType) bool {
	switch t.GetType().(type) {
	case *nebula.LiteralType_Simple:
		return t.GetSimple() == nebula.SimpleType_NONE
	default:
		return false
	}
}

func getTypeChecker(t *nebula.LiteralType) typeChecker {
	switch t.GetType().(type) {
	case *nebula.LiteralType_CollectionType:
		return collectionTypeChecker{
			literalType: t,
		}
	case *nebula.LiteralType_MapValueType:
		return mapTypeChecker{
			literalType: t,
		}
	case *nebula.LiteralType_Schema:
		return schemaTypeChecker{
			literalType: t,
		}
	case *nebula.LiteralType_UnionType:
		return unionTypeChecker{
			literalType: t,
		}
	case *nebula.LiteralType_StructuredDatasetType:
		return structuredDatasetChecker{
			literalType: t,
		}
	default:
		if isNoneType(t) {
			return noneTypeChecker{}
		}

		return trivialChecker{
			literalType: t,
		}
	}
}

func AreTypesCastable(upstreamType, downstreamType *nebula.LiteralType) bool {
	typeChecker := getTypeChecker(downstreamType)

	// if upstream is a singular union we check if the downstream type is castable from the union variant
	if upstreamType.GetUnionType() != nil && len(upstreamType.GetUnionType().GetVariants()) == 1 {
		variants := upstreamType.GetUnionType().GetVariants()
		if len(variants) == 1 && typeChecker.CastsFrom(variants[0]) {
			return true
		}
	}

	return typeChecker.CastsFrom(upstreamType)
}
