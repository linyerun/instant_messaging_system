# im_system
## gorm
- 获取记录数为0, 再获取error会出现err != nil
- 设置零值使用`default`
```go
type Friend struct {
	Email01   string    `gorm:"column:email01"`
	Email02   string    `gorm:"column:email02"`
	RoomId    uint      `gorm:"column:room_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:NULL"` // 设置默认值
}
```
- 使用 `first`获取一条参数
```go
func GetUserByEmail(email string) (user *models.User, err error) {
	user = new(models.User)
	err = MySQLClient.First(user, "email = ?", email).Error
	return
}
```

