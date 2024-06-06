CREATE TABLE if NOT EXISTS env (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE,
    origin_host VARCHAR(255) UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


CREATE TABLE if NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    env_id integer REFERENCES env(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE if NOT EXISTS feature_flags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    value BOOLEAN DEFAULT true,
    env_id integer REFERENCES env (id) ON DELETE CASCADE,
    group_id integer REFERENCES groups (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
