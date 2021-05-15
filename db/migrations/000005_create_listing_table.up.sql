create table listing (
    id uuid primary key,
    user_id uuid not null, 
	name varchar(64) not null,
    address varchar(64) not null,
	link varchar(64) not null,
    constraint FK_user_id foreign key (user_id) references users(id)
);