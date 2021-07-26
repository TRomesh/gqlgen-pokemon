-- name: GetPokemon :one
SELECT * FROM pokemon
WHERE id = $1;

-- name: ListPokemons :many
SELECT * FROM pokemon
ORDER BY name;

-- name: CreatePokemon :one
INSERT INTO pokemon (name, power, description)
VALUES ($1, $2, $3)
    RETURNING *;

-- name: UpdatePokemon :one
UPDATE pokemon
SET name = $2, power = $3, description = $4
WHERE id = $1
    RETURNING *;

-- name: DeletePokemon :one
DELETE FROM pokemon
WHERE id = $1
    RETURNING *;

-- name: GetBattle :one
SELECT * FROM battle
WHERE id = $1;

-- name: ListBattles :many
SELECT * FROM battle
ORDER BY opponent;

-- name: CreateBattle :one
INSERT INTO battle (opponent, location, pokemon)
VALUES ($1, $2, $3)
    RETURNING *;

-- name: UpdateBattle :one
UPDATE battle
SET opponent = $2, location = $3
WHERE id = $1
    RETURNING *;

-- name: DeleteBattle :one
DELETE FROM battle
WHERE id = $1
    RETURNING *;