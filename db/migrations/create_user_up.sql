CREATE TABLE IF NOT EXISTS
  "public"."users" (
    "id" serial primary key,
    "email" varchar(255) null,
    "name" varchar(255) null,
    "password" varchar(255) null,
    "created_at" timestamp not null default NOW(),
    "updated_at" TIMESTAMP null
)
