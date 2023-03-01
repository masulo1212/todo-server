CREATE TABLE "pending_tasks" (
  "id" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "date" varchar NOT NULL,
  "is_done" bool NOT NULL DEFAULT false,
  "is_delete" bool NOT NULL DEFAULT false,
  "is_favorite" bool NOT NULL DEFAULT false
);

CREATE TABLE "complete_tasks" (
  "id" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "date" varchar NOT NULL,
  "is_done" bool NOT NULL DEFAULT false,
  "is_delete" bool NOT NULL DEFAULT false,
  "is_favorite" bool NOT NULL DEFAULT false
);

CREATE TABLE "fav_tasks" (
  "id" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "date" varchar NOT NULL,
  "is_done" bool NOT NULL DEFAULT false,
  "is_delete" bool NOT NULL DEFAULT false,
  "is_favorite" bool NOT NULL DEFAULT false
);

CREATE TABLE "bin_tasks" (
  "id" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "date" varchar NOT NULL,
  "is_done" bool NOT NULL DEFAULT false,
  "is_delete" bool NOT NULL DEFAULT false,
  "is_favorite" bool NOT NULL DEFAULT false
);
