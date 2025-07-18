CREATE TABLE IF NOT EXISTS usr (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	is_admin bool NOT NULL,
	username varchar(20) NOT NULL UNIQUE,
	email varchar(254) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS alias (
	usr_id uuid references usr(id),
	name varchar(20) NOT NULL UNIQUE,
	description varchar(255) NOT NULL,
	enabled bool NOT NULL
);
