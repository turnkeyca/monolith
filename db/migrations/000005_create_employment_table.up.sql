create table employment (
    id uuid primary key,
    user_id uuid, 
	employer varchar(32),
	occupation varchar(32),
	duration varchar(16),
	additional_details varchar(512),
    annual_salary numeric(8, 2),
	rent_affordability varchar(512),
	last_updated varchar(35),
	created_on varchar(35),
    constraint FK_user_id foreign key (user_id) references users(id)
);