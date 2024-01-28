CREATE TABLE "users" (
		"id" BIGSERIAL PRIMARY KEY,
		"name" varchar NOT NULL,
		"surname" varchar NOT NULL,
		"address" varchar NOT NULL,
		"notification_radius" smallint NOT NULL DEFAULT 300,
		"email" varchar NOT NULL,
		"password" varchar NOT NULL,
		"role" varchar NOT NULL,
		"created_at" timestamptz NOT NULL DEFAULT (now()),
		"updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX "email" ON "users" ("email");