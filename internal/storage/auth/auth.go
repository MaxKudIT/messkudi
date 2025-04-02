package auth

//	func (us *userStorage) GetUserId(userCr auth.UserCredentials) uuid.UUID {
//		var id uuid.UUID
//		var hashhex string
//		const GET_USERPASS_QUERY = "SELECT password FROM users WHERE name = $1"
//		if err := userRep.db.QueryRow(GET_USERPASS_QUERY, userCr.Name).Scan(&hashhex); err != nil {
//			log.Print(exceptions.GetUserExc().Error())
//		}
//		hash, _ := hex.DecodeString(hashhex)
//		utils2.CompareToHash(hash, []byte(userCr.Password))
//		const GET_USERID_QUERY = "SELECT id FROM users WHERE name = $1 AND password = $2"
//
//		if err := userRep.db.QueryRow(GET_USERID_QUERY, userCr.Name, hashhex).Scan(&id); err != nil {
//			log.Print(exceptions.GetUserExc().Error())
//		}
//		return id
//	}
