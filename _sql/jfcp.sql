
CREATE DATABASE paka
  WITH OWNER = pguser1
       ENCODING = 'UTF8'
       COLUMNSPACE = pg_default
       LC_COLLATE = 'en_US.UTF-8'
       LC_CTYPE = 'en_US.UTF-8'
       CONNECTION LIMIT = -1;

CREATE TABLE tbl_jfcp_angle_param (
id serial PRIMARY KEY NOT NULL,
device_serial varchar(12) NOT NULL,
park_number varchar(32) NOT NULL,
map_block_id varchar(32) NOT NULL,
bid integer NOT NULL,
bangle integer NOT NULL,
nid integer NOT NULL,
nangle integer NOT NULL,
crop_x integer NOT NULL,
crop_y integer NOT NULL,
crop_w integer NOT NULL,
crop_h integer NOT NULL,
remark varchar(255)
);

COMMENT ON COLUMN tbl_jfcp_angle_param.device_serial IS '设备串号';
COMMENT ON COLUMN tbl_jfcp_angle_param.park_number IS '停车位编号';
COMMENT ON COLUMN tbl_jfcp_angle_param.map_block_id IS '地图区块编号';
COMMENT ON COLUMN tbl_jfcp_angle_param.bid IS '底座角度序号';
COMMENT ON COLUMN tbl_jfcp_angle_param.bangle IS '底座角度';
COMMENT ON COLUMN tbl_jfcp_angle_param.nid IS '头部角度序号';
COMMENT ON COLUMN tbl_jfcp_angle_param.nangle IS '头部角度';
COMMENT ON COLUMN tbl_jfcp_angle_param.crop_x IS '剪裁区域矩形x坐标';
COMMENT ON COLUMN tbl_jfcp_angle_param.crop_y IS '剪裁区域矩形y坐标';
COMMENT ON COLUMN tbl_jfcp_angle_param.crop_w IS '剪裁区域矩形宽度';
COMMENT ON COLUMN tbl_jfcp_angle_param.crop_h IS '剪裁区域矩形高度';
COMMENT ON COLUMN tbl_jfcp_angle_param.remark IS '备注';

