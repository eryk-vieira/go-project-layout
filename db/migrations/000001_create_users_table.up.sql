CREATE TABLE users (
  id uuid NOT NULL,
  name varchar(60) NOT NULL,
  username varchar(30) NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id)
)

