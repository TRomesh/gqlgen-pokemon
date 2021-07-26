CREATE TABLE IF NOT EXISTS pokemon (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    power  TEXT,
    description TEXT,
    date DATE
);

CREATE TABLE IF NOT EXISTS battle (
    id BIGSERIAL PRIMARY KEY,
    opponent TEXT ,
    location TEXT,
    date DATE,
    pokemon_id BIGINT NOT NULL,
    FOREIGN KEY (pokemon_id) REFERENCES pokemon(id) ON DELETE RESTRICT
);