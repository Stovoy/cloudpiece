DROP TABLE post;

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
