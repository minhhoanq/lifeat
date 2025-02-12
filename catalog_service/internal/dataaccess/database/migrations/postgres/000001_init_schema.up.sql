CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "carts" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "cart_items" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "cart_id" uuid NOT NULL,
  "sku_id" uuid NOT NULL,
  "quantity" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "description" varchar,
  "image" varchar NOT NULL,
  "name" varchar NOT NULL,
  "category_id" int NOT NULL,
  "brand_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "brands" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "skus" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "product_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "slug" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sku_attributes" (
  "id" serial PRIMARY KEY,
  "sku_id" uuid NOT NULL,
  "attribute_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "attributes" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "prices" (
  "id" serial PRIMARY KEY,
  "sku_id" uuid NOT NULL,
  "original_price" int NOT NULL,
  "effective_date" timestamp NOT NULL,
  "active" bool NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "inventories" (
  "id" serial PRIMARY KEY,
  "sku_id" uuid NOT NULL,
  "original_price" int NOT NULL,
  "reservations" timestamp NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);


ALTER TABLE "cart_items" ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");

ALTER TABLE "cart_items" ADD FOREIGN KEY ("sku_id") REFERENCES "skus" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("brand_id") REFERENCES "brands" ("id");

ALTER TABLE "skus" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "sku_attributes" ADD FOREIGN KEY ("sku_id") REFERENCES "skus" ("id");

ALTER TABLE "sku_attributes" ADD FOREIGN KEY ("attribute_id") REFERENCES "attributes" ("id");

ALTER TABLE "prices" ADD FOREIGN KEY ("sku_id") REFERENCES "skus" ("id");

ALTER TABLE "inventories" ADD FOREIGN KEY ("sku_id") REFERENCES "skus" ("id");