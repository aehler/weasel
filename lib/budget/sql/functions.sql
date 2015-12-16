delimiter //

create or replace function weasel_main.save_budget_operation(
	  _id bigint,
    _user_id bigint,
    _organization_id bigint,
    _sum numeric,
    _date_op text,
    _user_meta jsonb,
    _dims_meta jsonb
) returns bigint as $$

    declare
        _nid bigint;
        dt TIMESTAMP;
    begin

    select into dt _date_op::TIMESTAMP;

    update weasel_main.budget_operations set "sum" = _sum, date_op = dt, updated_at = current_timestamp,
    user_meta = _user_meta, dims_meta = _dims_meta
        where id = _id RETURNING id INTO _nid;

	   insert into weasel_main.budget_operations (user_id, organization_id, "sum", date_op, user_meta, dims_meta)
	   select _user_id, _organization_id, _sum, dt, _user_meta, _dims_meta
	   where NOT EXISTS(select 1 from weasel_main.budget_operations where id = _id)
			returning id into _nid;

    	return coalesce(_nid, 0);

    end;
$$ language plpgsql;


delimiter //

create or replace function weasel_main.get_plan_rows(
 _oid bigint,
 _period_id bigint
 ) returns setof record as $$

begin

return query
	with ref_trunc as ( select id,
		jsonb_array_elements(dims_meta)->>'ReferenceAlias' as rkey,
		jsonb_array_elements(dims_meta)->>'Label' as rlabel,
		jsonb_array_elements(dims_meta)->>'Value' as rvalue
		from budget_plan
		)

	select distinct on (budget_plan.id) budget_operations.id, sum, date_op, dims_meta, user_meta from budget_plan
	left join ref_trunc on ref_trunc.id = budget_plan.id
	where organization_id = _oid
	and
		ref_trunc.rkey = 'period' and ref_trunc.rvalue = _period_id::text
	order by date_op desc;

    end;
$$ language plpgsql;