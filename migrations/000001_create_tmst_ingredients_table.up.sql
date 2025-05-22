CREATE TABLE IF NOT EXISTS public.tmst_ingredients (
    id bigserial PRIMARY KEY,
    name varchar(255) NULL,
    unit varchar(255) NULL,
    price_per_unit numeric DEFAULT 0,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(36) NULL,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(36) NULL,
    deleted_at timestamp NULL,
    deleted_by varchar(36) NULL
);