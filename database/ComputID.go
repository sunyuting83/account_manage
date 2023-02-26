package database

type Comput struct {
	ID         uint `gorm:"primaryKey"`
	UsersID    uint
	ProjectsID uint
	ComputCode string `gorm:"index"`
	Accounts   []Accounts
	CreatedAt  int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt  int64 `gorm:"autoUpdateTime:milli"`
}

// 新建后台 生成唯一码 写入Leveldb key 编码 value UsersID ProjectsID 通过唯一key生成url
// 每次取号，通过key获取userid和projectsID 需要机器码的地方 将机器码写入数据库
// 机器码入库的时候自动生成Cache key computid value postgres ComputID
// 需要机器码寻找的时候， 先从Cache读取Computid 然后去库中查找 执行动作
