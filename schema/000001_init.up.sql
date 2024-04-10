CREATE TABLE IF NOT EXISTS users (
                                     id              SERIAL PRIMARY KEY,
                                     first_name      VARCHAR NOT NULL,
                                     last_name       VARCHAR NOT NULL,
                                     username        VARCHAR NOT NULL UNIQUE,
                                     password_hash   VARCHAR NOT NULL,
                                     created_at      TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS goals (
                                     id                SERIAL PRIMARY KEY,
                                     name              VARCHAR NOT NULL,
                                     description       VARCHAR NOT NULL,
                                     completion_status VARCHAR NOT NULL DEFAULT 'Progress',
                                     start_date        TIMESTAMP NOT NULL,
                                     end_date          TIMESTAMP NOT NULL,
                                     created_at        TIMESTAMP NOT NULL DEFAULT now(),
                                     updated_at        TIMESTAMP NOT NULL DEFAULT now(),
                                     user_id           INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
                                     UNIQUE(user_id, name)
);

CREATE TABLE IF NOT EXISTS tasks (
                                     id          SERIAL PRIMARY KEY,
                                     name        VARCHAR NOT NULL,
                                     description VARCHAR NOT NULL,
                                     frequency   VARCHAR NOT NULL,
                                     created_at  TIMESTAMP NOT NULL DEFAULT now(),
                                     updated_at  TIMESTAMP NOT NULL DEFAULT now(),
                                     goal_id     INT REFERENCES goals(id) ON DELETE CASCADE NOT NULL,
                                     UNIQUE (goal_id, name)
);