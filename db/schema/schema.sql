create type product_type as enum (
    'pc98',
    'windows',
    'zuns_music_collection',
    'akyus_untouched_score',
    'commercial_books',
    'other'
);

create table products (
    id            text                     not null primary key,
    name          text                     not null,
    short_name    text                     not null,
    product_type  product_type             not null,
    series_number numeric(5,2)             not null,
    created_at    timestamp with time zone not null default current_timestamp,
    updated_at    timestamp with time zone not null default current_timestamp
);

create table original_songs (
    id            text                     not null primary key,
    product_id    text                     not null,
    name          text                     not null,
    composer      text                     not null default '',
    arranger      text                     not null default '',
    track_number  integer                  not null,
    is_duplicate  boolean                  not null default false,
    created_at    timestamp with time zone not null default current_timestamp,
    updated_at    timestamp with time zone not null default current_timestamp
);
create index original_songs_product_id on original_songs (product_id);
