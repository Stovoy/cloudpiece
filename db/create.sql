CREATE TABLE post
(
  id integer NOT NULL,
  title character varying(140) NOT NULL,
  body text NOT NULL,
  CONSTRAINT id PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE post
  OWNER TO postgres;
