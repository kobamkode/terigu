CREATE TABLE public.users (
	id bigserial NOT NULL,
	"name" varchar NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);
