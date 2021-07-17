create table roommate (
    id uuid primary key,
    user_id uuid, 
    full_name varchar(64),
    email varchar(64),
    constraint FK_user_id foreign key (user_id) references users(id)
);