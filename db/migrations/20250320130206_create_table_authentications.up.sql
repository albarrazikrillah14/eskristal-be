-- Active: 1738660987275@@127.0.0.1@5432@rania_eskristal
CREATE TABLE IF NOT EXISTS authentications (
  id SERIAL PRIMARY KEY,
  token TEXT NOT NULL
);

CREATE INDEX token_idx ON authentications(token);
