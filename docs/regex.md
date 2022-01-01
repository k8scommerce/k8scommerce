# Common regex

match 
^\s{2}([A-Z|\[].*)\s(.*)
replace
  $2 $1 `json:"$2,omitempty"`