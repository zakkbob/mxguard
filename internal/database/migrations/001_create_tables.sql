CREATE TABLE user (
   id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
   is_admin bool NOT NULL,
   username varchar(20) NOT NULL UNIQUE
);

CREATE TABLE subdomain (
   id int SERIAL PRIMARY KEY,
   subdomain varchar(255), -- max subdomain length is 63 characters, but this will store nested ones too e.g 'dave.steve'
   verified bool NOT NULL
);

CREATE TABLE email (
   id int SERIAL PRIMARY KEY,
   email varchar(320) UNIQUE NOT NULL,
   user uuid references user(id) NOT NULL,
   verified bool NOT NULL
);

CREATE TABLE alias (
   id int SERIAL PRIMARY KEY NOT NULL,
   active bool NOT NULL,
   user uuid references user(id) NOT NULL,
   subdomain int references subdomain(id) NOT NULL,
   local_part varchar(64) NOT NULL,
   UNIQUE (subdomain, local_part)
);