-- name: ListCurrencies :many
SELECT * FROM currencies ORDER BY code;

-- name: ListRegions :many
SELECT * FROM regions ORDER BY name;