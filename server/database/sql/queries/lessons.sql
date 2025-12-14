-- name: AddLesson :exec
INSERT INTO lessons (topic, started_date, repetitions_dates, amount_of_cards)
VALUES (?, ?, ?, ?);
