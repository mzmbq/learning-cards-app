CREATE TABLE IF NOT EXISTS cards (
    id bigserial PRIMARY KEY,
    front text NOT NULL,
    back text NOT NULL,

    deck_id bigserial NOT NULL
);