CREATE TABLE public."session" (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	user_id uuid NOT NULL,
	expires_in bigint NOT NULL,
	created timestamp with time zone NOT NULL DEFAULT now(),
	CONSTRAINT session_pk PRIMARY KEY (id),
	CONSTRAINT session_fk FOREIGN KEY (user_id) REFERENCES public."user"(id)
);

-- Column comments

COMMENT ON COLUMN public."session".id IS 'RefreshToken value';
COMMENT ON COLUMN public."session".user_id IS 'FK of the user';