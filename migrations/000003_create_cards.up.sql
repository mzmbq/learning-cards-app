CREATE TABLE IF NOT EXISTS cards (
    id bigserial PRIMARY KEY,
    front text NOT NULL,
    back text NOT NULL,
    deck_id bigserial NOT NULL,

    ease real,
    interval bigint,
    "state" integer,
    step integer,
    due timestamp with time zone
);