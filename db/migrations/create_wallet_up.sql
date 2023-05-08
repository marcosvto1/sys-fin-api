CREATE TABLE
  "public"."wallets" (
    "id" serial primary key,
    "created_at" timestamp not null default NOW(),
    "name" varchar(255) null,
    "user_id" int null,
    "updated_at" TIMESTAMPTZ null
  )