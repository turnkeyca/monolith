create table reference (
    id uuid primary key,
    user_id uuid, 
    full_name varchar(64),
    email varchar(64),
	phone_number varchar(16),
    relationship varchar(32),
    additional_details varchar(512),
    last_updated varchar(35),
	created_on varchar(35),
    constraint FK_user_id foreign key (user_id) references users(id)
);