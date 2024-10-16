package query

const (
	GET_USER_BY_ID = `SELECT
						id, title, author, publisher, publication_year, page_count,
						dimensions, cover_type, price, description, image_url,
						sold_quantity, average_rating, quantity_evaluate, 
						discount_percentage, product_type, is_active, 
						original_owner_id, created_at, updated_at
					  FROM products WHERE id = $1;
		        	  `
	CREATE_PRODUCT     = `CALL create_product($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	UPDATE_PRODUCT     = `CALL update_product($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	DELETE_PRODUCT     = `CALL delete_product($1)`
	PAGINATION_PRODUCT = `SELECT * FROM UnifiedProductPagination($1, $2, $3, $4, $5, $6, $7)`
)
