-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS profiles
(
    id UUID PRIMARY KEY,
    name     VARCHAR(60) NOT NULL,
    last_name     VARCHAR(60) NOT NULL,
    sex     VARCHAR(60),
    date_of_birth TIMESTAMP,
    city VARCHAR(60),
    height INTEGER,
    chest_size INTEGER,
    waist_size INTEGER,
    hip_size INTEGER,
    weight INTEGER,
    avatar_url VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS photos
(
    id serial,
    photo_url VARCHAR(100) PRIMARY KEY,
    profile_id UUID NOT NULL,
    position integer,
    foreign key (profile_id) REFERENCES profiles(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE profiles;
DROP TABLE photos;
-- +goose StatementEnd
