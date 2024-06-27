FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error)
FindListByPage(ctx context.Context, page int, pageSize int, condition string) ([]*{{.upperStartCamelObject}}, error)
FindCountByCondition(ctx context.Context, condition string) (int64, error)