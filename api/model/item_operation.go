package model

type ItemOperation struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	OperationId int64  `xorm:"not null default 0 comment('日志编号') BIGINT(20)"`
	Submitter   string `xorm:"not null default '无' comment('操作人') index VARCHAR(50)"`
	ItemId      int64  `xorm:"not null default 0 comment('菜单编号') index BIGINT(20)"`
	ItemName    string `xorm:"default '' comment('菜品名称') VARCHAR(255)"`
	OperateType int    `xorm:"not null default 0 comment('操作类型，1:新增商品，2:修改商品，3:删除商品') TINYINT(3)"`
	BeforeEdit  string `xorm:"not null default '' comment('修改前的字段') VARCHAR(5000)"`
	AfterEdit   string `xorm:"default '' comment('修改后的字段') VARCHAR(5000)"`
	CreateTime  int64  `xorm:"not null default 0 comment('新增时间') index BIGINT(20)"`
	ModifyTime  int64  `xorm:"not null default 0 comment('修改时间') BIGINT(20)"`
	IsDelete    int    `xorm:"not null default 0 comment('是否删除，0：未删除 1：已删除') TINYINT(3)"`
}
