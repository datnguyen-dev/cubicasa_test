DROP TABLE IF EXISTS public.hubs;

CREATE TABLE public.hubs
(
    hub_id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default",
    geo_location text COLLATE pg_catalog."default",
    CONSTRAINT hubs_pkey PRIMARY KEY (hub_id)
);

DROP TABLE IF EXISTS public.teams;
CREATE TABLE public.teams
(
    team_id text COLLATE pg_catalog."default",
    name text COLLATE pg_catalog."default",
    type text COLLATE pg_catalog."default",
    hub_id text COLLATE pg_catalog."default",
    CONSTRAINT teams_pkey PRIMARY KEY (team_id)
);

DROP TABLE IF EXISTS public.users;
CREATE TABLE public.users
(
    user_id text COLLATE pg_catalog."default",
    role integer,
    email text COLLATE pg_catalog."default",
    team_id text COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);