-- +goose Up
CREATE TABLE IF NOT EXISTS lessons(
        id TEXT NOT NULL UNIQUE PRIMARY KEY, -- Pattern: ls-#
        topic TEXT NOT NULL UNIQUE,
        started_date DATE,
        r_one_date DATE,
        r_two_date DATE,
        r_three_date DATE,
        r_one_state TEXT, -- Acepted options: not_started, started, done
        r_two_state TEXT,  -- Acepted options: not_started, started, done
        r_three_state TEXT, -- Acepted options: not_started, started, done
        created_at TEXT NOT NULL DEFAULT (datetime('now', 'localtime')),
        updated_at TEXT NOT NULL DEFAULT (datetime('now', 'localtime'))
);

INSERT INTO lessons(id, topic) VALUES ('ls-0', 'Ejemplo');
-- +goose Down
DELETE FROM lessons;
DROP TABLE lessons;
