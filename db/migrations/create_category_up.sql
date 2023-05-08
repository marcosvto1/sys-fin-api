CREATE TABLE
  "public"."categories" (
    "id" serial primary key,
    "created_at" timestamp not null default NOW(),
    "name" varchar(255) null,
    "updated_at" TIMESTAMP null
  )