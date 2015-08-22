-- 压力数据

-- function
CREATE OR REPLACE FUNCTION nf_test_press_data(_in_record_count integer)
RETURNS void AS $$
DECLARE
	_serial varchar;
	_bid integer;
	_nid integer;
	_map_block_id varchar;
	_count integer;

	len_count integer;
	len_fill integer;
	prestr varchar;
BEGIN
	_count := 1;
	_bid := 0;
	_nid := 0;
	delete from public.tbl_jfc_device where remark='press_data_p2';
	delete from public.tbl_jfcp_angle_param where remark='press_data_p2';
	
LOOP
	len_count  := length(_count::varchar);
	
	len_fill := 12 - len_count;
	prestr := rpad('ptets_', len_fill, '0');
	

	_serial:= prestr || _count::varchar;
	_map_block_id:= 'r' || _serial::varchar;

		INSERT INTO public.tbl_jfc_device(
		    device_serial, ref_floor_id, position_x, position_y, ble_uuid, 
		    ble_major_id, ble_minor_id, wifi_ssid, wifi_key, remark)
		VALUES (_serial, 3, 100, 100, 'B9407F30-F5F8-466E-AFF9-25556B57FE6D',  
		    10, 1, 'my3579', 'mymy3579', 'press_data_p2');

		-- some computations
		INSERT INTO tbl_jfcp_angle_param (device_serial, ref_map_block_id, bid, bangle, nid, nangle, crop_x, crop_y, crop_w, crop_h, remark) 
			VALUES (_serial, _map_block_id, _bid, 60, _nid, 50, 400, 300, 640, 320, 'press_data_p2');
		
		_count := _count + 1;
	
	EXIT WHEN _count > _in_record_count; 
	END LOOP;
END

$$ LANGUAGE plpgsql;