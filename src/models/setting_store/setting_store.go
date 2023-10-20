package setting_store

type SettingStore interface {
	GetLastSetting() (Setting, error)
	CreateSetting(request CreateSettingRequest) (Setting, error)
}
