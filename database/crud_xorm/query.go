/**
 * @Author: alienongwlx@gmail.com
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/20 10:33
 */

package crud_xorm

import (
	"errors"
	"fmt"
	"github.com/xormplus/builder"
	"github.com/xormplus/xorm"
	"regexp"
	"strings"
)

//SQL injection
var InjectionMatch = regexp.MustCompile(`(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`)

/**
@description: QueryParam
@attribute Size: Query Records Count
@attribute Index: Query Records From Which Start
@attribute Count: Is Calculate The Count of Records
@attribute Conditions: Condition List
@attribute Order: Orders Map,-1 Decs,1 Asc
@attribute Fields: Select Which Columns
@attribute Params: Some Other Params
*/
type QueryParam struct {
	Size       int                    `json:"size"`
	Index      int                    `json:"index"`
	Count      bool                   `json:"count"`
	Conditions []Condition            `json:"conditions"`
	Order      []map[string]int       `json:"order"`
	Fields     []map[string]int       `json:"fields"`
	Params     map[string]interface{} `json:"query"`
}

/**
@description: Condition
@attribute Field: Field
@attribute Relation: "and", "or"
@attribute Operate: "like" , ">=" , "<=" , ">", "<", "nil" , "!=" , "in", "between"
@attribute Value: Value
*/
type Condition struct {
	Field    string      `json:"Field"`
	Relation string      `json:"Relation"`
	Operate  string      `json:"Operate"`
	Value    interface{} `json:"Value"`
}

/**
@description: NewQuery
@param : nil
@return: QueryParam
*/
func NewQuery() *QueryParam {
	var q = new(QueryParam)
	q.Count = true
	q.Index = 1
	q.Conditions = make([]Condition, 0)
	q.Order = make([]map[string]int, 0)
	q.Fields = make([]map[string]int, 0)
	q.Params = make(map[string]interface{}, 0)
	return q
}

/**
@description: NewCondition
@param Field: nil
@return: QueryParam
*/
func (q *QueryParam) NewCondition(Field string, Relation string, Operate string, Value interface{}) *QueryParam {
	c := Condition{Field: Field, Relation: Relation, Operate: Operate, Value: Value}
	q.Conditions = append(q.Conditions, c)
	return q
}

/**
@description: Query Orders
@param session: xorm.Session
@return: error
*/
func (q *QueryParam) Orders(session *xorm.Session) error {
	for _, order := range q.Order {
		for name, i := range order {
			if i > 0 {
				session.Asc(name)
			} else {
				session.Desc(name)
			}
		}
	}
	return nil
}

/**
@description: Query Limit
@param session: xorm.Session
@return:
*/
func (q *QueryParam) Limit(session *xorm.Session) {
	if q.Size > 0 {
		session.Limit(q.Size, (q.Index-1)*q.Size)
	}
}

/**
@description: Query Select Cols
@param session: xorm.Session
@return: error
*/
func (q *QueryParam) Cols(session *xorm.Session) error {
	for _, field := range q.Fields {
		var fs []string
		var nfs []string
		for k, v := range field {
			if v > 0 {
				fs = append(fs, k)
			} else {
				nfs = append(nfs, k)
			}
		}
		if len(fs) > 0 {
			session.Cols(fs...)
		}
		if len(nfs) > 0 {
			session.Omit(nfs...)
		}
	}
	return nil
}

/**
@description: Query Filter Condition
@param session: xorm.Session
@return: error
*/
func (q *QueryParam) Filter(session *xorm.Session) error {
	for k, v := range q.Params {
		q.NewCondition(k, "and", "=", v)
	}
	return HandleConditions(session, q)
}

/**
@description: Query GetItems
@param session: xorm.Session
@return: error
*/
func (q *QueryParam) GetItems(session *xorm.Session) error {
	err := q.Orders(session)
	if err != nil {
		return err
	}
	err = q.Cols(session)
	if err != nil {
		return err
	}
	q.Limit(session)
	return nil
}

