CREATE KEYSPACE public WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE public.users (
	Id int PRIMARY KEY,
	Name text, 
	Surname text, 
	Email text
);

INSERT INTO public.users (Id, Name, Surname, Email) VALUES (1, 'Arthur', 'Morgan', 'ArthurMorgan@mail.com');