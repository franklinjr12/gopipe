select
    start_byte,
    end_byte,
    data_type_id
from
    application_data_structures
where
    application_id = $1
    and version = $2