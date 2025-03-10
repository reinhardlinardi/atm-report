package transaction

type Repository interface {
	Load([]Transaction) (int64, error)
	CountDaily() ([]DailyCount, error)
	CountDailyByType() ([]DailyTypeCount, error)
	GetDailyMaxWithdraw() ([]DailyMaxWithdraw, error)
}
