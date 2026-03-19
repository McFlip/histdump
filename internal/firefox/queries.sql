-- name: GetVisits :many
SELECT moz_historyvisits.visit_date, moz_places.last_visit_date, moz_places.url, moz_places.title, moz_places.visit_count
FROM moz_historyvisits, moz_places
WHERE moz_historyvisits.place_id = moz_places.id
ORDER BY moz_historyvisits.visit_date DESC;