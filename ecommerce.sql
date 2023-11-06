CREATE TABLE "order_items" (
  "id" BIGSERIAL PRIMARY KEY,
  "order_id" integer NOT NULL,
  "image" varchar NOT NULL,
  "product_id" integer NOT NULL,
  "quantity" integer NOT NULL,
  "price" decimal NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "is_admin" bool NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "orders" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" integer NOT NULL,
  "shipping_address" varchar NOT NULL,
  "shipping_city" varchar NOT NULL,
  "shipping_postal_code" varchar NOT NULL,
  "shipping_country" varchar NOT NULL,
  "payment_method" varchar NOT NULL,
  "payment_id" varchar NOT NULL,
  "payment_status" varchar NOT NULL,
  "payment_update_time" varchar NOT NULL,
  "payment_email_address" varchar NOT NULL,
  "items_price" decimal NOT NULL,
  "tax_price" decimal NOT NULL,
  "shipping_price" decimal NOT NULL,
  "total_price" decimal NOT NULL,
  "is_paid" bool NOT NULL DEFAULT false,
  "paid_at" timestamp,
  "is_delivered" bool NOT NULL DEFAULT false,
  "delivered_at" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "products" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "brand" varchar NOT NULL,
  "category" varchar NOT NULL,
  "image" varchar NOT NULL,
  "rating" decimal,
  "num_reviews" integer NOT NULL,
  "price" decimal NOT NULL,
  "count_in_stock" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "reviews" (
  "id" BIGSERIAL PRIMARY KEY,
  "rating" decimal NOT NULL,
  "comment" varchar NOT NULL,
  "user_id" integer NOT NULL,
  "product_id" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
