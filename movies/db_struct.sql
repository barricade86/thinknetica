drop table if exists movies;
drop table if exists actors;
drop table if exists directors;
drop table if exists studio;
drop table if exists movies_and_studios;
drop table if exists movies_actors_and_directors;
create table if not exists movies_and_studios(movie_id int not null,studio_id int not null);
create unique index movies_and_studios_unq on movies_and_studios (movie_id,studio_id);
create table movies_actors_and_directors(
    movie_id int not null,
    director_id int not null,
    actor_id int not null
);
create unique index movies_actors_and_directors_unq on movies_actors_and_directors (
    movie_id,director_id,actor_id
);
CREATE TYPE rating AS ENUM ('PG-10', 'PG-13', 'PG-18');
create table if not exists movies(
    id serial primary key, 
    title TEXT, 
    year smallint not null, 
    box_office int not null, 
    ratings rating
);
create table if not exists actors(
    id serial primary key, 
    name TEXT not null, 
    birth_date date not null
);
create table if not exists directors(
    id serial primary key, 
    name TEXT not null, 
    birth_date date not null
);
create table if not exists studio(
    id serial primary key, 
    title TEXT
);
----------------------------------------------------------------------------
insert into movies(id, title,year,box_office,ratings) 
values(1,'Nightmare on the Elm street',1984,25504513,'PG-18');
insert into movies(id, title,year,box_office,ratings) 
values(2,'Freddy is dead:The final Nightmare',1991,34872033,'PG-18');
insert into movies(id, title,year,box_office,ratings) 
values(3,'Wishmaster',1997,15738769,'PG-18');
insert into movies(id, title,year,box_office,ratings) 
values(4,'Wishmaster2: Evil never dies',1999,15000769,'PG-18');
insert into movies(id, title,year,box_office,ratings) 
values(5,'DragonHeart',1996,51367375,'PG-18');
insert into movies(id, title,year,box_office,ratings) 
values(6,'A Nightmare on Elm Street Part 2: Freddy''s Revenge',1985,29999213,'PG-18');
--------------------------------------------------------------------------------------------------
insert into actors(id,name,birth_date) values(1,'Robert Englund','06-06-1947');
insert into actors(id,name,birth_date) values(2,'Andrew Divoff','07-02-1955');
insert into actors(id,name,birth_date) values(3,'Dennis Quaid','04-09-1954');
----------------------------------------------------------------------------------------------------
insert into directors(id,name,birth_date) values(1,'Wes Craven','08-02-1939');
insert into directors(id,name,birth_date) values(2,'Robert Kurtzman','11-25-1964');
insert into directors(id,name,birth_date) values(3,'Jack Sholder','06-08-1945');
insert into directors(id,name,birth_date) values(4,'Rob Cohen','03-12-1949');
------------------------------------------------------------------------------------------
insert into studio(id,title) values(1,'New Line Cinema');
insert into studio(id,title) values(2,'Universal');
insert into studio(id,title) values(3,'Live Entertaiment');
-----------------------------------------------------------------------------------------
insert into movies_and_studios(movie_id,studio_id) values(1,1);
insert into movies_and_studios(movie_id,studio_id) values(2,1);
insert into movies_and_studios(movie_id,studio_id) values(6,1);
------------------------------------------------------------------------------------------
insert into movies_actors_and_directors(movie_id,director_id,actor_id) values(1,1,1);
insert into movies_actors_and_directors(movie_id,director_id,actor_id) values(3,2,1);
insert into movies_actors_and_directors(movie_id,director_id,actor_id) values(3,2,2);
insert into movies_actors_and_directors(movie_id,director_id,actor_id) values(5,4,3);
insert into movies_actors_and_directors(movie_id,director_id,actor_id) values(6,3,1);
insert into movies_actors_and_directors(movie_id,director_id,actor_id) values(4,3,2);
-------------------------------------------------------------------------------------------
select m.id as movieId,m.title as movieName,a.name as actorName 
from movies m inner join movies_actors_and_directors mad on m.id=mad.movie_id 
              inner join actors a on mad.actor_id=a.id WHERE a.name='Robert Englund';
select * from movies m where m.box_office>1000;                       
select m.id,m.title,d.name 
from movies m inner join movies_actors_and_directors mad on m.id=mad.movie_id 
              inner join directors d on mad.director_id=d.id WHERE d.name='Jack Sholder';