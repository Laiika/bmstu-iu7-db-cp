-- ТАБЛИЦЫ

create table if not exists analyzer_types
(
    name        text primary key,
    max_sensors int not null
);

create table if not exists gas_analyzers
(
    id                text primary key,
    type              text not null,
    part_number       text not null,
    job_number        text not null,
    software_version  text not null,

    foreign key (type) references analyzer_types(name) on delete cascade
);

create table if not exists gases
(
    name    text primary key,
    formula text,
    type    text not null
);

create table if not exists sensors
(
    id                text primary key,
    type              text not null,
    gas               text not null,
    analyzer_id       text,
    low_limit_alarm   text not null,
    upper_limit_alarm text not null,

    foreign key (gas) references gases(name) on delete cascade,
    foreign key (analyzer_id) references gas_analyzers(id) on delete set null
);

create table if not exists events
(
    id            int generated always as identity primary key,
    signal_time   timestamp not null,
    sensor_id     text not null,
    peak_readings real not null,

    foreign key (sensor_id) references sensors(id) on delete cascade
);

create table if not exists types_gases
(
    gas           text not null,
    analyzer_type text not null,

    primary key (analyzer_type, gas),
    foreign key (gas) references gases(name) on delete cascade,
    foreign key (analyzer_type) references analyzer_types(name) on delete cascade
);

create table if not exists users
(
    id       int generated always as identity primary key,
    name     text unique not null,
    password text not null,
    role     text not null
);

-- РОЛИ

-- Обычный пользователь
create role simple_user;
grant select on public.gas_analyzers to simple_user;
grant select on public.sensors to simple_user;
grant select on public.gases to simple_user;
grant select on public.events to simple_user;
grant select on public.analyzer_types to simple_user;
grant select on public.types_gases to simple_user;

create user user1 with PASSWORD 'user1' in role simple_user;

-- Сотрудник
create role employee inherit;
grant simple_user to employee;
grant insert, delete on public.gas_analyzers to employee;
grant insert, update, delete on public.sensors to employee;
grant insert on public.gases to employee;
grant insert, delete on public.events to employee;
grant insert on public.types_gases to employee;
grant insert, delete on public.analyzer_types to employee;

create user empl1 with PASSWORD 'empl1' in role employee;

-- Администратор
create role admin;
grant create, usage on schema public to admin;
grant all privileges on all tables in schema public to admin;

create user admin1 with PASSWORD 'admin1' in role admin;

-- ТРИГГЕР

create or replace function check_analyzer()
returns trigger
as $$
begin
    if new.analyzer_id is null then
        return new;
    end if;

    if (select count(*) from sensors where analyzer_id = new.analyzer_id) >=
    (select ats.max_sensors from analyzer_types ats join
    (select type from gas_analyzers where id = new.analyzer_id) ga on ga.type = ats.name) then
        raise exception 'gas analyzer with id % is already has max number of sensors', new.analyzer_id;
    end if;

    if new.gas not in
    (select gas from types_gases where analyzer_type in
    (select type from gas_analyzers where id = new.analyzer_id)) then
        raise exception 'gas analyzer cannot work with %', new.gas;
    end if;

    return new;
end;
$$ language plpgsql;

create or replace trigger check_analyzer_trigger
before insert or update on sensors
for each row execute function check_analyzer();
