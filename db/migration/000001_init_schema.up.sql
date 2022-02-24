CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "content" text NOT NULL,
  "org_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "deleted_comments" (
  "id" bigserial PRIMARY KEY,
  "content" text NOT NULL,
  "org_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "comments" ("org_name");

CREATE FUNCTION moveDeleted() RETURNS trigger AS $$
  BEGIN
    INSERT INTO "deleted_comments" VALUES((OLD).*);
    RETURN OLD;
  END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER moveDeleted
BEFORE DELETE ON "comments"
FOR EACH ROW
EXECUTE PROCEDURE moveDeleted();