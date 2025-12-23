-- +goose Up
CREATE TABLE IF NOT EXISTS lessons (
id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
topic TEXT NOT NULL UNIQUE,
started_date TEXT NOT NULL,
repetitions_dates TEXT NOT NULL,
amount_of_cards INTEGER NOT NULL,
created_at TEXT NOT NULL DEFAULT (datetime ('now', 'localtime')),
updated_at TEXT NOT NULL DEFAULT (datetime ('now', 'localtime'))
) ;

INSERT INTO lessons (topic, started_date, repetitions_dates, amount_of_cards)
VALUES ('Ejemplo', '0000-00-00', json ('[]'), 0) ;
-- +goose Down
DELETE FROM lessons ;
DROP TABLE lessons ;
