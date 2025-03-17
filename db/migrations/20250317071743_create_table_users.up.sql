-- Active: 1738660987275@@127.0.0.1@5432@rania_eskristal
CREATE TABLE IF NOT EXISTS users(
  id UUID PRIMARY KEY NOT NULL,
  role_id UUID,
  name TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  username TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP,
  CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE SET NULL
);