/**
@description: Query
@param session: xorm.Session
@param bean: Use For Counting Of The Results
@param beans:  Use For Storing The Results
@return: count,error
*/
func (q *QueryParam) Query(sess *xorm.Session, bean interface{}, beans interface{}) (int64, error) {
	q.Filter(sess)
	count := int64(0)
	if q.Count {
		count, _ = sess.Count(bean)
	}
	err := q.GetItems(sess)
	if err != nil {
		return 0, err
	}
	err = sess.Find(beans)
	if err != nil {
		return 0, err
	}
	return count, nil
}

/**
@description: Query By Sql String
@param session: xorm.Session
@param sql: Sql String
@param bean: Use For Counting Of The Results
@param beans:  Use For Storing The Results
@param sqlArgs: params from Sql String
@return: count,error
*/
func (q *QueryParam) SqlQuery(sess *xorm.Session, sql string, bean interface{}, beans interface{}, sqlArgs ...interface{}) (int64, error) {
	selectFields := make([]string, 0)
	for _, field := range q.Fields {
		for k, v := range field {
			if InjectionMatch.MatchString(k) {
				//logging.Warn("发现SQL注入：%v", k)
				return 0, errors.New("发现SQL注入：" + fmt.Sprintf("%v", k))
			}
			if v > 0 {
				selectFields = append(selectFields, k)
			}
		}
	}
	var selectStr string
	if len(selectFields) == 0 {
		selectStr = "*"
	} else {
		selectStr = strings.Join(selectFields, ",")
	}
	sql = `select ` + selectStr + ` from ( ` + sql + ` ) temptable where 1=1 `
	for k, v := range q.Params {
		q.NewCondition(k, "and", "=", v)
	}
	args, err := q.HandleSqlConditions(&sql)
	if err != nil {
		return 0, err
	}
	sqlArgs = append(sqlArgs, args...)
	count := int64(0)
	if q.Count {
		count, _ = sess.SQL(`select count(*) from ( `+sql+` ) temptable`, sqlArgs...).Count(bean)
	}
	sql += q.HandleSqlOrders(sess)
	sess.SQL(sql, sqlArgs...)
	q.Limit(sess)
	err = sess.Find(beans)
	if err != nil {
		return 0, err
	}
	return count, nil
}

/**
@description: HandleSqlConditions,Parse QueryParam's Conditions  to Sql
@param sql: sql String
@return: sql,count,error
*/
func (q *QueryParam) HandleSqlConditions(sql *string) (args []interface{}, err error) {
	type QueryItem struct {
		Relation string
		Str      string
		Args     []interface{}
	}
	qis := make([]QueryItem, 0)
	for _, condition := range q.Conditions {
		if InjectionMatch.MatchString(condition.Field) {
			return nil, errors.New("无效的字段名称：" + fmt.Sprintf("%v", condition.Field))
		}
		var conf interface{}
		if len(condition.Operate) == 0 {
			condition.Operate = "="
		}
		condition.Field = "temptable." + condition.Field
		switch condition.Operate {
		case "like":
			conf = builder.Like{condition.Field, condition.Value.(string)}
		case "=":
			conf = builder.Eq{condition.Field: condition.Value}
		case ">=":
			conf = builder.Gte{condition.Field: condition.Value}
		case ">":
			conf = builder.Gt{condition.Field: condition.Value}
		case "<=":
			conf = builder.Lte{condition.Field: condition.Value}
		case "<":
			conf = builder.Lt{condition.Field: condition.Value}
		case "nil":
			conf = builder.IsNull{condition.Field}
		case "nnil":
			conf = builder.NotNull{condition.Field}
		case "!=":
			conf = builder.Neq{condition.Field: condition.Value}
		case "in":
			o := condition.Value.([]interface{})
			if len(o) > 0 {
				conf = builder.In(condition.Field, o)
			} else {
				conf = builder.Eq{"ID": "NULL"}
			}
		case "between":
			o := condition.Value.([]interface{})
			conf = builder.Between{Col: condition.Field, LessVal: o[0], MoreVal: o[1]}
		}
		str, args, err := builder.ToSQL(conf)
		if err != nil {
			//	logging.Warn("条件错误：%v", err)
			return nil, errors.New("条件错误：" + fmt.Sprintf("%v", err))
		}
		qi := QueryItem{
			Relation: condition.Relation,
			Str:      str,
			Args:     args,
		}
		qis = append(qis, qi)
	}

	retArgs := make([]interface{}, 0)
	for _, qi := range qis {
		if qi.Relation == "and" {
			*sql += " AND (" + qi.Str + ")"
			retArgs = append(retArgs, qi.Args...)
		} else if qi.Relation == "and" {
			*sql += " OR (" + qi.Str + ")"
			retArgs = append(retArgs, qi.Args...)
		}
	}
	return retArgs, err
}

