CREATE TABLE
    IF NOT EXISTS todos (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        todo text,
        done boolean,
        createdDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updateDate TIMESTAMP,
        active boolean
    );