create table pet (
    id uuid primary key,
    user_id uuid not null, 
    breed varchar(64),
    weight numeric(10, 9),
    constraint FK_user_id foreign key (user_id) references users(id)
);