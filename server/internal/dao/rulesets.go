// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"server/internal/dao/internal"
)

// internalRulesetsDao is internal type for wrapping internal DAO implements.
type internalRulesetsDao = *internal.RulesetsDao

// rulesetsDao is the data access object for table rulesets.
// You can define custom methods on it to extend its functionality as you wish.
type rulesetsDao struct {
	internalRulesetsDao
}

var (
	// Rulesets is globally public accessible object for table rulesets operations.
	Rulesets = rulesetsDao{
		internal.NewRulesetsDao(),
	}
)

// Fill with you ideas below.
