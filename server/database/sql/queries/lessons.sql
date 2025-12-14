-- name: AddLesson :exec
INSERT INTO lessons (topic, amount_of_cards)
VALUES (?, ?);
