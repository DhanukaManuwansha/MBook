CREATE TABLE "User" (
  "user_id" varchar(64) PRIMARY KEY,
  "user_name" varchar(40) UNIQUE NOT NULL,
  "first_name" varchar(40) NOT NULL,
  "last_name" varchar(40) NOT NULL,
  "nic" varchar(20) UNIQUE NOT NULL,
  "tell_no" varchar(20) NOT NULL,
  "address" varchar(100),
  "user_email" varchar(40) UNIQUE NOT NULL,
  "user_pwd" varchar(64) NOT NULL,
  "is_email_verified" boolean NOT NULL,
  "is_tell_no_verified" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now())
);

CREATE TABLE "Doctor" (
  "doctor_id" bigserial PRIMARY KEY,
  "reg_number" varchar(20) NOT NULL,
  "dob" date NOT NULL,
  "user_id" varchar(20) NOT NULL
);

CREATE TABLE "Specialization" (
  "specialization_id" bigserial PRIMARY KEY,
  "specialization_name" varchar(40) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now())
);

CREATE TABLE "Specialization_Doctor" (
  "specialization_id" bigserial REFERENCES "Specialization" ("specialization_id") ON UPDATE CASCADE ON DELETE CASCADE, 
  "doctor_id" bigserial REFERENCES "Doctor" ("doctor_id") ON UPDATE CASCADE,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  CONSTRAINT "specialization_doctor_id" PRIMARY KEY ("specialization_id", "doctor_id", "created_at")  -- explicit pk
);

CREATE TABLE "Nurse" (
  "nurse_id" bigserial PRIMARY KEY,
  "reg_number" varchar(20) NOT NULL,
  "dob" DATE NOT NULL,
  "nurse_rank" varchar(5) NOT NULL,
  "active_status" bool NOT NULL,
  "center_id" bigserial NOT NULL,
  "user_id" varchar(100) NOT NULL
);

CREATE TABLE "CenterAdmin" (
  "centerAdmin_id" bigserial PRIMARY KEY,
  "center_id" bigserial NOT NULL,
  "user_id" varchar(100) NOT NULL
);

CREATE TABLE "SuperAdmin" (
  "superAdmin_id" bigserial PRIMARY KEY,
  "user_id" varchar(100) NOT NULL
);

CREATE TABLE "Center" (
  "center_id" bigserial PRIMARY KEY,
  "center_name" varchar(50) UNIQUE NOT NULL,
  "address" varchar(100),
  "tell_no" varchar(20) NOT NULL,
  "email" varchar(40) NOT NULL,
  "website" varchar(20) UNIQUE NOT NULL,
  "active_status" bool NOT NULL,
  "online_status" bool NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "centertype_id" bigserial NOT NULL
);

CREATE TABLE "CenterType" (
  "centertype_id" bigserial PRIMARY KEY,
  "center_type" varchar(10) NOT NULL,
  "centertype_des" varchar(30),
  "created_at" timestamptz NOT NULL DEFAULT(now())
);

CREATE TABLE "SuperAdmin_Center" (
  "superadmin_id" bigserial REFERENCES "SuperAdmin" ("superAdmin_id") ON UPDATE CASCADE ON DELETE CASCADE, 
  "center_id" bigserial REFERENCES "Center" ("center_id") ON UPDATE CASCADE,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  CONSTRAINT "superAdmin_center_id" PRIMARY KEY ("superadmin_id", "center_id", "created_at")  -- explicit pk
);

CREATE TABLE "Ward" (
  "ward_id" bigserial PRIMARY KEY,
  "ward_no" int NOT NULL,
  "tot_of_patient" bigint NOT NULL,
  "ward_type" varchar(10) NOT NULL,
  "period" varchar(10) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "center_id" bigserial NOT NULL
);

CREATE TABLE "BedTicket" (
  "bedTicket_id" bigserial PRIMARY KEY,
  "bed_no" int NOT NULL,
  "assigneddate_time" date NOT NULL,
  "leavedate_time" date NOT NULL,
  "ticket_type" varchar(20) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "patient_id" varchar(64) NOT NULL,
  "ward_id" bigserial NOT NULL
);

