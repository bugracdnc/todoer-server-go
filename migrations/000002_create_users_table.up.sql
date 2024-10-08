CREATE TABLE
    IF NOT EXISTS users (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        token text,
        name text,
        active boolean
    );

ALTER TABLE todos
ADD COLUMN user_id uuid;