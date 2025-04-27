package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type TabUser struct {
	Id           int       `orm:"column(user_id);auto" description:"Standard field for the primary key"`
	Name         string    `orm:"column(name);size(255)" description:"A regular string field"`
	Email        string    `orm:"column(email);size(255);null" description:"A pointer to a string, allowing for null values"`
	Age          int8      `orm:"column(age)" description:"An unsigned 8-bit integer"`
	Birthday     time.Time `orm:"column(birthday);type(datetime);null" description:"A pointer to time.Time, can be null"`
	MemberNumber string    `orm:"column(member_number);size(255);null" description:"Uses sql.NullString to handle nullable strings"`
	Remark       string    `orm:"column(remark);size(128);null" description:"备注"`
	ActivatedAt  time.Time `orm:"column(activated_at);type(datetime);null" description:"Uses sql.NullTime for nullable time fields"`
	CreatedAt    time.Time `orm:"column(created_at);type(datetime);auto_now_add" description:"Automatically managed by GORM for creation time"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(datetime);auto_now_add" description:"Automatically managed by GORM for update time"`
}

func (t *TabUser) TableName() string {
	return "tab_user"
}

func init() {
	orm.RegisterModel(new(TabUser))
}

// AddTabUser insert a new TabUser into database and returns
// last inserted Id on success.
func AddTabUser(m *TabUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTabUserById retrieves TabUser by Id. Returns error if
// Id doesn't exist
func GetTabUserById(id int) (v *TabUser, err error) {
	o := orm.NewOrm()
	v = &TabUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTabUser retrieves all TabUser matches certain condition. Returns empty list if
// no records exist
func GetAllTabUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TabUser))
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

	var l []TabUser
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

// UpdateTabUser updates TabUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateTabUserById(m *TabUser) (err error) {
	o := orm.NewOrm()
	v := TabUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTabUser deletes TabUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTabUser(id int) (err error) {
	o := orm.NewOrm()
	v := TabUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TabUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
