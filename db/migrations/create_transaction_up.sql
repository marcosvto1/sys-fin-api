CREATE TABLE
  "public"."transactions" (
    "id" serial primary key,
    "description" varchar(255) null,
    "amount" DECIMAL(10, 2) null,
    "category_id" INT null,
    "wallet_id" INT null,
    "transaction_at" DATE null,
    "created_at" timestamp not null default NOW(),
    "updated_at" TIMESTAMP null
  )