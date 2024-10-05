package mysql

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate()
	if err != nil {

	}
}
