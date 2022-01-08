package repos

import (
	"fmt"
	"strings"
)

// func getWhere(table string, whereMap map[string]interface{}, orderByMap map[string]interface{}, out interface{}) error {
// 	var where []string
// 	for item, _ := range whereMap {
// 		where = append(where, fmt.Sprintf("%s = :%s", item, item))
// 	}

// 	var orderBy []string
// 	if len(orderByMap) > 0 {
// 		orderBy = append(orderBy, "ORDER BY 1")
// 	}
// 	for item, direction := range orderByMap {
// 		orderBy = append(orderBy, fmt.Sprintf("%s %s", item, direction))
// 	}

// 	query := fmt.Sprintf(`
// 			SELECT * FROM %s
// 			WHERE %s
// 			%s
// 		`,
// 		table,
// 		strings.Join(where, " AND "),
// 		strings.Join(orderBy, ", "))

// 	fmt.Println("query: ", query)

// 	nstmt, err := a.db.PrepareNamed(query)
// 	if err != nil {
// 		return fmt.Errorf("error::getWhere::%s", err.Error())
// 	}
// 	err = nstmt.Get(out, whereMap)
// 	return err
// }

func BuildOrderBy(on string, allowedFields map[string]string) (orderBy string, err error) {
	if on == "" {
		return "", nil
	}

	fmt.Println("ON: ", on)

	var sb []string
	fields := strings.Split(strings.ToLower(on), ",")

	for _, field := range fields {
		field = strings.TrimSpace(field)
		dir := "ASC"
		fieldName := field
		if isDesc := strings.HasPrefix(field, "-"); isDesc {
			dir = "DESC"
			fieldName = field[1:]
		}

		allowed := false
		for allowedField, tableName := range allowedFields {
			if fieldName == allowedField {
				sb = append(sb, fmt.Sprintf("%s.%s %s", tableName, fieldName, dir))
				allowed = true
				break
			}
		}

		if !allowed {
			return "", fmt.Errorf("error: %s is not allowed to be sorted on", fieldName)
		}
	}

	if len(sb) == 0 {
		return orderBy, err
	}

	orderBy = fmt.Sprintf("ORDER BY %s", strings.Join(sb, ", "))
	return orderBy, err
}
