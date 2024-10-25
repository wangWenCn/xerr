package xerr

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"runtime"
	"strings"

	"github.com/wangWenCn/traceLog"
	"github.com/zeromicro/go-zero/core/logx"
)

var sourceMap map[string][]ErrLine

type ErrLine struct {
	Code string
	Line int
}

type CodeError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
	err     []error
}

func (e *CodeError) Error() string {
	return e.Message
}

func NewErrCodeMsg(errCode int64, errMsg string, e ...error) *CodeError {
	err := &CodeError{Code: errCode, Message: errMsg, err: e}
	logErrorInfo(err)
	return err
}

func logErrorInfo(err *CodeError) {
	pc, file, line, ok := runtime.Caller(2)
	if sourceMap[file] == nil {
		sourceMap[file] = parseCode(file)
	}
	if !ok {
		logx.WithContext(traceLog.GetGoroutineContext()).WithCallerSkip(2).Error(err)
		return
	}
	fields := make([]logx.LogField, 0)
	funcName := runtime.FuncForPC(pc).Name()
	split := strings.Split(funcName, ".")
	if len(split) > 1 {
		split[1] = strings.Trim(split[1], "*()")
		fields = append(fields, logx.Field("module", split[1]))
		fields = append(fields, logx.Field("function", split[2]))
	}
	fields = append(fields, logx.Field("sentence", getErrLine(file, line)))
	if err != nil {
		if len(err.err) > 0 && err.err[0] != nil {
			fields = append(fields, logx.Field("error", err.err[0].Error()))
		} else {
			fields = append(fields, logx.Field("error", err.Error()))
		}
	}
	logx.WithContext(traceLog.GetGoroutineContext()).WithCallerSkip(2).Errorw("err", fields...)
}

func getErrLine(file string, line int) string {

	if len(file) == 0 {
		return ""
	}
	l, r := 0, len(sourceMap[file])-1
	if l > r {
		return ""
	}
	for l < r {
		mid := (l + r + 1) / 2
		if sourceMap[file][mid].Line <= line {
			l = mid
		} else {
			r = mid - 1
		}
	}

	if sourceMap[file][l].Line <= line {
		return sourceMap[file][l].Code
	}
	return ""
}

func NewErrCode(errCode int64, e ...error) *CodeError {
	err := &CodeError{
		Code:    errCode,
		Message: MapErrMsg(errCode),
		err:     e,
	}
	logErrorInfo(err)
	return err
}

func NewErrMsg(errMsg string, e ...error) *CodeError {
	err := &CodeError{
		Code:    ServerCommonError,
		Message: errMsg,
		err:     e,
	}
	logErrorInfo(err)
	return err
}

func isErrIdent(expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok {
		if strings.Contains(strings.ToLower(ident.Name), "err") {
			return true
		}
	}
	return false
}

func findErrAssignments(f *ast.File, fSet *token.FileSet) []ast.Stmt {
	var errAssignments []ast.Stmt
	ast.Inspect(f, func(n ast.Node) bool {
		if assign, ok := n.(*ast.AssignStmt); ok {
			for _, lhs := range assign.Lhs {
				if isErrIdent(lhs) {
					if assign.Tok == token.ASSIGN || assign.Tok == token.DEFINE {
						errAssignments = append(errAssignments, assign)
					}
				}
			}
		}
		//if ifStmt, ok := n.(*ast.IfStmt); ok {
		//	var buf strings.Builder
		//	_ = printer.Fprint(&buf, fSet, ifStmt)
		//	for _, stmt := range ifStmt.Body.List {
		//		if ret, ok := stmt.(*ast.ReturnStmt); ok {
		//			for _, result := range ret.Results {
		//				if ident, ok := result.(*ast.Ident); ok && strings.ToLower(ident.Name) == "xerr" {
		//					fmt.Printf("stmt: %#v\n", stmt)
		//				}
		//			}
		//		}
		//	}
		//}
		return true
	})
	return errAssignments
}

func parseCode(filePath string) []ErrLine {
	fSet := token.NewFileSet()
	file, err := parser.ParseFile(fSet, filePath, nil, parser.AllErrors)
	if err != nil {
		return []ErrLine{}
	}

	errAssignments := findErrAssignments(file, fSet)

	lines := make([]ErrLine, 0)

	for _, stmt := range errAssignments {
		start := fSet.Position(stmt.Pos())
		var buf strings.Builder
		_ = printer.Fprint(&buf, fSet, stmt)
		line := buf.String()
		line = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(line, "\n", ""), "\t", ""))
		if strings.Contains(line, "Transaction") {
			continue
		}
		lines = append(lines, ErrLine{
			Code: line,
			Line: start.Line,
		})
	}
	return lines
}

func init() {
	sourceMap = make(map[string][]ErrLine)
}
