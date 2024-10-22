CREATE TABLE IF NOT EXISTS decks (
    id bigserial PRIMARY KEY,
    name text NOT NULL,

    user_id bigserial NOT NULL,

    background text NOT NULL
);