/**
@description:  HandleSqlOrders,Parse QueryParam's Orders  to Sql
@param sql: sql String
@return: sql,count,error
*/
func (q *QueryParam) HandleSqlOrders(session *xorm.Session) string {
	var strOrder = ""
	if len(q.Order) > 0 {
		strOrder = " order by "
		var items = make([]string, 0)
		for _, order := range q.Order {
			for name, i := range order {
				name = "temptable." + name
				if i > 0 {
					items = append(items, name+" ASC")
				} else {
					items = append(items, name+" DESC")
				}
			}
		}
		strOrder += strings.Join(items, ",")
	}
	return strOrder
}

/**
@description: Attach  QueryParam's Conditionss  to session
@param session: xorm.Session
@param q: QueryParam
@return: sql,count,error
*/
func HandleConditions(session *xorm.Session, q *QueryParam) error {
	type QueryItem struct {
		Relation string
		Str      string
		Args     []interface{}
	}
	qis := make([]QueryItem, 0)
	for _, condition := range q.Conditions {
		if InjectionMatch.MatchString(condition.Field) {
			//logging.Warn("发现SQL注入：%v", condition)
			return errors.New("无效的Field名称：" + fmt.Sprintf("%v", condition.Field))
		}
		var conf interface{}
		if len(condition.Operate) == 0 {
			condition.Operate = "="
		}
		switch condition.Operate {
		case "like":
			conf = builder.Like{condition.Field, condition.Value.(string)}
		case "=":
			conf = builder.Eq{condition.Field: condition.Value}
		case ">=":
			conf = builder.Gte{condition.Field: condition.Value}
		case ">":
			conf = builder.Gt{condition.Field: condition.Value}
		case "<=":
			conf = builder.Lte{condition.Field: condition.Value}
		case "<":
			conf = builder.Lt{condition.Field: condition.Value}
		case "nil":
			conf = builder.IsNull{condition.Field}
		case "nnil":
			conf = builder.NotNull{condition.Field}
		case "!=":
			conf = builder.Neq{condition.Field: condition.Value}
		case "in":
			o := condition.Value.([]interface{})
			if len(o) > 0 {
				conf = builder.In(condition.Field, o)
			} else {
				conf = builder.Eq{"ID": "NULL"}
			}
		case "between":
			o := condition.Value.([]interface{})
			if len(o) == 2 {
				conf = builder.Between{Col: condition.Field, LessVal: o[0], MoreVal: o[1]}
			}
		}
		str, args, err := builder.ToSQL(conf)
		if err != nil {
			//logging.Error("条件错误：%v", err)
			return errors.New("条件错误：" + fmt.Sprintf("%v", err))
		}
		qi := QueryItem{
			Relation: condition.Relation,
			Str:      str,
			Args:     args,
		}
		qis = append(qis, qi)
	}
	for _, qi := range qis {
		if qi.Relation == "and" {
			session.And("("+qi.Str+")", qi.Args...)
		} else if qi.Relation == "or" {
			session.Or("("+qi.Str+")", qi.Args...)
		}
	}
	return nil
}
