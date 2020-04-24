/**
 * @Author: alienongwlx@gmail.com
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/18 10:28
 */

package crud_xorm

import (
	"fmt"
	"github.com/xormplus/builder"
	"github.com/xormplus/xorm"
)

/**
@description: ICRUDService Interface,Must Implement These Methods
@method Insert: C
@method Update: U
@method Delete: D
@method Deletes: D
@method Get: Get One Record
@method Query: R
*/
type ICRUDService interface {
	Insert(bean interface{}) error
	Update(id interface{}, bean interface{}, fields []string) error
	Delete(id interface{}, bean interface{}) error
	Deletes(ids []interface{}, bean interface{}) error
	Get(id interface{}, bean interface{}) error
	Query(m *QueryParam, bean interface{}, beans interface{}) (int64, error)
}

/**
@description: ICRUDEvent Interface,Can Implement One or All Methods
@method BeforeInsert: the method before insert
@method AfterInsert: the method after insert
@method BeforeUpdate: the method before update
@method AfterUpdate: the method after update
@method BeforeDelete: the method before update
@method AfterDelete: the method after update
*/
type ICRUDEvent interface {
	BeforeInsert(bean interface{})
	AfterInsert(bean interface{})
	BeforeUpdate(bean interface{})
	AfterUpdate(bean interface{})
	BeforeDelete(bean interface{})
	AfterDelete(bean interface{})
}

/**
@description: Struct Set ICRUDEvent Interface,Can Implement One or All Methods
*/
type CRUDEvent struct {
	ICRUDEvent
}

/**
@description: CRUDService
@attribute db: Xorm DB Engine
@attribute session: Xorm DB Engine Session
@attribute event: CRUDEvent
*/
type CRUDService struct {
	db      *xorm.Engine
	session *xorm.Session
	event   *CRUDEvent
}

/**
@description: Some Implement Methods
*/
func (s *CRUDService) BeforeInsert(bean interface{}) {}

func (s *CRUDService) AfterInsert(bean interface{}) {}

func (s *CRUDService) BeforeUpdate(bean interface{}) {}

func (s *CRUDService) AfterUpdate(bean interface{}) {}

func (s *CRUDService) BeforeDelete(bean interface{}) {}

func (s *CRUDService) AfterDelete(bean interface{}) {}

/**
@description: NewCRUD
@param db: Xorm Engine
@param c: ICRUDEvent
@return: CRUDService
*/
func NewCRUD(db *xorm.Engine, c ICRUDEvent) *CRUDService {
	crud := new(CRUDService)
	crud.db = db
	crud.event = &CRUDEvent{c}
	return crud
}

/**
@description: DB
@param : nil
@return: xorm.Engine
*/
func (s *CRUDService) DB() *xorm.Engine {
	return s.db
}

/**
@description: Set The DB
@param db: xorm.Engine
@return:
*/
func (s *CRUDService) SetDB(db *xorm.Engine) {
	s.db = db
}

/**
@description: Session
@param : nil
@return: Session
*/
func (s *CRUDService) Session() *xorm.Session {
	if s.session != nil {
		return s.session
	}
	return s.db.NewSession()
}

/**
@description: SetSession
@param session: xorm.Session
@return:
*/
func (s *CRUDService) SetSession(session *xorm.Session) {
	s.session = session
}

/**
@description: Insert
@param bean: insert record
@return: error
*/
func (s *CRUDService) Insert(bean interface{}) error {
	if s.event != nil {
		s.event.BeforeInsert(bean)
		_, err := s.Session().Insert(bean)
		if err != nil {
			return err
		}
		s.event.AfterInsert(bean)
		return err
	}
	_, err := s.Session().Insert(bean)
	return err
}

/**
@description: Update
@param id: Record Id
@param bean: Record Updated
@param fields: Update Some Fields
@return: error
*/
func (s *CRUDService) Update(id interface{}, bean interface{}, fields []string) error {
	var sess = s.Session().Where(s.DB().Quote("id")+"=?", id)
	if len(fields) == 0 {
		sess.AllCols()
	} else {
		sess.Cols(fields...)
	}
	if s.event != nil {
		s.event.BeforeUpdate(bean)
		_, err := sess.Update(bean)
		if err != nil {
			return err
		}
		s.event.AfterUpdate(bean)
		return err
	}
	_, err := sess.Update(bean)
	return err
}

/**
@description: Delete One Record
@param id: Record Id
@param bean: Record Updated
@return: error
*/
func (s *CRUDService) Delete(id interface{}, bean interface{}) error {
	ti, err := s.DB().TableInfo(bean)
	if err != nil {
		return err
	}
	deleteSQL := fmt.Sprintf(`DELETE FROM %s WHERE 1=1`, ti.Name)
	var args []interface{}
	str, arg, _ := builder.ToSQL(builder.Eq{s.DB().Quote("id"): id})
	deleteSQL += " AND " + str + " "
	args = append(args, arg...)
	if s.event != nil {
		s.event.BeforeDelete(bean)
		_, err := s.Session().SQL(deleteSQL, args...).Execute()
		if err != nil {
			return err
		}
		s.event.AfterDelete(bean)
		return err
	}
	_, err = s.Session().SQL(deleteSQL, args...).Execute()
	return err
}

/**
@description: Delete Some Record
@param id: Record Ids
@param bean: Record Updated
@return: error
*/
func (s *CRUDService) Deletes(ids []interface{}, bean interface{}) error {
	ti, err := s.DB().TableInfo(bean)
	if err != nil {
		return err
	}
	deleteSQL := fmt.Sprintf(`DELETE FROM %s WHERE 1=1`, ti.Name)
	var args []interface{}
	conf := builder.In(s.DB().Quote("_id"), ids)
	str, arg, _ := builder.ToSQL(conf)
	deleteSQL += " AND " + str + " "
	args = append(args, arg...)
	_, err = s.Session().SQL(deleteSQL, args...).Execute()
	//logging.Info("已删除: %d", n)
	return err
}

/**
@description: Get The Record By Id
@param id: Record Id
@param bean: Record Updated
@return: error
*/
func (s *CRUDService) Get(id interface{}, bean interface{}) error {
	_, err := s.Session().Where(s.DB().Quote("_id")+"=?", id).Get(bean)
	return err
}

/**
@description: Query The Records
@param m: Query Condition
@param bean: Record
@param beans: Records Return
@return: Count And Error
*/
func (s *CRUDService) Query(m *QueryParam, bean interface{}, beans interface{}) (int64, error) {
	return m.Query(s.Session(), bean, beans)
}
