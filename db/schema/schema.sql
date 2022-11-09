create table products (
    id            text                     not null primary key,
    title         text                     not null,
    short_title   text                     not null,
    product_type  text                     not null,
    series_number numeric(5,2)             not null,
    created_at    timestamp with time zone not null default current_timestamp,
    updated_at    timestamp with time zone not null default current_timestamp
);

create table original_songs (
    id            text                     not null primary key,
    product_id    text                     not null,
    title         text                     not null,
    composer      text                     not null default '',
    arranger      text                     not null default '',
    track_number  integer                  not null,
    is_duplicate  boolean                  not null default false,
    created_at    timestamp with time zone not null default current_timestamp,
    updated_at    timestamp with time zone not null default current_timestamp
);
create index original_songs_product_id on original_songs (product_id);