CREATE TABLE "NurseNotes" (
  "nursingnotes_id" bigserial PRIMARY KEY,
  "remark" varchar(20) NOT NULL,
  "nursing_care" varchar(30),
  "notes_datetime" date,
  "created_at" timestamptz DEFAULT(now()) NOT NULL,
  "nurse_id" bigserial NOT NULL,
  "patient_id" varchar(64) NOT NULL
);

CREATE TABLE "Doctor_Center" (
  "doctor_id" bigserial REFERENCES "Doctor" ("doctor_id") ON UPDATE CASCADE ON DELETE CASCADE, 
  "center_id" bigserial REFERENCES "Center" ("center_id") ON UPDATE CASCADE,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  CONSTRAINT "doctor_center_id" PRIMARY KEY ("doctor_id", "center_id", "created_at")  -- explicit pk
);

CREATE TABLE "Patient_Center" (
  "patient_id" varchar(64) NOT NULL , 
  "center_id" bigserial REFERENCES "Center" ("center_id") ON UPDATE CASCADE,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  CONSTRAINT "patient_center_id" PRIMARY KEY ("patient_id", "center_id", "created_at")  -- explicit pk
);

CREATE TABLE "DrugOrder" (
  "drugorder_id" bigserial PRIMARY KEY,
  "dose" varchar(10) NOT NULL,
  "dosage" varchar(10) NOT NULL,
  "givendate" date NOT NULL,
  "giveuntil" date NOT NULL,
  "route" varchar(10) NOT NULL,
  "frequency" int NOT NULL,
  "dosaage_time" time NOT NULL,
  "drug_status" varchar(20),
  "drug_id" bigserial,
  "omit_status" int NOT NULL,
  "patient_id" varchar(64) NOT NULL,
  "prescription_id" bigserial NOT NULL,
  "session_id" bigserial,
  "doctor_id" bigserial NOT NULL
);

CREATE TABLE "Doctor_DrugOrder" (
  "doctor_id" bigserial REFERENCES "Doctor" ("doctor_id") ON UPDATE CASCADE ON DELETE CASCADE, 
  "drugorder_id" bigserial REFERENCES "DrugOrder" ("drugorder_id") ON UPDATE CASCADE,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  CONSTRAINT "doctor_drugorder_id" PRIMARY KEY ("doctor_id", "drugorder_id", "created_at")  -- explicit pk
);

CREATE TABLE "Prescription" (
  "prescription_id" bigserial PRIMARY KEY,
  "pres_date" date NOT NULL,
  "pres_time" time NOT NULL,
  "active_status" bool NOT NULL,
  "notes" varchar(20),
  "sketch" varchar(20),
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "patient_id" varchar(64) NOT NULL,
  "doctor_id" bigserial NOT NULL
);

CREATE TABLE "Drug" (
  "drug_id" bigserial PRIMARY KEY,
  "drug_name" varchar(20) NOT NULL,
  "scientific_name" varchar(20),
  "drug_category" varchar(10) NOT NULL,
  "storage_temperature" varchar(10),
  "dangerous_level" varchar(10) NOT NULL,
  "manufacture" varchar(20) NOT NULL,
  "no_of_units" int NOT NULL,
  "notes" varchar(30),
  "created_at" timestamptz NOT NULL DEFAULT(now())
);

CREATE TABLE "Observation" (
  "observation_id" bigserial PRIMARY KEY,
  "obs_datetime" timestamptz NOT NULL,
  "bp_sys" float NOT NULL,
  "bp_dia" float NOT NULL,
  "pr" float NOT NULL,
  "rr" float NOT NULL,
  "temp" float NOT NULL,
  "notes" varchar(20),
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "patient_id" varchar(64) NOT NULL,
  "nurse_id" bigserial NOT NULL

);
--BP - Sys/Dia

CREATE TABLE "TestReport" (
  "testreport_id" bigserial PRIMARY KEY,
  "report_name" varchar(15) NOT NULL,
  "report_duration" DATE NOT NULL,
  "doctor_name" varchar(15) NOT NULL,
  "report_description" varchar(20),
  "report" varchar(30) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "added_by_nurse" bool NOT Null,
  "testreporttype_id" bigserial NOT NULL,
  "patient_id" varchar(64) NOT NULL,
  "nurse_id" bigserial 
  
);

