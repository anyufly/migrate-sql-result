package result

type SQLForTable struct {
	table string
	sql   string
}

func NewSQLForTable(table string, sql string) *SQLForTable {
	return &SQLForTable{
		table: table,
		sql:   sql,
	}
}

func (s *SQLForTable) Table() string {
	return s.table
}

func (s *SQLForTable) SQL() string {
	return s.sql
}

type SQLForTableList []*SQLForTable

func (sqlList SQLForTableList) ToMap() map[string][]string {
	sqlMap := make(map[string][]string)
	for _, sql := range sqlList {
		if sql != nil {
			if tableSQLList, ok := sqlMap[sql.Table()]; ok {
				tableSQLList = append(tableSQLList, sql.SQL())
				sqlMap[sql.Table()] = tableSQLList
			} else {
				sqlMap[sql.Table()] = []string{sql.SQL()}
			}
		}
	}
	return sqlMap
}

type MigrateSQLResult struct {
	up   SQLForTableList
	down SQLForTableList
}

func NewMigrateSQLResult() *MigrateSQLResult {
	return &MigrateSQLResult{
		up:   make([]*SQLForTable, 0, 10),
		down: make([]*SQLForTable, 0, 10),
	}
}

func (result *MigrateSQLResult) AppendUp(sqlList ...*SQLForTable) {
	for _, sql := range sqlList {
		if sql != nil {
			result.up = append(result.up, sql)
		}
	}
}

func (result *MigrateSQLResult) Up() map[string][]string {
	return result.up.ToMap()
}

func (result *MigrateSQLResult) AppendDown(sqlList ...*SQLForTable) {
	for _, sql := range sqlList {
		if sql != nil {
			result.down = append(result.down, sql)
		}
	}

}

func (result *MigrateSQLResult) Down() map[string][]string {
	return result.down.ToMap()
}

func (result *MigrateSQLResult) Empty() bool {
	return len(result.up) == 0
}
