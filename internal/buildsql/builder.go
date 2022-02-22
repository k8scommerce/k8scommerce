package buildsql

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

//
// Sample query string formats
// Delimiter is hyphen: http://www.blooberry.com/indexdot/html/topics/urlencoding.htm
//
// Filter: firstName = 'bob' ORDER BY 'id' DESC
// Protip: the '-' sign prefixing the 'id' field indicates a DESC
// no prefix indicates an ASC
//
// https://example.org/?filter=u-firstName-bob&sortOn=-u-id
//
// filter: field format is: 'table prefix' 'hyphen' 'fieldname' 'hyphen' 'field value'
// Example: u-firstName-bob
// u-firstName-bob =		 u      		-      	firstName       -			bob
// 						 	 |				|			|			|			 |
//  					table prefix	hyphen  fieldName	hyphen	 field value
//
//
// sortOn: field format is: 'optional ASC/DESC prefix' 'table prefix' 'hyphen' 'fieldname'
// Example: u-firstName-bob
// -u-id =		 			- 						u      		-      		   id
// 						 	|						|			|				|
//  			optional ASC/DESC prefix		table prefix  hyphen	fieldName
//
//
//
// Filter: firstName = 'bob' AND lastName = 'philips' ORDER BY 'id' DESC
// https://example.org/?filter=u-firstName-bob&filter=u-lastName-philips&sortOn=u-lastName&sortOn=-u-firstName
//
// Assume the filter is always an "AND"
// check the allowedFields for the fieldnames
// return an error if a unknown fieldname
//
// In both AllowedFilterFields and AllowedSortFields
// the map[string]string maps to:
// map [string]		 string
// 		  |            |
//    fieldName   table alias
//
// Example:
// map[string]string{
//		"id":     "p",  // product alias
//		"name":   "p",  // product alias
//		"slug":   "p",  // product alias
//		"sku":    "v",  // product alias
//		"amount": "pr", // price alias
//	}
//

var Delimiter string = "-"

type SortDirection string

const (
	ASC  SortDirection = "ASC"
	DESC SortDirection = "DESC"
)

type FilterField struct {
	TableAlias string
	FieldName  string
	Operator   Operator
	Value      interface{}
}
type SortField struct {
	TableAlias string
	FieldName  string
	Direction  SortDirection
}

type Where struct {
	CombinedName string
	SqlString    string
	Named        string
}

func NewQueryBuilder() QueryBuilder {
	return QueryBuilder{}
}

type QueryBuilder struct {
	AllowedFilterFields map[string]string
	AllowedSortFields   map[string]string
	Filters             []FilterField
	Sorts               []SortField
	SearchTables        map[string]int
}

// AllowedFiltersFieldsFromMap
// resets AllowedFilterFields
// example:
// map[string]string{
//		"id":     "p",  // product alias
//		"name":   "p",  // product alias
//		"slug":   "p",  // product alias
//		"sku":    "v",  // product alias
//		"amount": "pr", // price alias
//	}
// func (b *QueryBuilder) AllowedFiltersFieldsFromStringMap(allowed map[string]string) {
// 	b.AllowedFilterFields = allowed
// }

func (b *QueryBuilder) ParseParamString(paramString string) error {
	if paramString == "" {
		paramString = "?"
	}
	// fmt.Println("paramString: ", paramString)

	b.SearchTables = make(map[string]int)

	if strings.Index(paramString, "?") != 0 {
		pathParts := strings.Split(paramString, "?")

		// fmt.Println("pathParts", pathParts)
		if len(pathParts) > 1 {
			paramString = pathParts[1]
		}
	}

	prefix := paramString[0:1]
	if prefix != "?" {
		paramString = "?" + paramString
	}

	// let's let the url parser do the work
	u, err := url.Parse(paramString)
	if err != nil {
		return err
	}
	q := u.Query()
	// fmt.Println(q)

	// fmt.Println("Q:", q)

	// parse filters
	if filters, ok := q["filter"]; ok {
		count := 0
		for _, filter := range filters {
			filter = strings.TrimSpace(filter)
			parts := strings.Split(filter, Delimiter)

			if len(parts) < 3 {
				return fmt.Errorf("filter: %s has too few params", filter)
			}

			// we don't care if the parts len is longer,
			// it could mean the filter content has a hyphen
			// so we'll keep every remaining after the second hyphen
			// this works well for sku numbers with hyphens
			val := strings.SplitAfterN(filter, Delimiter, 4)

			value := val[3]
			if Operator(parts[2]).isLike() {
				value = "%" + val[3] + "%"
			}

			filterField := FilterField{
				TableAlias: parts[0],
				FieldName:  parts[1],
				Operator:   Operator(parts[2]),
				Value:      value,
			}
			b.SearchTables[filterField.TableAlias] = count + 1
			b.Filters = append(b.Filters, filterField)
		}
	}

	// parse sorts
	if sortOns, ok := q["sortOn"]; ok {
		count := 0
		for _, sort := range sortOns {
			// check for the direction first
			// since the delimiter is the same as the
			// sort direction prefix
			sort := strings.TrimSpace(sort)
			dir := ASC
			if isDesc := strings.HasPrefix(sort, "-"); isDesc {
				dir = DESC
				sort = sort[1:]
			}

			parts := strings.Split(sort, Delimiter)
			if len(parts) < 1 {
				return fmt.Errorf("sortOn: %s has too few params", sort)
			}

			sortField := SortField{
				TableAlias: parts[0],
				FieldName:  parts[1],
				Direction:  dir,
			}
			b.SearchTables[sortField.TableAlias] = count + 1
			b.Sorts = append(b.Sorts, sortField)
		}
	}

	// fmt.Printf("\n#%+v", b.Filters)
	// fmt.Printf("\n#%+v\n\n", b.Sorts)
	return nil
}

