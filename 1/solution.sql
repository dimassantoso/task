create table "user"
(
    ID serial constraint user_pk primary key,
    "UserName" varchar(32),
    "Parent" int4 not null
);

INSERT INTO "Users"("UserName", "Parent") VALUES ('Ali', 2);
INSERT INTO "Users"("UserName", "Parent") VALUES ('Budi', 0);
INSERT INTO "Users"("UserName", "Parent") VALUES ('Cecep', 1);

SELECT a.ID as "ID", a."UserName", (SELECT "UserName" FROM "Users" WHERE "Users".ID = a."Parent") as "ParentUserName" FROM "Users" a