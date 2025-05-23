package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type TabUserCard struct {
	Id         int       `orm:"column(user_card_id);auto" description:"用户卡ID"`
	UserId     int64     `orm:"column(user_id)" description:"用户ID"`
	CardTypeId int64     `orm:"column(card_type_id)" description:"卡类型ID"`
	Remark     string    `orm:"column(remark);size(255);null" description:"登录地点"`
	Status     string    `orm:"column(status);size(1);null" description:"卡状态（1=在用 2=停用）"`
	CreatedAt  time.Time `orm:"column(created_at);type(datetime);auto_now_add" description:"创建时间"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(datetime);auto_now_add" description:"修改时间"`
}

func (t *TabUserCard) TableName() string {
	return "tab_user_card"
}

func init() {
	orm.RegisterModel(new(TabUserCard))
}

// AddTabUserCard insert a new TabUserCard into database and returns
// last inserted Id on success.
func AddTabUserCard(m *TabUserCard) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTabUserCardById retrieves TabUserCard by Id. Returns error if
// Id doesn't exist
func GetTabUserCardById(id int) (v *TabUserCard, err error) {
	o := orm.NewOrm()
	v = &TabUserCard{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTabUserCard retrieves all TabUserCard matches certain condition. Returns empty list if
// no records exist
func GetAllTabUserCard(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TabUserCard))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []TabUserCard
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTabUserCard updates TabUserCard by Id and returns error if
// the record to be updated doesn't exist
func UpdateTabUserCardById(m *TabUserCard) (err error) {
	o := orm.NewOrm()
	v := TabUserCard{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTabUserCard deletes TabUserCard by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTabUserCard(id int) (err error) {
	o := orm.NewOrm()
	v := TabUserCard{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TabUserCard{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
