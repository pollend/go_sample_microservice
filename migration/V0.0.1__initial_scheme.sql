create table pet_store
(
	id serial not null
		constraint pet_store_pk
			primary key,
	name varchar(80)
);

alter table pet_store owner to petstore_api;

create table pet
(
	id serial not null
		constraint pet_pk
			primary key,
	breed varchar(100),
	name varchar(80),
	pet_store_id integer
		constraint pet_pet_store_id_fk
			references pet_store
);

alter table pet owner to petstore_api;

create unique index pet_id_uindex
	on pet (id);

create unique index pet_store_id_uindex
	on pet_store (id);

