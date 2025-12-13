-- +goose Up
CREATE TABLE IF NOT EXISTS lessons(
        id TEXT NOT NULL UNIQUE PRIMARY KEY, -- Pattern: ls-#
        topic TEXT NOT NULL UNIQUE,
        started_date TEXT,
        amount_of_cards INTEGER NOT NULL,
        created_at TEXT NOT NULL DEFAULT (datetime('now', 'localtime')),
        updated_at TEXT NOT NULL DEFAULT (datetime('now', 'localtime'))
);

INSERT INTO lessons(id, topic, amount_of_cards) VALUES ('ls-0', 'Ejemplo', 0);
-- +goose Down
DELETE FROM lessons;
DROP TABLE lessons;
