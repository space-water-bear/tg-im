func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp {{.upperStartCamelObject}}
	err := m.QueryRowCtx(ctx, &resp, {{.cacheKeyVariable}}, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query :=  fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
		return conn.QueryRowCtx(ctx, v, query, {{.lowerStartCamelPrimaryKey}})
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{else}}query := fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
	var resp {{.upperStartCamelObject}}
	err := m.conn.QueryRowCtx(ctx, &resp, query, {{.lowerStartCamelPrimaryKey}})
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{end}}
}

func (m *default{{.upperStartCamelObject}}Model) FindListByPage(ctx context.Context, page int, pageSize int, condition string) ([]*{{.upperStartCamelObject}}, error) {
    if page <= 0 {
        page = 1
    }
    offset := (page - 1) * pageSize
    query := fmt.Sprintf("select %s from %s where %s limit %d offset %d", {{.lowerStartCamelObject}}Rows, m.table, condition, pageSize, offset)
    var resp []*{{.upperStartCamelObject}}
    {{if .withCache}}
    err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
    {{else}}
    err := m.conn.QueryRowsCtx(ctx, &resp, query)
    {{end}}
    switch err {
    case nil:
        return resp, nil
    default:
        return nil, err
    }
}

func (m *default{{.upperStartCamelObject}}Model) FindCountByCondition(ctx context.Context, condition string) (int64, error) {
    query := fmt.Sprintf("select count(id) from %s where %s", m.table, condition)
    var count int64
    {{if .withCache}}
    err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
    {{else}}
    err := m.conn.QueryRowCtx(ctx, &count, query)
    {{end}}
    switch err {
    case nil:
        return count, nil
    default:
        return 0, err
    }
}
