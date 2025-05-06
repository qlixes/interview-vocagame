drop table logs;
create table logs (
	id serial primary key,
	car_id varchar(64) not null,
	ticket_id varchar(128) not null,
	park_id bigint not null,
	duration int default 0,
	amount decimal default 0.0
);

drop table parks;
create table parks(
    id serial primary key,
    car_id varchar(64) not null,
    ticket_id varchar(64) not null
);
