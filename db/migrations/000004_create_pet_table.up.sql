create table pet (
    id uuid primary key,
    user_id uuid, 
    pet_type varchar(8),
	breed varchar(16),
    size_type varchar(8),
    last_updated varchar(35),
	created_on varchar(35),
    constraint FK_user_id foreign key (user_id) references users(id)
);