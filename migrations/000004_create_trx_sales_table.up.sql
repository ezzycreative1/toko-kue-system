CREATE TABLE IF NOT EXISTS public.trx_sales (
    id bigserial PRIMARY KEY,
    tanggal timestamp NULL,
    product_id bigserial NOT NULL,
    quantity integer NULL DEFAULT 0,
    total_price numeric DEFAULT 0,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(36) NULL,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(36) NULL,
    deleted_at timestamp NULL,
    deleted_by varchar(36) NULL,
    CONSTRAINT fk_trx_sales_product_id FOREIGN KEY (product_id) REFERENCES public.tmst_products(id) ON DELETE CASCADE
);

CREATE INDEX trx_sales_product_id_idx ON public.trx_sales (product_id);