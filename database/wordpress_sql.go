package database

const (
	SqlWordpressPosts = `select p.ID, 
							p.post_content, 
							p.post_title, 
							p.post_status, 
							p.post_name, 
							p.post_date from wiku_posts p
			left join wiku_term_relationships tr on p.ID = tr.object_id
			left join wiku_term_taxonomy tt  on tr.term_taxonomy_id = tt.term_taxonomy_id
			left join wiku_terms t  on tt.term_id = t.term_id
				where t.term_id = ?`
)