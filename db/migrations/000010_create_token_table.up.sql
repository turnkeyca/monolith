create table token (
    id uuid primary key,
    token varchar(256),
    created_on varchar(35),
    status varchar(8)
);