CREATE TABLE post
(
  id serial PRIMARY KEY,
  title character varying(140) NOT NULL,
  body text NOT NULL
)
WITH (
  OIDS=FALSE
);
ALTER TABLE post
  OWNER TO postgres;

CREATE TABLE component
(
  id serial PRIMARY KEY,
  name character varying(80) NOT NULL,
  version int NOT NULL,
  date_created timestamp NOT NULL
)