CREATE TABLE "TestReportType" (
  "testreporttype_id" bigserial PRIMARY KEY,
  "test_name" varchar(15) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now())
);

CREATE TABLE "Sessions" (
  "session_id" bigserial PRIMARY KEY,
  "session_date" date NOT NULL,
  "session_time" time NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "center_id" bigserial NOT NULL,
  "patient_id" varchar(64) NOT NULL,
  "doctor_id" bigserial NOT NULL
);

CREATE TABLE "DrugSummery" (
  "drugsummery_id" bigserial PRIMARY KEY,
  "summery_date" date NOT NULL,
  "summery_status" bool NOT NULL,
  "first_dose_is_given" int NOT NULL,
  "second_dose_is_given" int NOT NULL,
  "third_dose_is_given" int NOT NULL,
  "fourth_dose_is_given" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "drugorder_id" bigserial,
  "patient_id" varchar(64) NOT NULL
);

CREATE TABLE "CriticalFacts" (
  "criticalfact_id" bigserial PRIMARY KEY,
  "illness_id" bigserial NOT NULL,
  "patient_id" varchar NOT NULL,
  "doctor_id" bigserial NOT NULL
);

CREATE TABLE "Illness" (
  "illness_id" bigserial PRIMARY KEY,
  "illness_name" varchar(20) NOT NULL,
  "illness_information" varchar(40),
  "created_at" timestamptz NOT NULL DEFAULT(now())
);

ALTER TABLE "Doctor" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("user_id");

ALTER TABLE "Nurse" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("user_id");

ALTER TABLE "Nurse" ADD FOREIGN KEY ("center_id") REFERENCES "Center" ("center_id");

ALTER TABLE "CenterAdmin" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("user_id");

ALTER TABLE "CenterAdmin" ADD FOREIGN KEY ("center_id") REFERENCES "Center" ("center_id");

ALTER TABLE "SuperAdmin" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("user_id");

ALTER TABLE "Center" ADD FOREIGN KEY ("centertype_id") REFERENCES "CenterType" ("centertype_id");

ALTER TABLE "Ward" ADD FOREIGN KEY ("center_id") REFERENCES "Center" ("center_id");

ALTER TABLE "BedTicket" ADD FOREIGN KEY ("ward_id") REFERENCES "Ward" ("ward_id");

ALTER TABLE "NurseNotes" ADD FOREIGN KEY ("nurse_id") REFERENCES "Nurse" ("nurse_id");

ALTER TABLE "DrugOrder" ADD FOREIGN KEY ("prescription_id") REFERENCES "Prescription" ("prescription_id");

ALTER TABLE "DrugOrder" ADD FOREIGN KEY ("session_id") REFERENCES "Sessions" ("session_id");

ALTER TABLE "DrugOrder" ADD FOREIGN KEY ("drug_id") REFERENCES "Drug" ("drug_id");

ALTER TABLE "DrugOrder" ADD FOREIGN KEY ("doctor_id") REFERENCES "Doctor" ("doctor_id");

ALTER TABLE "Prescription" ADD FOREIGN KEY ("doctor_id") REFERENCES "Doctor" ("doctor_id");

ALTER TABLE "TestReport" ADD FOREIGN KEY ("testreporttype_id") REFERENCES "TestReportType" ("testreporttype_id");

ALTER TABLE "TestReport" ADD FOREIGN KEY ("nurse_id") REFERENCES "Nurse" ("nurse_id");

ALTER TABLE "Sessions" ADD FOREIGN KEY ("doctor_id") REFERENCES "Doctor" ("doctor_id");

ALTER TABLE "Sessions" ADD FOREIGN KEY ("center_id") REFERENCES "Center" ("center_id");

ALTER TABLE "CriticalFacts" ADD FOREIGN KEY ("illness_id") REFERENCES "Illness" ("illness_id");

ALTER TABLE "CriticalFacts" ADD FOREIGN KEY ("doctor_id") REFERENCES "Doctor" ("doctor_id");

ALTER TABLE "Observation" ADD FOREIGN KEY ("nurse_id") REFERENCES "Nurse" ("nurse_id");
