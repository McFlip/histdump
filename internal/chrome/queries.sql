-- name: GetVisits :many
SELECT visits.visit_time, urls.last_visit_time, urls.url, urls.title, urls.visit_count
FROM urls, visits
WHERE visits.url = urls.id
ORDER BY visit_time DESC;
