create table employment (
    id uuid primary key,
    user_id uuid not null, 
	employer varchar(64) not null,
	occupation varchar(64) not null,
	duration varchar(64) not null,
	additional_details varchar(256),
    annual_salary numeric(10, 9),
    constraint FK_user_id foreign key (user_id) references users(id)
);