CREATE TABLE contacts (
	id serial4 NOT NULL,
	contact_id int4 NULL,
	first_name varchar(128) NULL,
	last_name varchar(64) NULL,
	phone varchar(64) NULL,
	email varchar(64) NULL,
	CONSTRAINT contacts_contact_id_key UNIQUE (contact_id),
	CONSTRAINT contacts_pkey PRIMARY KEY (id)
);

-- DROP TABLE contacts;


CREATE TABLE tasks (
	id serial4 NOT NULL,
	task_id int4 NULL,
	title varchar(128) NULL,
	status varchar(64) NULL,
	priority varchar(64) NULL,
	created_at timestamp NULL,
	created_by varchar(64) NULL,
	due_date date NULL,
	CONSTRAINT tasks_pkey PRIMARY KEY (id),
	CONSTRAINT tasks_task_id_key UNIQUE (task_id)
);

-- DROP TABLE tasks;