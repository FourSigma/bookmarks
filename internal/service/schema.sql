
CREATE TABLE public.bookmarks (
   id  uuid PRIMARY KEY,
   url text UNIQUE NOT NULL,
   data jsonb NOT NULL
);

GRANT ALL PRIVILEGES ON TABLE bookmarks TO web;
