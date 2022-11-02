CREATE TABLE "locations" (
                             "id" bigserial PRIMARY KEY,
                             "name" varchar NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "events" (
                          "id" bigserial PRIMARY KEY,
                          "location_id" bigint NOT NULL,
                          "name" bigint NOT NULL,
                          "datetime" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "events" ("location_id");

ALTER TABLE "events" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id");
