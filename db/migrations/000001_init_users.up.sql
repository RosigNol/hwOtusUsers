CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
    "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "name" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "role" VARCHAR NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "users_email_key" ON "users"("email");

INSERT INTO "users" (
  id,
  name,
  email,
  password,
  role,
  updated_at
) VALUES (
  'f525b2e1-d311-471c-a935-604412bc07bb', 'test', 'test@test.ru', '12345', 'user', '2024-04-05 02:20:00.298'
),(
  'a98489a4-9c6b-4b7c-9946-da94a9a6c79d', 'test2', 'test2@test.ru', '12345', 'user', '2024-04-05 02:20:00.298'
),(
  '9ef8ca20-31a0-4479-87e4-f1cc7170e51b', 'test3', 'test3@test.ru', '12345', 'user', '2024-04-05 02:20:00.298'
),(
  '892612a7-6f3d-403b-848d-8c753dd74ad7', 'test4', 'test4@test.ru', '12345', 'user', '2024-04-05 02:20:00.298'
);
