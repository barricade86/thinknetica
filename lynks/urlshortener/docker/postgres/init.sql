CREATE TABLE IF NOT EXISTS links(
    short_link TEXT NOT NULL, 
    original_link TEXT NOT NULL,
    CONSTRAINT short_link_pkey PRIMARY KEY (short_link)
 );