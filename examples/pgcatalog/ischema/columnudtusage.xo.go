// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/mbilski/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// ColumnUdtUsage represents a row from 'information_schema.column_udt_usage'.
type ColumnUdtUsage struct {
	UdtCatalog   pgtypes.SQLIdentifier `json:"udt_catalog"`   // udt_catalog
	UdtSchema    pgtypes.SQLIdentifier `json:"udt_schema"`    // udt_schema
	UdtName      pgtypes.SQLIdentifier `json:"udt_name"`      // udt_name
	TableCatalog pgtypes.SQLIdentifier `json:"table_catalog"` // table_catalog
	TableSchema  pgtypes.SQLIdentifier `json:"table_schema"`  // table_schema
	TableName    pgtypes.SQLIdentifier `json:"table_name"`    // table_name
	ColumnName   pgtypes.SQLIdentifier `json:"column_name"`   // column_name
}
