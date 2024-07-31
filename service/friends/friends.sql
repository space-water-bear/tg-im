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

 Date: 15/07/2024 17:11:21
*/


-- ----------------------------
-- Table structure for friends
-- ----------------------------
DROP TABLE IF EXISTS "public"."friends";
CREATE TABLE "public"."friends" (
  "id" int4 NOT NULL DEFAULT nextval('friends_id_seq'::regclass),
  "user_id" int4 NOT NULL,
  "friend_id" int4 NOT NULL,
  "status" int4 NOT NULL,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "friend_nickname" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "friend_avatar" varchar(255) COLLATE "pg_catalog"."default",
  "friend_custom_name" varchar(255) COLLATE "pg_catalog"."default",
  "remark" varchar(255) COLLATE "pg_catalog"."default",
  "request_status" bool,
  "friend_username" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "friend_email" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."friends" OWNER TO "postgres";
COMMENT ON COLUMN "public"."friends"."status" IS '1.同意，2.等待，3.拒绝，4.黑名单';
COMMENT ON COLUMN "public"."friends"."request_status" IS '是否请求';

-- ----------------------------
-- Indexes structure for table friends
-- ----------------------------
CREATE INDEX "idx_friend_id" ON "public"."friends" USING btree (
  "friend_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "idx_status" ON "public"."friends" USING btree (
  "status" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "idx_user_id" ON "public"."friends" USING btree (
  "user_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table friends
-- ----------------------------
ALTER TABLE "public"."friends" ADD CONSTRAINT "unique_friend" UNIQUE ("user_id", "friend_id");

-- ----------------------------
-- Primary Key structure for table friends
-- ----------------------------
ALTER TABLE "public"."friends" ADD CONSTRAINT "friends_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table friends
-- ----------------------------
ALTER TABLE "public"."friends" ADD CONSTRAINT "fk_friend" FOREIGN KEY ("friend_id") REFERENCES "public"."users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "public"."friends" ADD CONSTRAINT "fk_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
