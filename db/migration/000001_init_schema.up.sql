CREATE TABLE products (
  "id" bigserial PRIMARY KEY,
  "sku" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE countries (
  "id" bigserial PRIMARY KEY,
  "country_code" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE stocks (
  "id" bigserial PRIMARY KEY,
  "quantity" bigint NOT NULL DEFAULT 0 CHECK (quantity >=0),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "product_id" bigint NOT NULL,
  "country_id" bigint NOT NULL
);

CREATE UNIQUE INDEX ON "products" ("sku");

CREATE UNIQUE INDEX ON "countries" ("country_code");

CREATE UNIQUE INDEX ON "stocks" ("product_id", "country_id");

COMMENT ON COLUMN "stocks"."quantity" IS 'Must be 0 or positive';

ALTER TABLE "stocks" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "stocks" ADD FOREIGN KEY ("country_id") REFERENCES "countries" ("id");