-- angle params

DROP TABLE IF EXISTS tbl_jfcp_angle_param;
CREATE TABLE tbl_jfcp_angle_param (
id serial PRIMARY KEY NOT NULL,
device_serial varchar(12) NOT NULL,
ref_map_block_id varchar(32) NOT NULL,
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
COMMENT ON COLUMN tbl_jfcp_angle_param.ref_map_block_id IS '地图区块编号';
COMMENT ON COLUMN tbl_jfcp_angle_param.bid IS '底座角度序号';
COMMENT ON COLUMN tbl_jfcp_angle_param.bangle IS '底座角度';
COMMENT ON COLUMN tbl_jfcp_angle_param.nid IS '头部角度序号';
COMMENT ON COLUMN tbl_jfcp_angle_param.nangle IS '头部角度';
COMMENT ON COLUMN tbl_jfcp_angle_param.crop_x IS '剪裁区域矩形x坐标';
COMMENT ON COLUMN tbl_jfcp_angle_param.crop_y IS '剪裁区域矩形y坐标';
COMMENT ON COLUMN tbl_jfcp_angle_param.crop_w IS '剪裁区域矩形宽度';
COMMENT ON COLUMN tbl_jfcp_angle_param.crop_h IS '剪裁区域矩形高度';
COMMENT ON COLUMN tbl_jfcp_angle_param.remark IS '备注';

-- plate result

DROP TABLE IF EXISTS tbl_jfcp_plate_result;
CREATE TABLE tbl_jfcp_plate_result (
id serial PRIMARY KEY NOT NULL,
ref_floor_id integer NOT NULL, -- 所在楼层编号
ref_map_block_id varchar(32) NOT NULL, -- 对应地图中的块编号
car_status smallint NOT NULL default 0, -- 车辆状态，0:无车 1:有车 2:未知
plate_provice_code smallint, -- 车牌省份数字代码，
plate_provice_char varchar(8), -- 车牌省份汉字，例如粤B中的粤字
plate_city_code varchar(8), -- 车牌城市字母代码，例如粤B中的B
plate_number varchar(8), -- 5位的车牌号码
plate_literal varchar(12), -- 车牌字面值 例如 粤B12345
img_plate varchar(128), -- 车牌图片URL
img_crop varchar(128), -- 区域截图URL
update_time timestamp with time zone DEFAULT now(),
remark varchar(255)
);

COMMENT ON COLUMN tbl_jfcp_plate_result.ref_floor_id IS '引用楼层ID';
COMMENT ON COLUMN tbl_jfcp_plate_result.ref_map_block_id IS '引用地图区块ID';
COMMENT ON COLUMN tbl_jfcp_plate_result.car_status IS '车辆状态';
COMMENT ON COLUMN tbl_jfcp_plate_result.plate_provice_code IS '车牌省份数字代码';
COMMENT ON COLUMN tbl_jfcp_plate_result.plate_provice_char IS '车牌省份汉字';
COMMENT ON COLUMN tbl_jfcp_plate_result.plate_city_code IS '车牌城市字母代码';
COMMENT ON COLUMN tbl_jfcp_plate_result.plate_number IS '5位的车牌号码';
COMMENT ON COLUMN tbl_jfcp_plate_result.plate_literal IS '车牌字面值';
COMMENT ON COLUMN tbl_jfcp_plate_result.img_plate IS '车牌图片URL';
COMMENT ON COLUMN tbl_jfcp_plate_result.img_crop IS '区域截图URL';
COMMENT ON COLUMN tbl_jfcp_plate_result.update_time IS '更新时间';


-- plate result history

DROP TABLE IF EXISTS tbl_jfcp_plate_result_history;
CREATE TABLE tbl_jfcp_plate_result_history (
id serial PRIMARY KEY NOT NULL,
ref_floor_id integer NOT NULL, -- 所在楼层编号
ref_map_block_id varchar(32) NOT NULL, -- 对应地图中的块编号
ref_real_block_name varchar(64), -- 实际场景的名称
ref_device_serial varchar(24), -- 设备编号
ref_bid integer, -- bid编号
ref_nid integer, -- nid编号
update_time timestamp with time zone DEFAULT now(), -- 数据更新时间
car_status smallint NOT NULL default 0, -- 车辆状态，0:无车 1:有车 2:未知
plate_provice_code smallint, -- 车牌省份数字代码，
plate_provice_char varchar(8), -- 车牌省份汉字，例如粤B中的粤字
plate_city_code varchar(8), -- 车牌城市字母代码，例如粤B中的B
plate_number varchar(8), -- 5位的车牌号码
plate_literal varchar(12), -- 车牌字面值 例如 粤B12345
img_plate varchar(128), -- 历史车牌图片URL
img_crop varchar(128), -- 历史区域截图URL
remark varchar(255)
);

COMMENT ON COLUMN tbl_jfcp_plate_result_history.ref_floor_id IS '引用楼层ID';
COMMENT ON COLUMN tbl_jfcp_plate_result_history.ref_map_block_id IS '引用地图区块ID';
COMMENT ON COLUMN tbl_jfcp_plate_result_history.ref_real_block_name IS '实际名称';
COMMENT ON COLUMN tbl_jfcp_plate_result_history.update_time IS '数据更新时间';
COMMENT ON COLUMN tbl_jfcp_plate_result_history.car_status IS '车辆状态';
COMMENT ON COLUMN tbl_jfcp_plate_result_history.plate_provice_code IS '车牌省份数字代码';
COMMENT ON COLUMN tbl_jfcp_plate_result_history.plate_provice_char IS '车牌省份汉字';
COMMENT ON COLUMN tbl_jfcp_plate_result_history.plate_city_code IS '车牌城市字母代码';
COMMENT ON COLUMN tbl_jfcp_plate_result_history.plate_number IS '5位的车牌号码';
COMMENT ON COLUMN tbl_jfcp_plate_result_history.plate_literal IS '车牌字面值';
COMMENT ON COLUMN tbl_jfcp_plate_result_history.img_plate IS '车牌图片URL';
COMMENT ON COLUMN tbl_jfcp_plate_result_history.img_crop IS '区域截图URL';


-- function
CREATE OR REPLACE FUNCTION nf_save_plate_result(in_serial varchar(12), in_bid integer, in_nid integer, 
	in_car_status integer, in_plate_provice_code integer, in_plate_provice_char varchar(8), 
	in_plate_city_code varchar(8), in_plate_number varchar(8), in_plate_literal varchar(12), in_img_plate_unique varchar(128), in_img_plate_history varchar(128))
RETURNS void AS $$
DECLARE
	_map_block_id varchar;
	_floor_id integer;
	_block_name varchar;
	_exit_count integer;
BEGIN
	
	SELECT ref_map_block_id INTO STRICT _map_block_id FROM tbl_jfcp_angle_param WHERE device_serial = in_serial and bid = in_bid and nid = in_nid;
	SELECT ref_floor_id INTO STRICT _floor_id FROM tbl_jfc_device WHERE device_serial = in_serial;
	SELECT block_name INTO _block_name FROM tbl_map_realworld_element WHERE ref_floor_id = _floor_id and map_block_id = _map_block_id;

	SELECT COUNT(1) INTO _exit_count FROM tbl_jfcp_plate_result WHERE ref_floor_id = _floor_id  and ref_map_block_id = _map_block_id;
	IF _exit_count = 0 THEN
		INSERT INTO tbl_jfcp_plate_result (ref_floor_id, ref_map_block_id, car_status, plate_provice_code, plate_provice_char, plate_city_code, plate_number, plate_literal, img_plate) 
			VALUES (_floor_id, _map_block_id, in_car_status, in_plate_provice_code, in_plate_provice_char, in_plate_city_code, in_plate_number, in_plate_literal, in_img_plate_unique);
	ELSEIF _exit_count = 1 THEN
		UPDATE tbl_jfcp_plate_result SET car_status = in_car_status, plate_provice_code = in_plate_provice_code, plate_provice_char = in_plate_provice_char, 
		plate_city_code = in_plate_city_code, plate_number = in_plate_number, plate_literal = in_plate_literal, img_plate = in_img_plate_unique, update_time = now()
			WHERE ref_floor_id = _floor_id and ref_map_block_id = _map_block_id;	
	END IF;
	
	INSERT INTO tbl_jfcp_plate_result_history (ref_floor_id, ref_map_block_id, car_status, plate_provice_code, plate_provice_char, plate_city_code, plate_number, plate_literal, img_plate, ref_real_block_name, ref_device_serial, ref_bid, ref_nid)
		VALUES(_floor_id, _map_block_id, in_car_status, in_plate_provice_code, in_plate_provice_char, in_plate_city_code, in_plate_number, in_plate_literal, in_img_plate_history, _block_name, in_serial, in_bid, in_nid);
	
	EXCEPTION
		WHEN TOO_MANY_ROWS THEN
			RAISE EXCEPTION 'DEVICE ANGEL NOT UNIQUE';

END

$$ LANGUAGE plpgsql;


-- function
-- function
CREATE OR REPLACE FUNCTION nf_update_plate_crop(in_serial varchar(12), in_bid integer, in_nid integer,
			in_img_car_crop_unique varchar(128), in_img_car_crop_history varchar(128))
RETURNS void AS $$
DECLARE
	_map_block_id varchar;
	_floor_id integer;
	_result_history_id integer;

BEGIN
	-- find some info out
	SELECT ref_map_block_id INTO STRICT _map_block_id FROM tbl_jfcp_angle_param WHERE device_serial = in_serial and bid = in_bid and nid = in_nid;
	SELECT ref_floor_id INTO STRICT _floor_id FROM tbl_jfc_device WHERE device_serial = in_serial;
	
	-- update the result table
	UPDATE tbl_jfcp_plate_result
	SET img_crop = in_img_car_crop_unique
	WHERE ref_floor_id = _floor_id AND ref_map_block_id = _map_block_id;
	
	-- update the history result
	SELECT id INTO _result_history_id 
	FROM tbl_jfcp_plate_result_history 
	WHERE ref_device_serial = in_serial and ref_bid = in_bid and ref_nid = in_nid 
	ORDER BY id DESC 
	LIMIT 1;

	UPDATE tbl_jfcp_plate_result_history
	SET img_crop = in_img_car_crop_history
	WHERE id = _result_history_id;
END

$$ LANGUAGE plpgsql;