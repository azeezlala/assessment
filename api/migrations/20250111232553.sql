-- Create "customers" table
CREATE TABLE IF NOT EXISTS "public"."customers" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  "email" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_customers_email" UNIQUE ("email")
);
-- Create index "idx_customers_email" to table: "customers"
CREATE INDEX IF NOT EXISTS "idx_customers_email" ON "public"."customers" ("email");
-- Create "resources" table
CREATE TABLE IF NOT EXISTS "public"."resources" (
  "id" uuid NOT NULL,
  "name" text NULL,
  "type" text NULL,
  "region" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_resources_name" UNIQUE ("name")
);
-- Create index "idx_resources_name" to table: "resources"
CREATE INDEX IF NOT EXISTS "idx_resources_name" ON "public"."resources" ("name");
-- Create index "idx_resources_region" to table: "resources"
CREATE INDEX IF NOT EXISTS "idx_resources_region" ON "public"."resources" ("region");
-- Create index "idx_resources_type" to table: "resources"
CREATE INDEX IF NOT EXISTS "idx_resources_type" ON "public"."resources" ("type");
-- Create "customer_resources" table
CREATE TABLE IF NOT EXISTS "public"."customer_resources" (
  "id" uuid NOT NULL,
  "customer_id" uuid NOT NULL,
  "resource_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_customer_resources_customer" FOREIGN KEY ("customer_id") REFERENCES "public"."customers" ("id") ON UPDATE CASCADE ON DELETE SET NULL,
  CONSTRAINT "fk_customer_resources_resources" FOREIGN KEY ("resource_id") REFERENCES "public"."resources" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
-- Create index "idx_customer_resource" to table: "customer_resources"
CREATE UNIQUE INDEX IF NOT EXISTS "idx_customer_resource" ON "public"."customer_resources" ("customer_id", "resource_id");
-- Create index "idx_customer_resources_customer_id" to table: "customer_resources"
CREATE INDEX IF NOT EXISTS "idx_customer_resources_customer_id" ON "public"."customer_resources" ("customer_id");
-- Create index "idx_customer_resources_resource_id" to table: "customer_resources"
CREATE INDEX IF NOT EXISTS "idx_customer_resources_resource_id" ON "public"."customer_resources" ("resource_id");
