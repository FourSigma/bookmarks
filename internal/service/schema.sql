
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO web;
CREATE TABLE bookmarks (
   id  uuid PRIMARY KEY,
   url text UNIQUE NOT NULL,
   data jsonb NOT NULL
);
