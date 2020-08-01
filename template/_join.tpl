{{define "join"}}{{$v := .Name}}{{range .Joins}}{{$upperTable := .Table.UpperCamel}}{{$snakeTable := .Table.Snake}}
func (q *{{printf $v.UpperCamel}}Query) Join{{.Table.UpperCamel}}(this {{printf $v.LowerCamel}}, target {{.Table.LowerCamel}}, joinType query.JoinType, direction... query.JoinDirection) *Join{{printf $v.UpperCamel}}{{printf $upperTable}}Query {
    var first string
    var second string
    switch this { {{range .Column}}
    case {{printf $v.UpperCamel}}{{.This.UpperCamel}}:
        first = "{{printf $v.Snake}}.{{.This.Snake}}"
        if target != {{printf $upperTable}}{{.Target.UpperCamel}} {
            return nil
        }
        second = "{{$snakeTable}}.{{.Target.UpperCamel}}"{{end}}
    default:
        return nil
    }
    return &Join{{printf $v.UpperCamel}}{{printf $upperTable}}Query{
        db: q.db,
        joins: []*query.Join{
            query.NewJoin(joinType, first, second, direction...),
        },
    }
}
{{end}}{{end}}