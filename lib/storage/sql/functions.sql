delimiter //

create or replace function weasel_storage.save_files(
	  _id bigint,
    _reference_id bigint,
    _name   varchar,
    _alias varchar,
    _pid bigint,
    _fields jsonb,
    _is_group boolean
) returns bigint as $$

    declare
        _version bigint;
        _nid bigint;
    begin

		if _id = 0 then
			select into _version 0;
			update weasel_classifiers.counter set total = total + 1 where reference_id = _reference_id;
		else
	    	select into _version max(ver)+1 from weasel_classifiers.items where reference_id = _reference_id and "alias" = _alias;
	    	update weasel_classifiers.items set ver = _version, updated_at = current_timestamp where id = _id;
	   end if;

	   insert into weasel_classifiers.items (reference_id, name, alias, pid, fields, is_group, created_at, parents) values
	   (_reference_id, _name, _alias, _pid, _fields, _is_group, current_timestamp,
			(select coalesce(array_to_string(array_append(string_to_array(parents, '.')::bigint[], _pid), '.'),'') from weasel_classifiers.items where id = _pid)
		) returning id into _nid;

    	return coalesce(_nid, 0);

    end;
$$ language plpgsql;