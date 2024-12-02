package repository

type CreateUserTxParams struct {
	AddUserBaseParams
	AfterCreate func(user PreGoAccUserBase9999) error
}

type CreateUserTxResult struct {
	User PreGoAccUserBase9999
}

//func (s *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
//	var result = CreateUserTxResult{}
//
//	err := s.execTx(ctx, func(q *Queries) error {
//		var err error
//
//		result.User, err = q.AddUserBase(ctx, arg.AddUserBaseParams)
//		if err != nil {
//			return err
//		}
//
//	})
//	return result, nil
//}
