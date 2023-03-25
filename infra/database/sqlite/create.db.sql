CREATE TABLE IF NOT EXISTS vocations (
    id INTEGER PRIMARY KEY, 
    name VARCHAR NOT NULL, 
    createdAt DATETIME NOT NULL DEFAULT NOW(), 
    updatedAt DATETIME NOT NULL DEFAULT NOW(), 
    deletedAt DATETIME NULL 
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY, 
    nickname VARCHAR NOT NULL UNIQUE, 
    vocationId INTEGER NULL, 
    createdAt DATETIME NOT NULL DEFAULT NOW(), 
    updatedAt DATETIME NOT NULL DEFAULT NOW(), 
    deletedAt DATETIME NULL, 
    FOREIGN KEY (vocationId) REFERENCES vocations (id) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS backpacks (
    id INTEGER PRIMARY KEY, 
    userId INTEGER NOT NULL, 
    FOREIGN KEY (userId) REFERENCES users (id) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS items (
    id INTEGER PRIMARY KEY, 
    name INTEGER NOT NULL, 
    backpackId INTEGER NULL, 
    FOREIGN KEY (backpackId) REFERENCES backpacks (id) ON DELETE CASCADE 
);

IF (SELECT COUNT(*) FROM vocations LIMIT 1) === 0 THEN
    INSERT INTO vocations (name) VALUES ('Knight'), ('Sorcerer'), ('Paladin');
END
