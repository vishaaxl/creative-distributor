CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    phoneNumber text UNIQUE CHECK (phoneNumber ~ '^[789]\d{9}$'),
    activated bool NOT NULL,
    version integer NOT NULL DEFAULT 1
);
