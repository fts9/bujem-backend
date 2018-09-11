CREATE SEQUENCE user_id_seq INCREMENT BY 3;
CREATE SEQUENCE note_id_seq INCREMENT BY 3;

CREATE TABLE users (
    id          bigint primary key DEFAULT NEXTVAL('user_id_seq'),
    username    varchar(50),
    email       varchar(255),
    password    varchar,
    created     timestamp without time zone,
    modified    timestamp without time zone
);
CREATE TABLE notes (
    id bigint primary key DEFAULT NEXTVAL('note_id_seq'),
    note_type varchar(20),
    note_content varchar,
    created timestamp without time zone,
    modified timestamp without time zone,
    owner_user_id bigint REFERENCES users(id)
);

ALTER SEQUENCE user_id_seq OWNED BY users.id;
ALTER SEQUENCE note_id_seq OWNED BY notes.id;
