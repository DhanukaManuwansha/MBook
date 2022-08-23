CREATE TABLE "Patient" (
  "patient_id" varchar(64) PRIMARY KEY,
  "nic" varchar(20) UNIQUE NOT NULL,
  "first_name" varchar(50) NOT NULL,
  "middle_name" varchar(50) NOT NULL,
  "last_name" varchar(50) NOT NULL,
  "dob" date NOT NULL,
  "address_line_one" varchar(100),
  "address_line_two" varchar(100),
  "city" varchar(20),
  "state" varchar(20),
  "country" varchar(20) NOT NULL,
  "tell_no" varchar(20) NOT NULL,
  "email" varchar(70) UNIQUE NOT NULL,
  "password" varchar(50) NOT NULL,
  "sex" varchar(10),
  "height" float ,
  "weight" float ,
  "blood_group" varchar(10),
  "marital_state" varchar(10),
  "occupation" varchar(50),
  "nationality" varchar(30),
  "family_id" bigserial

);


CREATE TABLE "Family" (
  "family_id" bigserial PRIMARY KEY,
  "family_name" varchar(50) NOT NULL
);


ALTER TABLE "Patient" ADD FOREIGN KEY ("family_id") REFERENCES "Family" ("family_id");