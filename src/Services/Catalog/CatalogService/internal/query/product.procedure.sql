CREATE OR REPLACE PROCEDURE create_product(
    p_title VARCHAR,
    p_author VARCHAR,
    p_publisher VARCHAR,
    p_publication_year INTEGER,
    p_page_count INT,
    p_dimensions VARCHAR,
    p_cover_type VARCHAR,
    p_price DECIMAL(10, 2),
    p_description TEXT,
    p_image_url VARCHAR DEFAULT NULL,
    p_original_owner_id UUID DEFAULT NULL
) AS $$
DECLARE
    new_product_id UUID;
BEGIN
    new_product_id := gen_random_uuid();
    
    INSERT INTO products (
        id,
        title,
        author,
        publisher,
        publication_year,
        page_count,
        dimensions,
        cover_type,
        price,
        description,
        image_url,
        original_owner_id
    ) VALUES (
        new_product_id,
        p_title,
        p_author,
        p_publisher,
        p_publication_year,
        p_page_count,
        p_dimensions,
        p_cover_type,
        p_price,
        p_description,
        p_image_url,
        p_original_owner_id
    );
    RAISE NOTICE 'New product created with ID: %', new_product_id;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE update_product(
    p_id UUID, 
    p_title VARCHAR,
    p_author VARCHAR,
    p_publisher VARCHAR,
    p_publication_year INTEGER,
    p_page_count INT,
    p_dimensions VARCHAR,
    p_cover_type VARCHAR,
    p_price DECIMAL(10, 2),
    p_description TEXT,
    p_discount_percentage INT,
    p_product_type INT,
    p_is_active BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE products
    SET
        title = p_title,
        author = p_author,
        publisher = p_publisher,
        publication_year = p_publication_year,
        page_count = p_page_count,
        dimensions = p_dimensions,
        cover_type = p_cover_type,
        price = p_price,
        description = p_description,
        discount_percentage = p_discount_percentage,
        product_type = p_product_type,
        is_active = p_is_active,
        updated_at = CURRENT_TIMESTAMP 
    WHERE id = p_id;

    RAISE NOTICE 'Product with ID % has been updated', p_id;
    
EXCEPTION WHEN OTHERS THEN
    RAISE 'Error updating product with ID %', p_id;
END;
$$;

CREATE OR REPLACE PROCEDURE delete_product(
    p_id UUID
) AS $$
BEGIN
    DELETE FROM products
    WHERE id = p_id;

    RAISE NOTICE 'Product with ID: % has been deleted.', p_id;
END;
$$ LANGUAGE plpgsql;


-- DROP PROCEDURE update_product
