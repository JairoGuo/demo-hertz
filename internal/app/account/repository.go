package account

func Create(accounts []*Account) error {
	return DB.Create(accounts).Error
}

func FindAccountByUsername(username string) ([]*Account, error) {
	res := make([]*Account, 0)
	if err := DB.Where(DB.Where("username = ?", username)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckAccount(account, password string) ([]*Account, error) {
	res := make([]*Account, 0)
	if err := DB.Where("username = ?", account).Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
