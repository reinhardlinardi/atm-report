package transactionrepo

type Repository interface {
	InsertRows(data []Transaction) (int64, error)
	CountDaily() ([]DailyCount, error)
	CountByType() ([]ByTypeCount, error)
	CountDailyByType() ([]DailyByTypeCount, error)
	MaxWithdrawDaily() ([]DailyMaxWithdraw, error)
}
