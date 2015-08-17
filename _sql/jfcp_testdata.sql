-- area
INSERT INTO tbl_area (id, area_code, area_name, longitude, latitude, province, city_name, city_code, remark) 
	VALUES (2, 'SZ-HANANCHENG', '海岸城', 113.941751999999994, 22.5229919999999986, '广东', '深圳', '755', '深圳海岸城');
INSERT INTO tbl_area (id, area_code, area_name, longitude, latitude, province, city_name, city_code, remark) 
	VALUES (3, 'SZ-XIANGNAN', '向南瑞峰', 113.932162000000005, 22.5299359999999993, '广东', '深圳', '755', '深圳南山区桂庙路22号向南瑞峰');

-- building
INSERT INTO tbl_building (id, ref_area_id, building_code, building_name, longitude, latitude, remark) 
	VALUES (1, 2, 'HANC-WEST', '西座', 113.944405000000003, 22.5243730000000006, '深圳海岸城西座');
INSERT INTO tbl_building (id, ref_area_id, building_code, building_name, longitude, latitude, remark) 
	VALUES (2, 2, 'HANC-EAST', '东座', 113.943541999999994, 22.5239229999999999, '深圳海岸城东座');
INSERT INTO tbl_building (id, ref_area_id, building_code, building_name, longitude, latitude, remark) 
	VALUES (3, 3, 'XN-BUSSINESS', '创业中心', 113.932114999999996, 22.5296769999999995, '向南瑞峰创业中心');

-- floor 

INSERT INTO tbl_floor (id, ref_building_id, floor_number, floor_code, floor_name, floor_type, remark) 
	VALUES (1, 1, -1, 'HANCW-B1', 'B1', 2, '海岸城西座地下停车场');
INSERT INTO tbl_floor (id, ref_building_id, floor_number, floor_code, floor_name, floor_type, remark) 
	VALUES (2, 1, 1, 'HANCW-L1', 'L1', 1, '海岸城西座1楼商铺');
INSERT INTO tbl_floor (id, ref_building_id, floor_number, floor_code, floor_name, floor_type, remark) 
	VALUES (3, 1, 2, 'HANCW-L2', 'L2', 1, '海岸城西座2楼商铺');
INSERT INTO tbl_floor (id, ref_building_id, floor_number, floor_code, floor_name, floor_type, remark) 
	VALUES (4, 1, 3, 'HANCW-L3', 'L3', 1, '海岸城西座3楼商铺');
INSERT INTO tbl_floor (id, ref_building_id, floor_number, floor_code, floor_name, floor_type, remark) 
	VALUES (7, 2, 1, 'HANCE-L1', 'L1', 1, '海岸城东座1楼商铺');
INSERT INTO tbl_floor (id, ref_building_id, floor_number, floor_code, floor_name, floor_type, remark) 
	VALUES (5, 2, -1, 'HANCE-B1', 'B1', 2, '海岸城东座地下停车场');
INSERT INTO tbl_floor (id, ref_building_id, floor_number, floor_code, floor_name, floor_type, remark) 
	VALUES (8, 3, 3, 'SZ-XN-L3', 'L3', 1, '向南瑞峰创业中心3楼');

-- device
INSERT INTO public.tbl_jfc_device(device_serial, ref_floor_id, position_x, position_y, ble_uuid, ble_major_id, ble_minor_id, wifi_ssid, wifi_key)
    VALUES ('sz1234567890', 1, 100, 100, 'B9407F30-F5F8-466E-AFF9-25556B57FE6D', 100, 1, 'my3579', '12345');


-- jfcp angle

INSERT INTO tbl_jfcp_angle_param (device_serial, ref_map_block_id, bid, bangle, nid, nangle, crop_x, crop_y, crop_w, crop_h, remark) 
	VALUES ('sz1234567890', 'rB2186', 1, 30, 1, 20, 300, 300, 640, 320, '');

INSERT INTO tbl_jfcp_angle_param (device_serial, ref_map_block_id, bid, bangle, nid, nangle, crop_x, crop_y, crop_w, crop_h, remark) 
	VALUES ('sz1234567890', 'rB2187', 1, 30, 2, 50, 400, 300, 640, 320, '');

INSERT INTO tbl_jfcp_angle_param (device_serial, ref_map_block_id, bid, bangle, nid, nangle, crop_x, crop_y, crop_w, crop_h, remark) 
	VALUES ('sz1234567890', 'rB2188', 1, 30, 3, 70, 300, 350, 640, 320, '');

INSERT INTO tbl_jfcp_angle_param (device_serial, ref_map_block_id, bid, bangle, nid, nangle, crop_x, crop_y, crop_w, crop_h, remark) 
	VALUES ('sz1234567890', 'rB2016', 2, 130, 1, 40, 300, 300, 640, 320, '');

INSERT INTO tbl_jfcp_angle_param (device_serial, ref_map_block_id, bid, bangle, nid, nangle, crop_x, crop_y, crop_w, crop_h, remark) 
	VALUES ('sz1234567890', 'rB2017', 2, 130, 2, 60, 300, 300, 640, 320, '');

INSERT INTO tbl_jfcp_angle_param (device_serial, ref_map_block_id, bid, bangle, nid, nangle, crop_x, crop_y, crop_w, crop_h, remark) 
	VALUES ('sz1234567890', 'rB2018', 2, 130, 3, 80, 300, 300, 640, 320, '');

INSERT INTO tbl_jfcp_angle_param (device_serial, ref_map_block_id, bid, bangle, nid, nangle, crop_x, crop_y, crop_w, crop_h, remark) 
	VALUES ('sz1234567890', 'rB2019', 2, 130, 4, 90, 300, 300, 640, 320, '');