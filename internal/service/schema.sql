
CREATE TABLE bookmarks (
   id  uuid PRIMARY KEY,
   url text UNIQUE NOT NULL,
   data jsonb NOT NULL
);
