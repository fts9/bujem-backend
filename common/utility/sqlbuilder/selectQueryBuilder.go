package sqlbuilder

import (
	"errors"
	"fmt"
	"log"
)

const (
	keywordSelect    = "SELECT"
	keywordFrom      = "FROM"
	keywordWhere     = "WHERE"
	keywordAnd       = " AND \n"
	allColumns       = "*"
	columnDelimiter  = ", \n"
	objectDelimiter  = "."
	comparatorEquals = "="
	queryParameter   = "$%d"
)

type SelectQueryBuilder interface {
	Columns(...string) SelectQueryBuilder
	Table(string, string) SelectQueryBuilder
	Where(string) SelectQueryBuilder
	Build() (string, error)
}

type selectQueryBuilder struct {
	columns   []string
	table     []string
	where     []string
	logOutput bool
}

func NewSelectQueryWithLogging(logOutput bool) SelectQueryBuilder {
	return &selectQueryBuilder{
		columns: nil,
		table:   nil,
		where:   nil}
}

func NewSelectQuery() SelectQueryBuilder {
	return NewSelectQueryWithLogging(false)
}

func (builder *selectQueryBuilder) Columns(columns ...string) SelectQueryBuilder {
	builder.columns = append(builder.columns, columns...)
	return builder
}

func (builder *selectQueryBuilder) Table(schema string, table string) SelectQueryBuilder {
	var tableFQN = table
	if len(schema) > 0 {
		tableFQN = schema + objectDelimiter + tableFQN
	}
	builder.table = append(builder.table, tableFQN)
	return builder
}

func (builder *selectQueryBuilder) Where(column string) SelectQueryBuilder {
	builder.where = append(builder.where, column)
	return builder
}

func (builder *selectQueryBuilder) Build() (string, error) {
	var queryString = keywordSelect + " "
	if len(builder.columns) == 0 {
		queryString += allColumns
	} else {
		queryString += arrayToCSV(builder.columns, columnDelimiter)
	}

	queryString += " " + keywordFrom + " "
	if len(builder.table) == 0 {
		return "", errors.New("Cannot generate a SELECT query without a source table specified")
	}
	queryString += arrayToCSV(builder.table, columnDelimiter)

	if len(builder.where) > 0 {
		queryString += " " + keywordWhere + " "
		whereClauses := make([]string, len(builder.where))
		for whereColumnIndex := range builder.where {
			whereClauses[whereColumnIndex] = builder.where[whereColumnIndex] + " " + comparatorEquals + " " + getParameterValue(whereColumnIndex)
		}
		queryString += arrayToCSV(whereClauses, keywordAnd)
	}

	//if builder.logOutput {
	log.Printf("Generated SELECT query:")
	log.Printf(queryString)
	//}
	return queryString, nil
}

func arrayToCSV(arrayToConvert []string, delimiter string) string {
	csvList := ""
	if len(arrayToConvert) > 0 {
		for arrayIndex := range arrayToConvert {
			csvList += arrayToConvert[arrayIndex] + delimiter
		}
		return csvList[:len(csvList)-len(delimiter)]
	}
	return csvList
}

func getParameterValue(parameterIndex int) string {
	return fmt.Sprintf(queryParameter, parameterIndex+1)
}
