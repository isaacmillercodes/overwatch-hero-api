DROP DATABASE IF EXISTS overwatch_api;
CREATE DATABASE overwatch_api;

\c overwatch_api;

CREATE EXTENSION pgcrypto;

CREATE TABLE heroes (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL,
  description text NOT NULL,
  role varchar(255) NOT NULL,
  type varchar(255) NOT NULL,
  image varchar(255) NOT NULL,
  difficulty int NOT NULL
);

CREATE TABLE game_types (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL,
  has_defense boolean NOT NULL
);

\i /Users/Isaac/go-workspace/src/overwatch-api/db/seed.sql;
