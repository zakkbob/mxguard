CREATE TABLE IF NOT EXISTS user(
   id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
   is_admin bool NOT NULL,
   username varchar(20) NOT NULL UNIQUE
);
