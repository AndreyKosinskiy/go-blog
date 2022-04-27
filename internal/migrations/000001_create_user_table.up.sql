CREATE TABLE IF NOT EXISTS users(
    id UUID DEFAULT gen_random_uuid(),
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    is_delete BOOLEAN DEFAULT FALSE,
    create_at TIMESTAMP NOT NULL DEFAULT NOW(),
    update_at TIMESTAMP,
    delete_at TIMESTAMP
)