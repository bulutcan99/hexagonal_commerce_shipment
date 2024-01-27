CREATE TYPE "users_role_enum" AS ENUM ('admin', 'customer');

CREATE TABLE "users" (
		"id" BIGSERIAL PRIMARY KEY,
		"name" varchar NOT NULL,
		"surname" varchar NOT NULL,
		"address" varchar NOT NULL,
		"notification_radius" int NOT NULL,
		"email" varchar NOT NULL,
		"password" varchar NOT NULL,
		"role" varchar NOT NULL DEFAULT 'customer',
		"created_at" timestamptz NOT NULL DEFAULT (now()),
		"updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX "email" ON "users" ("email");
```