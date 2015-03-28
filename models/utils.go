package models

func zname(name, entity string) string {
	return db_prefix + name + "_" + entity
}
