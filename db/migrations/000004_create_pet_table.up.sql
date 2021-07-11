create table pet (
    id uuid primary key,
    user_id uuid, 
    pet_type varchar(64),
	breed varchar(64),
    size_type varchar(64),
    constraint FK_user_id foreign key (user_id) references users(id)
);