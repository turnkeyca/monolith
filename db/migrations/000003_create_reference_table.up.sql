create table reference (
    id uuid primary key,
    user_id uuid not null, 
    full_name varchar(64) not null,
    email varchar(64),
	phone_number varchar(64),
    relationship varchar(64),
    additional_details varchar(256),
    constraint FK_user_id foreign key (user_id) references users(id)
);