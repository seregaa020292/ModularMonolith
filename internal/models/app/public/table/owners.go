//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Owners = newOwnersTable("public", "owners", "")

type ownersTable struct {
	postgres.Table

	// Columns
	ID            postgres.ColumnString
	FullName      postgres.ColumnString
	LicenseNumber postgres.ColumnString
	Phone         postgres.ColumnString
	Email         postgres.ColumnString
	CreatedAt     postgres.ColumnTimestampz
	UpdatedAt     postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OwnersTable struct {
	ownersTable

	EXCLUDED ownersTable
}

// AS creates new OwnersTable with assigned alias
func (a OwnersTable) AS(alias string) *OwnersTable {
	return newOwnersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OwnersTable with assigned schema name
func (a OwnersTable) FromSchema(schemaName string) *OwnersTable {
	return newOwnersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OwnersTable with assigned table prefix
func (a OwnersTable) WithPrefix(prefix string) *OwnersTable {
	return newOwnersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OwnersTable with assigned table suffix
func (a OwnersTable) WithSuffix(suffix string) *OwnersTable {
	return newOwnersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOwnersTable(schemaName, tableName, alias string) *OwnersTable {
	return &OwnersTable{
		ownersTable: newOwnersTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newOwnersTableImpl("", "excluded", ""),
	}
}

func newOwnersTableImpl(schemaName, tableName, alias string) ownersTable {
	var (
		IDColumn            = postgres.StringColumn("id")
		FullNameColumn      = postgres.StringColumn("full_name")
		LicenseNumberColumn = postgres.StringColumn("license_number")
		PhoneColumn         = postgres.StringColumn("phone")
		EmailColumn         = postgres.StringColumn("email")
		CreatedAtColumn     = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn     = postgres.TimestampzColumn("updated_at")
		allColumns          = postgres.ColumnList{IDColumn, FullNameColumn, LicenseNumberColumn, PhoneColumn, EmailColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns      = postgres.ColumnList{FullNameColumn, LicenseNumberColumn, PhoneColumn, EmailColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return ownersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:            IDColumn,
		FullName:      FullNameColumn,
		LicenseNumber: LicenseNumberColumn,
		Phone:         PhoneColumn,
		Email:         EmailColumn,
		CreatedAt:     CreatedAtColumn,
		UpdatedAt:     UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
