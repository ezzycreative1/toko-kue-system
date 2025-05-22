CREATE TABLE IF NOT EXISTS public.tmst_recipes (
    id bigserial PRIMARY KEY,
    product_id bigserial NOT NULL,
    ingredient_id BIGSERIAL NOT NULL,
    quantity integer DEFAULT 0,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(36) NULL,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(36) NULL,
    deleted_at timestamp NULL,
    deleted_by varchar(36) NULL,
    CONSTRAINT fk_tmst_recipes_product_id FOREIGN KEY (product_id) REFERENCES public.tmst_products(id) ON DELETE CASCADE,
    CONSTRAINT fk_tmst_recipes_ingredient_id FOREIGN KEY (ingredient_id) REFERENCES public.tmst_ingredients(id) ON DELETE CASCADE
);

CREATE INDEX tmst_recipes_product_id_idx ON public.tmst_recipes (product_id);
CREATE INDEX tmst_recipes_ingredient_id_idx ON public.tmst_recipes (ingredient_id);