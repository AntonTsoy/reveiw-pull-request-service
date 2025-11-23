package models

type User struct {
	
}

/*
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    team_id INTEGER REFERENCES teams (id) ON DELETE SET NULL,
    is_active BOOLEAN NOT NULL DEFAULT true
);
*/