// AllowedFiltersFieldsFromReflectionMap
// resets AllowedFilterFields
// the map takes two fields: string key and an interface
// the key maps to the table alias
// the interface is a struct with 'json', 'db' tags
// it uses reflection to determin the allowed fields
func (b *QueryBuilder) Build(paramString string, allowed map[string]interface{}) (where string, orderBy string, namedParamMap map[string]interface{}, err error) {
	namedParamMap = make(map[string]interface{})
	wheres := make(map[string][]Where)
	sb := []string{}

	if err := b.ParseParamString(paramString); err != nil {
		return "", "", nil, err
	}

	fieldsByTableAlias := make(map[string][]FilterField)
	for _, filter := range b.Filters {
		fieldsByTableAlias[filter.FieldName] = append(fieldsByTableAlias[filter.FieldName], filter)
	}

	sortsByTableAlias := make(map[string][]SortField)
	for _, sort := range b.Sorts {
		sortsByTableAlias[sort.FieldName] = append(sortsByTableAlias[sort.FieldName], sort)
	}

	for tableAlias, tableStruct := range allowed {
		// let's not bother with the table unless we're actually referencing it
		// let's also keep track of how many fields we have yet to match
		// we created a count in the b.SearchTables[tableAlias] map
		// when we hit zero we'll move on to the next table instead
		// of continuing out looping
		// remainingFields, ok := b.SearchTables[tableAlias]
		// if !ok {
		// 	continue
		// }

		rv := reflect.ValueOf(tableStruct)
		for i := 0; i < rv.NumField(); i++ {
			tag := rv.Type().Field(i).Tag.Get("db")
			if tag == "" {
				continue
			}

			fields, ok := fieldsByTableAlias[tag]
			if ok {
				for i, field := range fields {
					if field.TableAlias == tableAlias {
						// we use named pairs here with sqlx
						namedParam := fmt.Sprintf("filter_%s_%s_%d", field.TableAlias, field.FieldName, i)
						namedParamMap[namedParam] = field.Value
						sqlString := fmt.Sprintf("%s.%s %s :%s", field.TableAlias, field.FieldName, field.Operator.Convert(), namedParam)
						combined := fmt.Sprintf("%s.%s", tableAlias, field.FieldName)
						wheres[combined] = append(wheres[combined], Where{
							CombinedName: fmt.Sprintf("%s.%s", tableAlias, field.FieldName),
							SqlString:    sqlString,
							Named:        namedParam,
						})
					}
				}
			}

			sorts, ok := sortsByTableAlias[tag]
			if ok {
				for _, sort := range sorts {
					if sort.TableAlias == tableAlias {
						sb = append(sb, fmt.Sprintf("%s.%s %s", tableAlias, sort.FieldName, sort.Direction))
					}
				}
			}
		}
	}

	where = b.AssembledWheres(wheres)
	// orderBy = fmt.Sprintf("ORDER BY %s", strings.Join(sb, ", "))
	orderBy = strings.Join(sb, ", ")
	if orderBy != "" {
		orderBy = fmt.Sprintf("ORDER BY %s", orderBy)
	}

	return where, orderBy, namedParamMap, err
}

func (b *QueryBuilder) AssembledWheres(whereMap map[string][]Where) string {
	where := []string{}
	for _, ws := range whereMap {
		if len(ws) > 1 {
			orGroup := []string{}
			for _, w := range ws {
				orGroup = append(orGroup, w.SqlString)
			}
			where = append(where, "("+strings.Join(orGroup, " OR ")+")")
		} else {
			where = append(where, ws[0].SqlString)
		}
	}

	out := strings.Join(where, " AND ")
	if out != "" {
		return fmt.Sprintf(" AND %s", out)
	}
	return ""
}

func BuildOrderBy(on string, allowedFields map[string]string) (orderBy string, err error) {
	if on == "" {
		return "", nil
	}

	var sb []string
	fields := strings.Split(strings.ToLower(on), ",")

	// fmt.Println("FIELDS: ", fields)

	for _, field := range fields {
		field = strings.TrimSpace(field)
		dir := "ASC"
		fieldName := field
		if isDesc := strings.HasPrefix(field, "-"); isDesc {
			dir = "DESC"
			fieldName = field[1:]
		}

		// fmt.Println("fieldName: ", fieldName)

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

	orderBy = strings.Join(sb, ", ")
	if orderBy != "" {
		orderBy = fmt.Sprintf("ORDER BY %s", orderBy)
	}

	return orderBy, err
}

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
