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

CREATE INDEX user_email_index ON users(email);
CREATE INDEX user_username_index ON users(username);


-- GENERATE 1jt user
-- INSERT INTO users (id, role_id, name, email, username, password)
-- SELECT 
--   gen_random_uuid(),
--   (CASE WHEN random() < 0.8 THEN (SELECT id FROM roles ORDER BY random() LIMIT 1) ELSE NULL END),
--   'User_' || gs::TEXT AS name,
--   'user_' || gs::TEXT || '@example.com' AS email,
--   'username_' || gs::TEXT AS username,
--   'password_' || md5(random()::TEXT) AS password
-- FROM generate_series(1, 1000000) AS gs;

-- -- Pastikan ekstensi pgcrypto aktif untuk menggunakan gen_random_uuid()
-- CREATE EXTENSION IF NOT EXISTS "pgcrypto";