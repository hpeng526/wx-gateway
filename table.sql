CREATE TABLE T_USER
(id integer not null primary key, wx_id text, template_id text, create_time datetime);

CREATE TABLE T_USER_GROUP
(id INTEGER NOT NULL, user_id INTEGER NOT NULL, owner INTEGER NOT NULL DEFAULT 0);