select
    id
from
    applications
where
    name = $1
    and company_id = $2