create table employment (
    id uuid primary key,
    user_id uuid, 
	employer varchar(64),
	occupation varchar(64),
	duration varchar(64),
	additional_details_employment varchar(256),
    annual_salary numeric(10, 2),
	rent_affordability varchar(256),
    constraint FK_user_id foreign key (user_id) references users(id)
);