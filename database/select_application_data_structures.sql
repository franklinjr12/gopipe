select
	a.start_byte,
	a.end_byte,
	d.name as type
from
	application_data_structures a
	inner join data_types d on d.id = a.data_type_id
where
	a.application_id = $1
	and a.version = $2
order by
	a.start_byte