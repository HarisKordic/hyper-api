-- Create Carbon Footprint History table
CREATE TABLE carbon_footprint_history (
    id SERIAL PRIMARY KEY,
    month DATE NOT NULL,
    amount DECIMAL(4,2) NOT NULL
);

-- Create Pollution Levels table
CREATE TABLE pollution_levels (
    id SERIAL PRIMARY KEY,
    pollutant VARCHAR(50) NOT NULL,
    level INTEGER NOT NULL,
    recorded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Map Users table
CREATE TABLE map_users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    latitude DECIMAL(9,6) NOT NULL,
    longitude DECIMAL(9,6) NOT NULL,
    activity TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    icon_type VARCHAR(50) NOT NULL,
    emission_amount DECIMAL(4,1) NOT NULL
);

-- Next auth
CREATE TABLE verification_token
(
    identifier TEXT        NOT NULL,
    expires    TIMESTAMPTZ NOT NULL,
    token      TEXT        NOT NULL,

    PRIMARY KEY (identifier, token)
);

CREATE TABLE accounts
(
    id                  SERIAL,
    "userId"            INTEGER      NOT NULL,
    type                VARCHAR(255) NOT NULL,
    provider            VARCHAR(255) NOT NULL,
    "providerAccountId" VARCHAR(255) NOT NULL,
    refresh_token       TEXT,
    access_token        TEXT,
    expires_at          BIGINT,
    id_token            TEXT,
    scope               TEXT,
    session_state       TEXT,
    token_type          TEXT,

    PRIMARY KEY (id)
);

CREATE TABLE sessions
(
    id             SERIAL,
    "userId"       INTEGER      NOT NULL,
    expires        TIMESTAMPTZ  NOT NULL,
    "sessionToken" VARCHAR(255) NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE users
(
    id              SERIAL,
    name            VARCHAR(255),
    email           VARCHAR(255),
    "emailVerified" TIMESTAMPTZ,
    image           TEXT,

    PRIMARY KEY (id)
);
