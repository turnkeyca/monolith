create table permission (
    id uuid primary key,
    user_id uuid,
    permission varchar(16),
    on_user_id uuid,
	last_updated varchar(35),
	created_on varchar(35),
    constraint FK_user_id foreign key (user_id) references users(id),
    constraint FK_on_user_id foreign key (on_user_id) references users(id)
);