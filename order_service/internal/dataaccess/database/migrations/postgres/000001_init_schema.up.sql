CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "orders" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid,
  "address" varchar,
  "payment_method" varchar,
  "status" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "order_id" uuid,
  "sku_id" uuid,
  "quantity" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
