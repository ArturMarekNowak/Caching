CREATE KEYSPACE public WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE public.users (
	Id uuid PRIMARY KEY,
	Name text, 
	Surname text, 
	Email text
);

INSERT INTO public.users (Id, Name, Surname, Email) VALUES (4fe271c7-ad29-4667-92ab-df8752a59618, 'Arthur', 'Morgan', 'ArthurMorgan@mail.com');