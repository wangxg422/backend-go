package system

import (
	"backend/common/enmu"
	"backend/global"
	sysModel "backend/model/system"
)

type SysMenuService struct {
}

func (m *SysMenuService) CreateMenu(d *sysModel.SysMenu) error {
	res := global.DB.Create(&d)

	return res.Error
}

func (m *SysMenuService) UpdateMenu(d *sysModel.SysMenu) error {
	res := global.DB.Model(&sysModel.SysMenu{MenuId: d.MenuId}).Updates(&d)
	return res.Error
}

func (m *SysMenuService) DeleteMenu(id int64) error {
	res := global.DB.Model(&sysModel.SysMenu{MenuId: id}).Update("del_flag", enmu.EnmuGroupApp.DelFlagDelete.GetCode())
	return res.Error
}

func (m *SysMenuService) ListMenu() ([]sysModel.SysMenu, error) {
	var list []sysModel.SysMenu

	res := global.DB.Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode()).Find(&list)

	return list, res.Error
}

func (m *SysMenuService) GetMenuById(id int64) (sysModel.SysMenu, error) {
	user := sysModel.SysMenu{
		MenuId: id,
	}

	res := global.DB.Take(&user, id).Where("del_flag = ?", enmu.EnmuGroupApp.StatusNormal.GetCode())
	return user, res.Error
}
