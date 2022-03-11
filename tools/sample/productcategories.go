package main

import (
	"fmt"
)

func associateProductCategories(categories Categories) {
	for _, category := range categories {
		repo.GetRawDB().MustExec(fmt.Sprintf(`
			with cte AS (
				select id as product_id, %d as category_id from product
				order by random() 
				limit random_between(12, %d)
			)
			insert into product_category 
			(product_id, category_id) 
			select product_id, category_id from cte
			ON CONFLICT (product_id, category_id) DO NOTHING;
		`, category.ID, 100))
	}
}
