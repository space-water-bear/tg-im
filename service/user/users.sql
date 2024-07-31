/*
 Navicat Premium Data Transfer

 Source Server         : pg-10.3.21.250
 Source Server Type    : PostgreSQL
 Source Server Version : 150005 (150005)
 Source Host           : 10.3.21.250:5432
 Source Catalog        : tg-im
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 150005 (150005)
 File Encoding         : 65001

 Date: 15/07/2024 17:11:11
*/


-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "username" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "nickname" varchar(255) COLLATE "pg_catalog"."default",
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "ip" varchar(255) COLLATE "pg_catalog"."default",
  "last_login" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "avatar" varchar(255) COLLATE "pg_catalog"."default",
  "status" int4
)
;
ALTER TABLE "public"."users" OWNER TO "postgres";

-- ----------------------------
-- Indexes structure for table users
-- ----------------------------
CREATE INDEX "idx_email" ON "public"."users" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_username" ON "public"."users" USING btree (
  "username" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Triggers structure for table users
-- ----------------------------
CREATE TRIGGER "update_users_updated_at" BEFORE UPDATE ON "public"."users"
FOR EACH ROW
EXECUTE PROCEDURE "public"."update_updated_at_column"();

-- ----------------------------
-- Uniques structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_username_key" UNIQUE ("username");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
