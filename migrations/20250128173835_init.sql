-- +goose Up
-- +goose StatementBegin
CREATE TABLE news (
  "Id" SERIAL PRIMARY KEY,
  "Title" TEXT NOT NULL,
  "Content" TEXT NOT NULL
);

CREATE TABLE news_categories (
  "NewsId" BIGINT NOT NULL,
  "CategoryId" BIGINT NOT NULL,
  PRIMARY KEY ("NewsId", "CategoryId")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS news_categories;
DROP TABLE IF EXISTS news;
-- +goose StatementEnd
