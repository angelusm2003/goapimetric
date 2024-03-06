-- Table: public.metrics

-- DROP TABLE IF EXISTS public.metrics;

CREATE TABLE IF NOT EXISTS public.metrics
(
    metric_id numeric,
    metric_name text COLLATE pg_catalog."default",
    date_creation date
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.metrics
    OWNER to postgres;