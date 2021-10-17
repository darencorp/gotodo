create table if not exists todos(
    id serial not null,
    title varchar,
    description varchar,
    status varchar,
    created_at timestamp,
    deleted_at timestamp,
    primary key (id)
)