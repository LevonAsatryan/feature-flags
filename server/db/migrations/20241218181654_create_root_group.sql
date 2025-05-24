-- +goose Up
-- +goose StatementBegin
INSERT INTO
    groups (name)
VALUES
    ('root');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM groups
WHERE
    name = 'root';

-- +goose StatementEnd
